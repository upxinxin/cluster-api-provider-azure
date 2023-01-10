//go:build e2e
// +build e2e

/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"context"
	"fmt"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	helmVals "helm.sh/helm/v3/pkg/cli/values"
	"sigs.k8s.io/cluster-api/test/framework/clusterctl"
)

const (
	cloudProviderAzureHelmRepoURL     = "https://raw.githubusercontent.com/kubernetes-sigs/cloud-provider-azure/master/helm/repo"
	cloudProviderAzureChartName       = "cloud-provider-azure"
	cloudProviderAzureHelmReleaseName = "cloud-provider-azure-oot"
	azureDiskCSIDriverHelmRepoURL     = "https://raw.githubusercontent.com/kubernetes-sigs/azuredisk-csi-driver/master/charts"
	azureDiskCSIDriverChartName       = "azuredisk-csi-driver"
	azureDiskCSIDriverHelmReleaseName = "azuredisk-csi-driver-oot"
	defaultNamespace                  = "default"
)

// InstallCalicoAndCloudProviderAzureHelmChart installs the official cloud-provider-azure helm chart
// and validates that expected pods exist and are Ready.
func InstallCalicoAndCloudProviderAzureHelmChart(ctx context.Context, input clusterctl.ApplyClusterTemplateAndWaitInput, cidrBlocks []string, hasWindows bool) {
	specName := "cloud-provider-azure-install"
	By("Installing cloud-provider-azure components via helm")
	options := &helmVals.Options{
		Values: []string{fmt.Sprintf("infra.clusterName=%s", input.ConfigCluster.ClusterName), fmt.Sprintf("cloudControllerManager.clusterCIDR=%s", strings.Join(cidrBlocks, `,`))},
	}
	InstallHelmChart(ctx, input, defaultNamespace, cloudProviderAzureHelmRepoURL, cloudProviderAzureChartName, cloudProviderAzureHelmReleaseName, options)

	// Install Calico CNI Helm Chart. We do this before waiting for the pods to be ready because there is a co-dependency between CNI (nodes ready) and cloud-provider being initialized.
	InstallCalicoHelmChart(ctx, input, cidrBlocks, hasWindows)

	clusterProxy := input.ClusterProxy.GetWorkloadCluster(ctx, input.ConfigCluster.Namespace, input.ConfigCluster.ClusterName)
	By("Waiting for Ready cloud-controller-manager deployment pods")
	for _, d := range []string{"cloud-controller-manager"} {
		waitInput := GetWaitForDeploymentsAvailableInput(ctx, clusterProxy, d, kubesystem, specName)
		WaitForDeploymentsAvailable(ctx, waitInput, e2eConfig.GetIntervals(specName, "wait-deployment")...)
	}
	By("Waiting for Ready cloud-node-manager daemonset pods")
	for _, ds := range []string{"cloud-node-manager", "cloud-node-manager-windows"} {
		waitInput := GetWaitForDaemonsetAvailableInput(ctx, clusterProxy, ds, kubesystem, specName)
		WaitForDaemonset(ctx, waitInput, e2eConfig.GetIntervals(specName, "wait-daemonset")...)
	}
}

// InstallAzureDiskCSIDriverHelmChart installs the official azure-disk CSI driver helm chart
func InstallAzureDiskCSIDriverHelmChart(ctx context.Context, input clusterctl.ApplyClusterTemplateAndWaitInput, hasWindows bool) {
	specName := "azuredisk-csi-drivers-install"
	By("Installing azure-disk CSI driver components via helm")
	options := &helmVals.Options{
		Values: []string{"controller.replicas=1", "controller.runOnControlPlane=true"},
	}
	// TODO: make this always true once HostProcessContainers are on for all supported k8s versions.
	if hasWindows {
		options.Values = append(options.Values, "windows.useHostProcessContainers=true")
	}
	InstallHelmChart(ctx, input, kubesystem, azureDiskCSIDriverHelmRepoURL, azureDiskCSIDriverChartName, azureDiskCSIDriverHelmReleaseName, options)
	clusterProxy := input.ClusterProxy.GetWorkloadCluster(ctx, input.ConfigCluster.Namespace, input.ConfigCluster.ClusterName)
	By("Waiting for Ready csi-azuredisk-controller deployment pods")
	for _, d := range []string{"csi-azuredisk-controller"} {
		waitInput := GetWaitForDeploymentsAvailableInput(ctx, clusterProxy, d, kubesystem, specName)
		WaitForDeploymentsAvailable(ctx, waitInput, e2eConfig.GetIntervals(specName, "wait-deployment")...)
	}
	By("Waiting for Running azure-disk-csi node pods")
	for _, ds := range []string{"csi-azuredisk-node", "csi-azuredisk-node-win"} {
		waitInput := GetWaitForDaemonsetAvailableInput(ctx, clusterProxy, ds, kubesystem, specName)
		WaitForDaemonset(ctx, waitInput, e2eConfig.GetIntervals(specName, "wait-daemonset")...)
	}
}
