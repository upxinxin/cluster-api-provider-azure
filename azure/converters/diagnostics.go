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

package converters

import (
	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2021-11-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	infrav1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
)

// GetDiagnosticsProfile converts a CAPZ Diagnostics option to a Azure SDK Diagnostics Profile.
func GetDiagnosticsProfile(diagnostics *infrav1.Diagnostics) *compute.DiagnosticsProfile {
	if diagnostics != nil && diagnostics.Boot != nil {
		switch diagnostics.Boot.StorageAccountType {
		case infrav1.DisabledDiagnosticsStorage:
			return &compute.DiagnosticsProfile{
				BootDiagnostics: &compute.BootDiagnostics{
					Enabled: to.BoolPtr(false),
				},
			}
		case infrav1.ManagedDiagnosticsStorage:
			return &compute.DiagnosticsProfile{
				BootDiagnostics: &compute.BootDiagnostics{
					Enabled: to.BoolPtr(true),
				},
			}
		case infrav1.UserManagedDiagnosticsStorage:
			return &compute.DiagnosticsProfile{
				BootDiagnostics: &compute.BootDiagnostics{
					Enabled:    to.BoolPtr(true),
					StorageURI: &diagnostics.Boot.UserManaged.StorageAccountURI,
				},
			}
		}
	}

	return nil
}
