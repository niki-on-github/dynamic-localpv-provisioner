// +build !ignore_autogenerated

/*
Copyright 2019 The OpenEBS Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorPool) DeepCopyInto(out *CStorPool) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(ResourceOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorPool.
func (in *CStorPool) DeepCopy() *CStorPool {
	if in == nil {
		return nil
	}
	out := new(CStorPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CStorVolume) DeepCopyInto(out *CStorVolume) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(ResourceOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CStorVolume.
func (in *CStorVolume) DeepCopy() *CStorVolume {
	if in == nil {
		return nil
	}
	out := new(CStorVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataItem) DeepCopyInto(out *DataItem) {
	*out = *in
	if in.Entries != nil {
		in, out := &in.Entries, &out.Entries
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataItem.
func (in *DataItem) DeepCopy() *DataItem {
	if in == nil {
		return nil
	}
	out := new(DataItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JivaVolume) DeepCopyInto(out *JivaVolume) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(ResourceOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JivaVolume.
func (in *JivaVolume) DeepCopy() *JivaVolume {
	if in == nil {
		return nil
	}
	out := new(JivaVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Options) DeepCopyInto(out *Options) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDetails) DeepCopyInto(out *ResourceDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDetails.
func (in *ResourceDetails) DeepCopy() *ResourceDetails {
	if in == nil {
		return nil
	}
	out := new(ResourceDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceOptions) DeepCopyInto(out *ResourceOptions) {
	*out = *in
	if in.IgnoreStepsOnError != nil {
		in, out := &in.IgnoreStepsOnError, &out.IgnoreStepsOnError
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceOptions.
func (in *ResourceOptions) DeepCopy() *ResourceOptions {
	if in == nil {
		return nil
	}
	out := new(ResourceOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSpec) DeepCopyInto(out *ResourceSpec) {
	*out = *in
	if in.JivaVolume != nil {
		in, out := &in.JivaVolume, &out.JivaVolume
		*out = new(JivaVolume)
		(*in).DeepCopyInto(*out)
	}
	if in.CStorVolume != nil {
		in, out := &in.CStorVolume, &out.CStorVolume
		*out = new(CStorVolume)
		(*in).DeepCopyInto(*out)
	}
	if in.CStorPool != nil {
		in, out := &in.CStorPool, &out.CStorPool
		*out = new(CStorPool)
		(*in).DeepCopyInto(*out)
	}
	if in.StoragePoolClaim != nil {
		in, out := &in.StoragePoolClaim, &out.StoragePoolClaim
		*out = new(StoragePoolClaim)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSpec.
func (in *ResourceSpec) DeepCopy() *ResourceSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceState) DeepCopyInto(out *ResourceState) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceState.
func (in *ResourceState) DeepCopy() *ResourceState {
	if in == nil {
		return nil
	}
	out := new(ResourceState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoragePoolClaim) DeepCopyInto(out *StoragePoolClaim) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(ResourceOptions)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoragePoolClaim.
func (in *StoragePoolClaim) DeepCopy() *StoragePoolClaim {
	if in == nil {
		return nil
	}
	out := new(StoragePoolClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeConfig) DeepCopyInto(out *UpgradeConfig) {
	*out = *in
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]ResourceDetails, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeConfig.
func (in *UpgradeConfig) DeepCopy() *UpgradeConfig {
	if in == nil {
		return nil
	}
	out := new(UpgradeConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeDetailedStatuses) DeepCopyInto(out *UpgradeDetailedStatuses) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.LastUpdatedTime.DeepCopyInto(&out.LastUpdatedTime)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeDetailedStatuses.
func (in *UpgradeDetailedStatuses) DeepCopy() *UpgradeDetailedStatuses {
	if in == nil {
		return nil
	}
	out := new(UpgradeDetailedStatuses)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResource) DeepCopyInto(out *UpgradeResource) {
	*out = *in
	out.ResourceDetails = in.ResourceDetails
	in.PreState.DeepCopyInto(&out.PreState)
	in.PostState.DeepCopyInto(&out.PostState)
	if in.SubResources != nil {
		in, out := &in.SubResources, &out.SubResources
		*out = make([]UpgradeSubResource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResource.
func (in *UpgradeResource) DeepCopy() *UpgradeResource {
	if in == nil {
		return nil
	}
	out := new(UpgradeResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResult) DeepCopyInto(out *UpgradeResult) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Config.DeepCopyInto(&out.Config)
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]UpgradeResultTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResult.
func (in *UpgradeResult) DeepCopy() *UpgradeResult {
	if in == nil {
		return nil
	}
	out := new(UpgradeResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeResult) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultConfig) DeepCopyInto(out *UpgradeResultConfig) {
	*out = *in
	out.ResourceDetails = in.ResourceDetails
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = make([]DataItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultConfig.
func (in *UpgradeResultConfig) DeepCopy() *UpgradeResultConfig {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultList) DeepCopyInto(out *UpgradeResultList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpgradeResult, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultList.
func (in *UpgradeResultList) DeepCopy() *UpgradeResultList {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeResultList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultStatus) DeepCopyInto(out *UpgradeResultStatus) {
	*out = *in
	in.Resource.DeepCopyInto(&out.Resource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultStatus.
func (in *UpgradeResultStatus) DeepCopy() *UpgradeResultStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeResultTask) DeepCopyInto(out *UpgradeResultTask) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeResultTask.
func (in *UpgradeResultTask) DeepCopy() *UpgradeResultTask {
	if in == nil {
		return nil
	}
	out := new(UpgradeResultTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeSubResource) DeepCopyInto(out *UpgradeSubResource) {
	*out = *in
	out.ResourceDetails = in.ResourceDetails
	in.PreState.DeepCopyInto(&out.PreState)
	in.PostState.DeepCopyInto(&out.PostState)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeSubResource.
func (in *UpgradeSubResource) DeepCopy() *UpgradeSubResource {
	if in == nil {
		return nil
	}
	out := new(UpgradeSubResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeTask) DeepCopyInto(out *UpgradeTask) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeTask.
func (in *UpgradeTask) DeepCopy() *UpgradeTask {
	if in == nil {
		return nil
	}
	out := new(UpgradeTask)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeTask) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeTaskList) DeepCopyInto(out *UpgradeTaskList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]UpgradeTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeTaskList.
func (in *UpgradeTaskList) DeepCopy() *UpgradeTaskList {
	if in == nil {
		return nil
	}
	out := new(UpgradeTaskList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *UpgradeTaskList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeTaskSpec) DeepCopyInto(out *UpgradeTaskSpec) {
	*out = *in
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = new(Options)
		**out = **in
	}
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeTaskSpec.
func (in *UpgradeTaskSpec) DeepCopy() *UpgradeTaskSpec {
	if in == nil {
		return nil
	}
	out := new(UpgradeTaskSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UpgradeTaskStatus) DeepCopyInto(out *UpgradeTaskStatus) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.CompletedTime.DeepCopyInto(&out.CompletedTime)
	if in.UpgradeDetailedStatuses != nil {
		in, out := &in.UpgradeDetailedStatuses, &out.UpgradeDetailedStatuses
		*out = make([]UpgradeDetailedStatuses, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UpgradeTaskStatus.
func (in *UpgradeTaskStatus) DeepCopy() *UpgradeTaskStatus {
	if in == nil {
		return nil
	}
	out := new(UpgradeTaskStatus)
	in.DeepCopyInto(out)
	return out
}
