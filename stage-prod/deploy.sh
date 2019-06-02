docker build -t spikelabs/corebos-kubernetes -t spikelabs/corebos-kubernetes:$SHA .
docker push spikelabs/corebos-kubernetes
docker push spikelabs/corebos-kubernetes:$SHA

curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl

mkdir ~/.kube

curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $DIGITAL_OCEAN_TOKEN" "https://api.digitalocean.com/v2/kubernetes/clusters/$CLUSTER_ID/kubeconfig" > kubeconfig.yaml

mv kubeconfig.yaml ~/.kube/config

kubectl apply -f ./stage-prod/k8s

kubectl set image deployments/corebos-kubernetes-deployment corebos-kubernetes=spikelabs/corebos-kubernetes:$SHA
