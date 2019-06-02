docker build -t spikelabs/corebos-kubernetes -t spikelabs/corebos-kubernetes:$SHA .
docker push spikelabs/corebos-kubernetes
docker push spikelabs/corebos-kubernetes:$SHA

apt-get update && apt-get install -y apt-transport-https nano bash curl gnupg gnupg2 gnupg1
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | tee -a /etc/apt/sources.list.d/kubernetes.list
apt-get update
apt-get install -y kubectl

mkdir ~/.kube

curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $DIGITAL_OCEAN_TOKEN" "https://api.digitalocean.com/v2/kubernetes/clusters/$CLUSTER_ID/kubeconfig" > kubeconfig.yaml

mv kubeconfig ~/.kube/config

kubectl apply -f ./stage-prod/k8s