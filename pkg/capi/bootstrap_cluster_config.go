// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.
package capi

import (
	"os"
)

const (
	BootstrapImageEnvVar = "VZ_BOOTSTRAP_IMAGE"
	bootstrapClusterName = "vz-capi-bootstrap"
)

type bootstrapClusterConfig struct{}

func (r bootstrapClusterConfig) GetClusterName() string {
	return bootstrapClusterName
}

func (r bootstrapClusterConfig) GetType() string {
	return KindClusterType
}

func (r bootstrapClusterConfig) GetContainerImage() string {
	return os.Getenv(BootstrapImageEnvVar)
}