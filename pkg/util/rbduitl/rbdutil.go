package rbdutil

import (
	"github.com/goodrain/rainbond-operator/pkg/apis/rainbond/v1alpha1"
	"github.com/goodrain/rainbond-operator/pkg/util/constants"
	"path"
)

type Labels map[string]string

func LabelsForRainbond() map[string]string {
	return map[string]string{
		"belongTo": "rainbond-operator",
	}
}

// GetStorageClass returns storage class name based on rainbondcluster.
func GetStorageClass(cluster *v1alpha1.RainbondCluster) string {
	if cluster.Spec.StorageClassName == "" {
		return constants.DefStorageClass
	}
	return cluster.Spec.StorageClassName
}

// GetImageRepository returns image repository name based on rainbondcluster.
func GetImageRepository(cluster *v1alpha1.RainbondCluster) string {
	if cluster.Spec.ImageHub == nil {
		return constants.DefImageRepository
	}
	return path.Join(cluster.Spec.ImageHub.Domain, cluster.Spec.ImageHub.Namespace)
}
