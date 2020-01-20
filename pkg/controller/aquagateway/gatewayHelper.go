package aquagateway

import (
	"fmt"
	"os"

	"github.com/niso120b/aqua-operator/pkg/controller/common"

	operatorv1alpha1 "github.com/niso120b/aqua-operator/pkg/apis/operator/v1alpha1"
	"github.com/niso120b/aqua-operator/pkg/consts"
	"github.com/niso120b/aqua-operator/pkg/utils/extra"
	"github.com/niso120b/aqua-operator/pkg/utils/k8s/services"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type GatewayParameters struct {
	Gateway *operatorv1alpha1.AquaGateway
}

type AquaGatewayHelper struct {
	Parameters GatewayParameters
}

func newAquaGatewayHelper(cr *operatorv1alpha1.AquaGateway) *AquaGatewayHelper {
	params := GatewayParameters{
		Gateway: cr,
	}

	return &AquaGatewayHelper{
		Parameters: params,
	}
}

func (gw *AquaGatewayHelper) newDeployment(cr *operatorv1alpha1.AquaGateway) *appsv1.Deployment {
	pullPolicy, registry, repository, tag := extra.GetImageData("gateway", cr.Spec.Infrastructure.Version, cr.Spec.GatewayService.ImageData)

	image := os.Getenv("RELATED_IMAGE_GATEWAY")
	if image == "" {
		image = fmt.Sprintf("%s/%s:%s", registry, repository, tag)
	}

	labels := map[string]string{
		"app":                cr.Name + "-gateway",
		"deployedby":         "aqua-operator",
		"aquasecoperator_cr": cr.Name,
		"type":               "aqua-gateway",
	}
	annotations := map[string]string{
		"description": "Deploy the aqua gateway server",
	}
	env_vars := gw.getEnvVars(cr)
	deployment := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "apps/v1",
			Kind:       "Deployment",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        fmt.Sprintf(consts.GatewayDeployName, cr.Name),
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: extra.Int32Ptr(int32(cr.Spec.GatewayService.Replicas)),
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: cr.Spec.Infrastructure.ServiceAccount,
					Containers: []corev1.Container{
						{
							Name:            "aqua-gateway",
							Image:           image,
							ImagePullPolicy: corev1.PullPolicy(pullPolicy),
							Ports: []corev1.ContainerPort{
								{
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: 3622,
								},
							},
							Env: env_vars,
						},
					},
				},
			},
		},
	}

	if cr.Spec.GatewayService.Resources != nil {
		deployment.Spec.Template.Spec.Containers[0].Resources = *cr.Spec.GatewayService.Resources
	}

	if cr.Spec.GatewayService.LivenessProbe != nil {
		deployment.Spec.Template.Spec.Containers[0].LivenessProbe = cr.Spec.GatewayService.LivenessProbe
	}

	if cr.Spec.GatewayService.ReadinessProbe != nil {
		deployment.Spec.Template.Spec.Containers[0].ReadinessProbe = cr.Spec.GatewayService.ReadinessProbe
	}

	if cr.Spec.GatewayService.NodeSelector != nil {
		if len(cr.Spec.GatewayService.NodeSelector) > 0 {
			deployment.Spec.Template.Spec.NodeSelector = cr.Spec.GatewayService.NodeSelector
		}
	}

	if cr.Spec.GatewayService.Affinity != nil {
		deployment.Spec.Template.Spec.Affinity = cr.Spec.GatewayService.Affinity
	}

	if cr.Spec.GatewayService.Tolerations != nil {
		if len(cr.Spec.GatewayService.Tolerations) > 0 {
			deployment.Spec.Template.Spec.Tolerations = cr.Spec.GatewayService.Tolerations
		}
	}

	if len(cr.Spec.Common.ImagePullSecret) != 0 {
		deployment.Spec.Template.Spec.ImagePullSecrets = []corev1.LocalObjectReference{
			corev1.LocalObjectReference{
				Name: cr.Spec.Common.ImagePullSecret,
			},
		}
	}

	return deployment
}

func (gw *AquaGatewayHelper) getEnvVars(cr *operatorv1alpha1.AquaGateway) []corev1.EnvVar {
	envsHelper := common.NewAquaEnvsHelper(cr.Spec.Infrastructure, cr.Spec.Common, cr.Spec.ExternalDb, cr.Name)
	result, _ := envsHelper.GetDbEnvVars()

	result = append(result, corev1.EnvVar{
		Name:  "HEALTH_MONITOR",
		Value: "0.0.0.0:8082",
	})

	result = append(result, corev1.EnvVar{
		Name:  "AQUA_CONSOLE_SECURE_ADDRESS",
		Value: fmt.Sprintf("%s:443", fmt.Sprintf(consts.ServerServiceName, cr.Name)),
	})

	result = append(result, corev1.EnvVar{
		Name:  "SCALOCK_GATEWAY_PUBLIC_IP",
		Value: fmt.Sprintf(consts.GatewayServiceName, cr.Name),
	})

	return result
}

func (gw *AquaGatewayHelper) newService(cr *operatorv1alpha1.AquaGateway) *corev1.Service {
	selectors := map[string]string{
		"app": fmt.Sprintf("%s-gateway", cr.Name),
	}

	ports := []corev1.ServicePort{
		{
			Port:       3622,
			TargetPort: intstr.FromInt(3622),
			Name:       "aqua-gate",
		},
		{
			Port:       8443,
			TargetPort: intstr.FromInt(8443),
			Name:       "aqua-gate-ssl",
		},
	}

	service := services.CreateService(cr.Name,
		cr.Namespace,
		fmt.Sprintf(consts.GatewayServiceName, cr.Name),
		fmt.Sprintf("%s-gateway", cr.Name),
		"Service for aqua gateway components",
		cr.Spec.GatewayService.ServiceType,
		selectors,
		ports)

	return service
}
