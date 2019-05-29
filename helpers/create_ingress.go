package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	v1beta "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"os"
)

func CreateIngress(ingressData *pb.Ingress, clientSet *kubernetes.Clientset) error {
	ingressClient := clientSet.ExtensionsV1beta1().Ingresses(os.Getenv("NAMESPACE"))

	ingress := &v1beta.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name: ingressData.Name,
			Annotations: map[string]string{
				"kubernetes.io/ingress.class":       "nginx",
				"certmanager.k8s.io/cluster-issuer": "corebos-cluster-issuer",
			},
		},
		Spec: v1beta.IngressSpec{
			TLS: []v1beta.IngressTLS{
				{
					Hosts: []string{
						ingressData.SubDomain,
					},
					SecretName: ingressData.Name + "-cert",
				},
			},
			Rules: []v1beta.IngressRule{
				{
					Host: ingressData.SubDomain,
					IngressRuleValue: v1beta.IngressRuleValue{
						HTTP: &v1beta.HTTPIngressRuleValue{
							Paths: []v1beta.HTTPIngressPath{
								{
									Backend: v1beta.IngressBackend{
										ServiceName: ingressData.Resource,
										ServicePort: intstr.IntOrString{
											IntVal: 80,
											Type:   intstr.Int,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := ingressClient.Create(ingress)
	return err
}
