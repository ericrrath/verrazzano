// Copyright (c) 2020, 2021, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package helm

import (
	"errors"
	"os/exec"
	"testing"

	"go.uber.org/zap"

	"github.com/stretchr/testify/assert"
)

const ns = "my_namespace"
const chartdir = "my_charts"
const release = "my_release"
const missingRelease = "no_release"

// upgradeRunner is used to test Helm upgrade without actually running an OS exec command
type upgradeRunner struct {
	t *testing.T
}

// getValuesRunner is used to test Helm get values without actually running an OS exec command
type getValuesRunner struct {
	t *testing.T
}

// badRunner is used to test Helm errors without actually running an OS exec command
type badRunner struct {
	t *testing.T
}

// foundRunner is used to test helm status command
type foundRunner struct {
	t *testing.T
}

// TestGetValues tests the Helm get values command
// GIVEN a set of upgrade parameters
//  WHEN I call Upgrade
//  THEN the Helm upgrade returns success and the cmd object has correct values
func TestGetValues(t *testing.T) {
	assert := assert.New(t)
	SetCmdRunner(getValuesRunner{t: t})
	defer SetDefaultRunner()

	stdout, err := GetValues(zap.S(), release, ns)
	assert.NoError(err, "GetValues returned an error")
	assert.NotZero(stdout, "GetValues stdout should not be empty")
}

// TestUpgrade tests the Helm upgrade command
// GIVEN a set of upgrade parameters
//  WHEN I call Upgrade
//  THEN the Helm upgrade returns success and the cmd object has correct values
func TestUpgrade(t *testing.T) {
	overrideYaml := "my-override.yaml"

	assert := assert.New(t)
	SetCmdRunner(upgradeRunner{t: t})
	defer SetDefaultRunner()

	stdout, stderr, err := Upgrade(zap.S(), release, ns, chartdir, false, false, "", overrideYaml)
	assert.NoError(err, "Upgrade returned an error")
	assert.Len(stderr, 0, "Upgrade stderr should be empty")
	assert.NotZero(stdout, "Upgrade stdout should not be empty")
}

// TestUpgradeFail tests the Helm upgrade command failure condition
// GIVEN a set of upgrade parameters and a fake runner that fails
//  WHEN I call Upgrade
//  THEN the Helm upgrade returns an error
func TestUpgradeFail(t *testing.T) {
	assert := assert.New(t)
	SetCmdRunner(badRunner{t: t})
	defer SetDefaultRunner()

	stdout, stderr, err := Upgrade(zap.S(), release, ns, "", false, false, "")
	assert.Error(err, "Upgrade should have returned an error")
	assert.Len(stdout, 0, "Upgrade stdout should be empty")
	assert.NotZero(stderr, "Upgrade stderr should not be empty")
}

// TestIsReleaseInstalled tests checking if a Helm release is installed
// GIVEN a release name and namespace
//  WHEN I call IsReleaseInstalled
//  THEN the function returns success and found equal true
func TestIsReleaseInstalled(t *testing.T) {
	assert := assert.New(t)
	SetCmdRunner(foundRunner{t: t})
	defer SetDefaultRunner()
	SetChartStatusFunction(func(releaseName string, namespace string) (string, error) {
		return ChartStatusDeployed, nil
	})
	defer SetDefaultChartStatusFunction()

	found, err := IsReleaseInstalled(release, ns)
	assert.NoError(err, "IsReleaseInstalled returned an error")
	assert.True(found, "Release not found")
}

// TestIsReleaseNotInstalled tests checking if a Helm release is not installed
// GIVEN a release name and namespace
//  WHEN I call IsReleaseInstalled
//  THEN the function returns success and the correct found status
func TestIsReleaseNotInstalled(t *testing.T) {
	assert := assert.New(t)
	SetCmdRunner(foundRunner{t: t})
	defer SetDefaultRunner()
	SetChartStatusFunction(func(releaseName string, namespace string) (string, error) {
		return ChartNotFound, nil
	})
	defer SetDefaultChartStatusFunction()

	found, err := IsReleaseInstalled(missingRelease, ns)
	assert.NoError(err, "IsReleaseInstalled returned an error")
	assert.False(found, "Release should not be found")
}

// TestIsReleaseInstalledFailed tests failure when checking if a Helm release is installed
// GIVEN a bad release name and namespace
//  WHEN I call IsReleaseInstalled
//  THEN the function returns a failure
func TestIsReleaseInstalledFailed(t *testing.T) {
	assert := assert.New(t)
	SetCmdRunner(foundRunner{t: t})
	defer SetDefaultRunner()

	found, err := IsReleaseInstalled("", ns)
	assert.Error(err, "IsReleaseInstalled should have returned an error")
	assert.False(found, "Release should not be found")
}

// Run should assert the command parameters are correct then return a success with stdout contents
func (r upgradeRunner) Run(cmd *exec.Cmd) (stdout []byte, stderr []byte, err error) {
	assert := assert.New(r.t)
	assert.Contains(cmd.Path, "helm", "command should contain helm")
	assert.Contains(cmd.Args[0], "helm", "args should contain helm")
	assert.Contains(cmd.Args[1], "upgrade", "args should contain upgrade")
	assert.Contains(cmd.Args[2], release, "args should contain release name")
	assert.Contains(cmd.Args[3], chartdir, "args should contain chart dir ")

	return []byte("success"), []byte(""), nil
}

// Run should assert the command parameters are correct then return a success with stdout contents
func (r getValuesRunner) Run(cmd *exec.Cmd) (stdout []byte, stderr []byte, err error) {
	assert := assert.New(r.t)
	assert.Contains(cmd.Path, "helm", "command should contain helm")
	assert.Contains(cmd.Args[0], "helm", "args should contain helm")
	assert.Contains(cmd.Args[1], "get", "args should contain get")
	assert.Contains(cmd.Args[2], "values", "args should contain get")
	assert.Contains(cmd.Args[3], release, "args should contain release name")
	return []byte("success"), []byte(""), nil
}

// Run should return an error with stderr contents
func (r badRunner) Run(cmd *exec.Cmd) (stdout []byte, stderr []byte, err error) {
	return []byte(""), []byte("error"), errors.New("error")
}

// Run should assert the command parameters are correct then return a success or error depending on release name
func (r foundRunner) Run(cmd *exec.Cmd) (stdout []byte, stderr []byte, err error) {
	assert := assert.New(r.t)
	assert.Contains(cmd.Path, "helm", "command should contain helm")
	assert.Contains(cmd.Args[0], "helm", "args should contain helm")
	assert.Contains(cmd.Args[1], "status", "args should contain status")

	if cmd.Args[2] == release {
		return []byte(""), []byte(""), nil
	}
	if cmd.Args[2] == missingRelease {
		return []byte(""), []byte("not found"), errors.New("not found error")
	}
	// simulate a Helm error
	return []byte(""), []byte("error"), errors.New("helm error")
}