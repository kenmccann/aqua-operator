package aquacsp

import (
	"fmt"

	operatorv1alpha1 "github.com/niso120b/aqua-operator/pkg/apis/operator/v1alpha1"
	"github.com/niso120b/aqua-operator/pkg/consts"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CspParameters struct {
	AquaCsp *operatorv1alpha1.AquaCsp
}

type AquaCspHelper struct {
	Parameters CspParameters
}

func newAquaCspHelper(cr *operatorv1alpha1.AquaCsp) *AquaCspHelper {
	params := CspParameters{
		AquaCsp: cr,
	}

	return &AquaCspHelper{
		Parameters: params,
	}
}

func (csp *AquaCspHelper) newAquaDatabase(cr *operatorv1alpha1.AquaCsp) *operatorv1alpha1.AquaDatabase {
	labels := map[string]string{
		"app":                cr.Name + "-csp",
		"deployedby":         "aqua-operator",
		"aquasecoperator_cr": cr.Name,
	}
	annotations := map[string]string{
		"description": "Deploy Aqua Database (not for production environments)",
	}
	aquadb := &operatorv1alpha1.AquaDatabase{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "operator.aquasec.com/v1alpha1",
			Kind:       "AquaDatabase",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        cr.Name,
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: operatorv1alpha1.AquaDatabaseSpec{
			Infrastructure: csp.Parameters.AquaCsp.Spec.Infrastructure,
			Common:         csp.Parameters.AquaCsp.Spec.Common,
			DbService:      csp.Parameters.AquaCsp.Spec.DbService,
			DiskSize:       csp.Parameters.AquaCsp.Spec.Common.DbDiskSize,
		},
	}

	return aquadb
}

func (csp *AquaCspHelper) newAquaGateway(cr *operatorv1alpha1.AquaCsp) *operatorv1alpha1.AquaGateway {
	labels := map[string]string{
		"app":                cr.Name + "-csp",
		"deployedby":         "aqua-operator",
		"aquasecoperator_cr": cr.Name,
	}
	annotations := map[string]string{
		"description": "Deploy Aqua Gateway",
	}
	aquadb := &operatorv1alpha1.AquaGateway{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "operator.aquasec.com/v1alpha1",
			Kind:       "AquaGateway",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        cr.Name,
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: operatorv1alpha1.AquaGatewaySpec{
			Infrastructure: csp.Parameters.AquaCsp.Spec.Infrastructure,
			Common:         csp.Parameters.AquaCsp.Spec.Common,
			GatewayService: csp.Parameters.AquaCsp.Spec.GatewayService,
			ExternalDb:     csp.Parameters.AquaCsp.Spec.ExternalDb,
		},
	}

	return aquadb
}

func (csp *AquaCspHelper) newAquaServer(cr *operatorv1alpha1.AquaCsp) *operatorv1alpha1.AquaServer {
	labels := map[string]string{
		"app":                cr.Name + "-csp",
		"deployedby":         "aqua-operator",
		"aquasecoperator_cr": cr.Name,
	}
	annotations := map[string]string{
		"description": "Deploy Aqua Server",
	}
	aquadb := &operatorv1alpha1.AquaServer{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "operator.aquasec.com/v1alpha1",
			Kind:       "AquaServer",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        cr.Name,
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: operatorv1alpha1.AquaServerSpec{
			Infrastructure: csp.Parameters.AquaCsp.Spec.Infrastructure,
			Common:         csp.Parameters.AquaCsp.Spec.Common,
			ServerService:  csp.Parameters.AquaCsp.Spec.ServerService,
			ExternalDb:     csp.Parameters.AquaCsp.Spec.ExternalDb,
			LicenseToken:   csp.Parameters.AquaCsp.Spec.LicenseToken,
			AdminPassword:  csp.Parameters.AquaCsp.Spec.AdminPassword,
		},
	}

	return aquadb
}

func (csp *AquaCspHelper) newAquaScanner(cr *operatorv1alpha1.AquaCsp) *operatorv1alpha1.AquaScanner {
	labels := map[string]string{
		"app":                cr.Name + "-csp",
		"deployedby":         "aqua-operator",
		"aquasecoperator_cr": cr.Name,
	}
	annotations := map[string]string{
		"description": "Deploy Aqua Scanner",
	}
	scanner := &operatorv1alpha1.AquaScanner{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "operator.aquasec.com/v1alpha1",
			Kind:       "AquaScanner",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:        cr.Name,
			Namespace:   cr.Namespace,
			Labels:      labels,
			Annotations: annotations,
		},
		Spec: operatorv1alpha1.AquaScannerSpec{
			Infrastructure: cr.Spec.Infrastructure,
			Common:         cr.Spec.Common,
			ScannerService: cr.Spec.ScannerService,
			Login: &operatorv1alpha1.AquaLogin{
				Username: "administrator",
				Password: cr.Spec.AdminPassword,
				Host:     fmt.Sprintf("http://%s:8080", fmt.Sprintf(consts.ServerServiceName, cr.Name)),
			},
		},
	}

	return scanner
}
