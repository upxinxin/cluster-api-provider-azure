#!/bin/bash
# Copyright 2020 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
# shellcheck source=hack/util.sh
source "${REPO_ROOT}/hack/util.sh"

# Verify the required Environment Variables are present.
capz::util::ensure_azure_envs

make envsubst

export REGISTRY="${REGISTRY:-registry.local/fake}"

# Cluster settings.
export CLUSTER_NAME="${CLUSTER_NAME:-capz-test}"
export WORKER_MACHINE_COUNT=2
export CONTROL_PLANE_MACHINE_COUNT=1
export KUBERNETES_VERSION="v1.25.0"

# Azure settings.
export AZURE_LOCATION="eastus2euap"
export AZURE_EXTENDEDLOCATION_TYPE="EdgeZone"
export AZURE_EXTENDEDLOCATION_NAME="microsoftrrdclab3"
export AZURE_RESOURCE_GROUP="${CLUSTER_NAME}"

# Machine settings.
export AZURE_CONTROL_PLANE_MACHINE_TYPE="Standard_D2s_v3"
export AZURE_NODE_MACHINE_TYPE="Standard_D2s_v3"

# identity secret settings.
export AZURE_CLUSTER_IDENTITY_SECRET_NAME="cluster-identity-secret"
export CLUSTER_IDENTITY_NAME=${CLUSTER_IDENTITY_NAME:="cluster-identity"}
export AZURE_CLUSTER_IDENTITY_SECRET_NAMESPACE="default"

# Generate SSH key.
capz::util::generate_ssh_key

echo "================ DOCKER BUILD ==============="
PULL_POLICY=IfNotPresent make modules docker-build

echo "================ MAKE CLEAN ==============="
make clean

echo "================ KIND RESET ==============="
make kind-reset

echo "================ INSTALL TOOLS ==============="
make install-tools

echo "================ CREATE CLUSTER ==============="
EXP_CLUSTER_RESOURCE_SET=true EXP_AKS=true EXP_MACHINE_POOL=true EXP_EDGEZONE=true make create-management-cluster
