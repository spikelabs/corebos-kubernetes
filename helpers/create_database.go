package helpers

import (
	pb "corebos-kubernetes/kubernetes"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"os"
)

func CreateDatabase(databaseData *pb.Database, claimName string, clientSet *kubernetes.Clientset) error {

	deploymentsClient := clientSet.AppsV1().Deployments(os.Getenv("NAMESPACE"))

	number := int32(1)

	database := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: databaseData.Name,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &number,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"component": databaseData.Label,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"component": databaseData.Label,
					},
				},
				Spec: apiv1.PodSpec{
					Volumes: []apiv1.Volume{
						{
							Name: databaseData.Name + "-storage",
							VolumeSource: apiv1.VolumeSource{
								PersistentVolumeClaim: &apiv1.PersistentVolumeClaimVolumeSource{
									ClaimName: claimName,
								},
							},
						},
					},
					Containers: []apiv1.Container{
						{
							Name:  databaseData.Label,
							Image: "mysql:5.7",
							Ports: []apiv1.ContainerPort{
								{
									ContainerPort: 3306,
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name:  "MYSQL_DATABASE",
									Value: databaseData.DbDatabase,
								},
								{
									Name:  "MYSQL_USER",
									Value: databaseData.DbUsername,
								},
								{
									Name:  "MYSQL_PASSWORD",
									Value: databaseData.DbPassword,
								},
								{
									Name:  "MYSQL_ROOT_PASSWORD",
									Value: "root",
								},
							},
							VolumeMounts: []apiv1.VolumeMount{
								{
									Name:      databaseData.Name + "-storage",
									MountPath: "/var/lib/mysql",
									SubPath:   "mysql",
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := deploymentsClient.Create(database)
	return err
}
