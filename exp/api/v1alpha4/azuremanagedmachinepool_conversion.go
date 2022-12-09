/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha4

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	infrav1exp "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this AzureManagedMachinePool to the Hub version (v1beta1).
func (src *AzureManagedMachinePool) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1exp.AzureManagedMachinePool)
	if err := Convert_v1alpha4_AzureManagedMachinePool_To_v1beta1_AzureManagedMachinePool(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &infrav1exp.AzureManagedMachinePool{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dst.Spec.Scaling = restored.Spec.Scaling
	dst.Spec.Name = restored.Spec.Name
	dst.Spec.Taints = restored.Spec.Taints
	dst.Spec.AvailabilityZones = restored.Spec.AvailabilityZones
	dst.Spec.MaxPods = restored.Spec.MaxPods
	dst.Spec.OsDiskType = restored.Spec.OsDiskType
	dst.Spec.OSType = restored.Spec.OSType
	dst.Spec.NodeLabels = restored.Spec.NodeLabels
	dst.Spec.EnableUltraSSD = restored.Spec.EnableUltraSSD
	dst.Spec.EnableNodePublicIP = restored.Spec.EnableNodePublicIP
	dst.Spec.NodePublicIPPrefixID = restored.Spec.NodePublicIPPrefixID
	dst.Spec.ScaleSetPriority = restored.Spec.ScaleSetPriority
	dst.Spec.AdditionalTags = restored.Spec.AdditionalTags
	if restored.Spec.KubeletConfig != nil {
		dst.Spec.KubeletConfig = restored.Spec.KubeletConfig
	}

	dst.Status.LongRunningOperationStates = restored.Status.LongRunningOperationStates
	dst.Status.Conditions = restored.Status.Conditions

	return nil
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *AzureManagedMachinePool) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1exp.AzureManagedMachinePool)
	if err := Convert_v1beta1_AzureManagedMachinePool_To_v1alpha4_AzureManagedMachinePool(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	return utilconversion.MarshalData(src, dst)
}

// Convert_v1beta1_AzureManagedMachinePoolSpec_To_v1alpha4_AzureManagedMachinePoolSpec is an autogenerated conversion function.
func Convert_v1beta1_AzureManagedMachinePoolSpec_To_v1alpha4_AzureManagedMachinePoolSpec(in *infrav1exp.AzureManagedMachinePoolSpec, out *AzureManagedMachinePoolSpec, s apiconversion.Scope) error {
	return autoConvert_v1beta1_AzureManagedMachinePoolSpec_To_v1alpha4_AzureManagedMachinePoolSpec(in, out, s)
}

// Convert_v1beta1_AzureManagedMachinePoolStatus_To_v1alpha4_AzureManagedMachinePoolStatus is an autogenerated conversion function.
func Convert_v1beta1_AzureManagedMachinePoolStatus_To_v1alpha4_AzureManagedMachinePoolStatus(in *infrav1exp.AzureManagedMachinePoolStatus, out *AzureManagedMachinePoolStatus, s apiconversion.Scope) error {
	return autoConvert_v1beta1_AzureManagedMachinePoolStatus_To_v1alpha4_AzureManagedMachinePoolStatus(in, out, s)
}

// ConvertTo converts this AzureManagedMachinePoolList to the Hub version (v1beta1).
func (src *AzureManagedMachinePoolList) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*infrav1exp.AzureManagedMachinePoolList)
	return Convert_v1alpha4_AzureManagedMachinePoolList_To_v1beta1_AzureManagedMachinePoolList(src, dst, nil)
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *AzureManagedMachinePoolList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*infrav1exp.AzureManagedMachinePoolList)
	return Convert_v1beta1_AzureManagedMachinePoolList_To_v1alpha4_AzureManagedMachinePoolList(src, dst, nil)
}
