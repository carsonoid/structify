package main

import (
	"time"

	v1 "k8s.io/api/core/v1"
)

func main() {
	obj := v1.ConfigMap{
		TypeMeta: v1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:            "test",
			GenerateName:    "",
			Namespace:       "default",
			SelfLink:        "/api/v1/namespaces/default/configmaps/test",
			UID:             "37419574-f36e-11e9-8c11-0242ac180002",
			ResourceVersion: "601",
			Generation:      0,
			CreationTimestamp: v1.Time{
				Time: time.Time{},
			},
			DeletionTimestamp:          nil,
			DeletionGracePeriodSeconds: nil,
			Labels:                     map[string]string{},
			Annotations:                map[string]string{},
			OwnerReferences:            nil,
			Finalizers:                 nil,
			ClusterName:                "",
			ManagedFields:              nil,
		},
		Data: map[string]string{
			"testkey1": "testvalue1",
			"testkey2": "testvalue2",
		},
		BinaryData: map[string][]uint8{},
	}
}
