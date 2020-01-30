kubectl create ns aqua
kubectl create -f  deploy/crds/operator_v1alpha1_aquadatabase_crd.yaml 
kubectl create -f  deploy/crds/operator_v1alpha1_aquagateway_crd.yaml 
kubectl create -f  deploy/crds/operator_v1alpha1_aquaserver_crd.yaml 
kubectl create -f  deploy/crds/operator_v1alpha1_aquaenforcer_crd.yaml
kubectl create -f  deploy/crds/operator_v1alpha1_aquacsp_crd.yaml 
kubectl create -f  deploy/crds/operator_v1alpha1_aquascanner_crd.yaml

kubectl create -f deploy/service_account.yaml -n aqua
kubectl create -f deploy/clusterrole.yaml
kubectl create -f deploy/clusterrole_binding.yaml
kubectl create -f deploy/operator.yaml -n aqua
