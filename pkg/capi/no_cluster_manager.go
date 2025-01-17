// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.
package capi

import (
	"fmt"
)

// compile time checking for interface implementation
var _ ClusterLifeCycleManager = &noClusterManager{}

const emptyKubeconfig = `
apiVersion: v1
kind: ""
clusters:
users:
contexts:
`

func newNoClusterManager(actualConfig ClusterConfig) (ClusterLifeCycleManager, error) {
	return &noClusterManager{
		config: actualConfig,
	}, nil
}

// noClusterManager ClusterLifecycleManager impl for testing - does not perform any cluster operations
type noClusterManager struct {
	config ClusterConfig
}

func (r *noClusterManager) GetKubeConfig() (string, error) {
	fmt.Println("get kubeconfig for noCluster")
	return emptyKubeconfig, nil
}

func (r *noClusterManager) Init() error {
	fmt.Println("Init noCluster")
	return nil
}

func (r *noClusterManager) GetConfig() ClusterConfig {
	return r.config
}

func (r *noClusterManager) Create() error {
	fmt.Printf("Creating noCluster with config %v\n", r.config)
	return nil
}

func (r *noClusterManager) Destroy() error {
	fmt.Println("Destroying noCluster")
	return nil
}
