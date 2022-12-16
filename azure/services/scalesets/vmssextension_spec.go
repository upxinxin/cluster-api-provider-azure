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

package scalesets

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-11-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-azure/azure"
)

// VMSSExtensionSpec defines the specification for a VM or VMScaleSet extension.
type VMSSExtensionSpec struct {
	azure.ExtensionSpec
	ResourceGroup string
}

// ResourceName returns the name of the VMSS extension.
func (s *VMSSExtensionSpec) ResourceName() string {
	return s.Name
}

// ResourceGroupName returns the name of the resource group.
func (s *VMSSExtensionSpec) ResourceGroupName() string {
	return s.ResourceGroup
}

// OwnerResourceName returns the name of the VMSS that owns this VMSS extension.
func (s *VMSSExtensionSpec) OwnerResourceName() string {
	return s.VMName
}

// Parameters returns the parameters for the VMSS extension.
func (s *VMSSExtensionSpec) Parameters(ctx context.Context, existing interface{}) (interface{}, error) {
	if existing != nil {
		_, ok := existing.(compute.VirtualMachineScaleSetExtension)
		if !ok {
			return nil, errors.Errorf("%T is not a compute.VirtualMachineScaleSetExtension", existing)
		}

		// VMSS extension already exists, nothing to update.
		return nil, nil
	}

	return compute.VirtualMachineScaleSetExtension{
		Name: to.StringPtr(s.Name),
		VirtualMachineScaleSetExtensionProperties: &compute.VirtualMachineScaleSetExtensionProperties{
			Publisher:          to.StringPtr(s.Publisher),
			Type:               to.StringPtr(s.Name),
			TypeHandlerVersion: to.StringPtr(s.Version),
			Settings:           s.Settings,
			ProtectedSettings:  s.ProtectedSettings,
		},
	}, nil
}
