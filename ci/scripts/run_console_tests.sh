#!/usr/bin/env bash
#
# Copyright (c) 2021, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.
#

if [ -z "$KUBECONFIG" ] || [ -z "$GO_REPO_PATH" ] || [ -z "${DOCKER_CREDS_USR}" ] ||
 [ -z "${DOCKER_CREDS_PSW}" ] || [ -z "$CONSOLE_REPO_BRANCH" ]; then
  echo "This script must only be called from Jenkins and requires a number of environment variables are set"
  exit 1
fi

# Temporarily clone the console repo until it is moved to the verrazzano repo
cd ${GO_REPO_PATH}
git clone https://${DOCKER_CREDS_USR}:${DOCKER_CREDS_PSW}@github.com/verrazzano/console.git
cd console
git checkout ${CONSOLE_REPO_BRANCH}

# Run the basic UI tests, and if they fail make sure to exit with a fail status
make run-ui-tests || exit 1

# Run the application page UI tests if specified
if [ "true" == "${RUN_APP_TESTS}" ]; then
  echo "Running Application Page UI tests"
  make run-app-page-test
fi