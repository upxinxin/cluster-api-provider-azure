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

package vmextensions

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-11-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	"sigs.k8s.io/cluster-api-provider-azure/azure"
)

// VMExtensionSpec defines the specification for a VM or VMScaleSet extension.
type VMExtensionSpec struct {
	azure.ExtensionSpec
	ResourceGroup string
	Location      string
}

// ResourceName returns the name of the VM extension.
func (s *VMExtensionSpec) ResourceName() string {
	return s.Name
}

// ResourceGroupName returns the name of the resource group.
func (s *VMExtensionSpec) ResourceGroupName() string {
	return s.ResourceGroup
}

// OwnerResourceName returns the name of the VM that owns this VM extension.
func (s *VMExtensionSpec) OwnerResourceName() string {
	return s.VMName
}

// Parameters returns the parameters for the VM extension.
func (s *VMExtensionSpec) Parameters(ctx context.Context, existing interface{}) (interface{}, error) {
	if existing != nil {
		_, ok := existing.(compute.VirtualMachineExtension)
		if !ok {
			return nil, errors.Errorf("%T is not a compute.VirtualMachineExtension", existing)
		}

		// VM extension already exists, nothing to update.
		return nil, nil
	}

	return compute.VirtualMachineExtension{
		VirtualMachineExtensionProperties: &compute.VirtualMachineExtensionProperties{
			Publisher:          to.StringPtr(s.Publisher),
			Type:               to.StringPtr(s.Name),
			TypeHandlerVersion: to.StringPtr(s.Version),
			Settings:           s.Settings,
			ProtectedSettings:  s.ProtectedSettings,
		},
		Location: to.StringPtr(s.Location),
	}, nil
}
