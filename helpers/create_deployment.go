package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func CreateDeployment (deploymentData *pb.Deployment, claimName string, clientset *kubernetes.Clientset) (error) {

	deploymentsClient := clientset.AppsV1().Deployments(os.Getenv("NAMESPACE"))

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: deploymentData.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deploymentData.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"component": deploymentData.Label,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"component": deploymentData.Label,
					},
				},
				Spec: apiv1.PodSpec{
					Volumes: []apiv1.Volume{
						{
							Name: deploymentData.Name + "-storage",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: claimName,
								},
							},
						},
					},
					Containers: []apiv1.Container{
						{
							Name:  deploymentData.Label,
							Image: "spikelabs/corebos-gdpr",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name: "COREBOS_DBSERVER",
									Value: deploymentData.DbHost,
								},
								{
									Name: "COREBOS_DBUSER",
									Value: deploymentData.DbUsername,
								},
								{
									Name: "COREBOS_DBPASS",
									Value: deploymentData.DbPassword,
								},
								{
									Name: "COREBOS_DBNAME",
									Value: deploymentData.DbDatabase,
								},
								{
									Name: "COREBOS_SITEURL",
									Value: deploymentData.SiteUrl,
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name: deploymentData.Name + "-storage",
									MountPath: "/www/storage",
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := deploymentsClient.Create(deployment)
	if err != nil {
		return err
	}
	return nil
}
