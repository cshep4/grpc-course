
.PHONY: create-gke-cluster
create-gke-cluster:
	export CLUSTER=test-grpc-cluster
	gcloud container clusters create $CLUSTER --preemptible --num-nodes=2
	gcloud components install gke-gcloud-auth-plugin
	export USE_GKE_GCLOUD_AUTH_PLUGIN=True
	gcloud container clusters get-credentials $CLUSTER

.PHONY: create-gh-packages-secret
create-gh-packages-secret:
	kubectl create secret docker-registry github-packages-secret \
    --docker-server=https://ghcr.io \
    --docker-username=<user-name> \
    --docker-password=<token>

.PHONY: create-cloudflare-token-secret
create-cloudflare-token-secret:
	kubectl apply -f deploy/cloudflare-token-secret.yaml

# https://cert-manager.io/docs/tutorials/getting-started-with-cert-manager-on-google-kubernetes-engine-using-lets-encrypt-for-ingress-ssl/#prerequisites
.PHONY: init-cert-manager
init-cert-manager:
	kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.5/cert-manager.yaml

.PHONY: init-nginx-controller
init-nginx-controller:
	helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	helm repo update
	helm install my-nginx ingress-nginx/ingress-nginx \
	  --namespace ingress-nginx \
	  --create-namespace \
	  --set controller.publishService.enabled=true

cleanup:
	gcloud container clusters delete test-grpc-cluster