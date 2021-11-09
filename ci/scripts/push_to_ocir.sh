#!/usr/bin/env bash
#
# Copyright (c) 2021, Oracle and/or its affiliates.
# Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.
#

# Exit when any command fails
set -e

SCRIPT_DIR=$(cd $(dirname "$0"); pwd -P)
TOOL_SCRIPT_DIR=${SCRIPT_DIR}/../../tools/scripts

if [ -z "$WORKSPACE" ] || [ -z "$OCI_OS_NAMESPACE" ] || [ -z "$OCI_OS_BUCKET" ] || [ -z "$OCIR_SCAN_REGISTRY" ] || [ -z "$OCIR_SCAN_REPOSITORY_PATH" ] || [ -z "$BRANCH_NAME" ]; then
  echo "This script must only be called from Jenkins and requires a number of environment variables are set"
  exit 1
fi

# We should have image tar files created already in ${WORKSPACE}/tar-files
if [ ! -d "${WORKSPACE}/tar-files" ]; then
  echo "No tar files were found to push into OCIR"
  exit 1
fi

BOM_FILE=${WORKSPACE}/tar-files/verrazzano-bom.json

if [ ! -f "${BOM_FILE}" ]; then
  echo "There is no verrazzano-bom.json from this run, so we can't push anything to OCIR"
  exit 1
fi

# This assumes that the docker login has happened, and that the OCI CLI has access as well with default profile

# This also currently assumes that the repository structure has been setup. That assumption will go away
# once we add in scripting which will ensure that the OCIR repositories for the images in the BOM are created
# and setup for scanning. Most of the time these already will exist and be setup, but if there is a new image
# or images it should get things setup for them.
#
# This will likely be done by enhancing the tests/e2e/config/scripts/create_ocir_repositories.sh script
# to handle our use cases as well.

# Push the images. NOTE: If a new image was added before we do the above "ensure" step, this may have the side
# effect of pushing that image to the root compartment rather than the desired sub-compartment (OCIR behaviour),
# and that new image will not be getting scanned until that is rectified (manually)

# TODO REMOVE "-d (dry run)" once tested
sh $TOOL_SCRIPT_DIR/vz-registry-image-helper.sh -t $OCIR_SCAN_REGISTRY -r $OCIR_SCAN_REPOSITORY_PATH -l ${WORKSPACE}/tar-files -b ${BOM_FILE}

# Finally push the current verrazzano-bom.json up as the last-ocir-pushed-verrazzano-bom.json so we know those were the latest images
# pushed up. This is used when polling for results to know which images were last pushed (which results are the latest)
if [[ "${BRANCH_NAME}" == "master" || "${BRANCH_NAME}" == release-* ]]; then
  echo "Pushing verrazzano-bom.json to object storage"
  oci --region us-phoenix-1 os object put --force --namespace ${OCI_OS_NAMESPACE} -bn ${OCI_OS_BUCKET} --name ${BRANCH_NAME}/last-ocir-pushed-verrazzano-bom.json --file ${WORKSPACE}/verrazzano-bom.json
fi

# TBD: We could also save the list of repositories as well, that may save the polling job some work so it doesn't need to figure that out
# or simply just rely on the BOM there and compute from that.