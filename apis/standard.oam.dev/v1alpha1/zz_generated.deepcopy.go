// +build !ignore_autogenerated

/*
Copyright 2021 The KubeVela Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryMetric) DeepCopyInto(out *CanaryMetric) {
	*out = *in
	if in.MetricsRange != nil {
		in, out := &in.MetricsRange, &out.MetricsRange
		*out = new(MetricsExpectedRange)
		(*in).DeepCopyInto(*out)
	}
	if in.TemplateRef != nil {
		in, out := &in.TemplateRef, &out.TemplateRef
		*out = new(v1.ObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryMetric.
func (in *CanaryMetric) DeepCopy() *CanaryMetric {
	if in == nil {
		return nil
	}
	out := new(CanaryMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompRolloutStatus) DeepCopyInto(out *CompRolloutStatus) {
	*out = *in
	in.RolloutStatus.DeepCopyInto(&out.RolloutStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompRolloutStatus.
func (in *CompRolloutStatus) DeepCopy() *CompRolloutStatus {
	if in == nil {
		return nil
	}
	out := new(CompRolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsExpectedRange) DeepCopyInto(out *MetricsExpectedRange) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(intstr.IntOrString)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsExpectedRange.
func (in *MetricsExpectedRange) DeepCopy() *MetricsExpectedRange {
	if in == nil {
		return nil
	}
	out := new(MetricsExpectedRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpecWorkload) DeepCopyInto(out *PodSpecWorkload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpecWorkload.
func (in *PodSpecWorkload) DeepCopy() *PodSpecWorkload {
	if in == nil {
		return nil
	}
	out := new(PodSpecWorkload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSpecWorkload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpecWorkloadList) DeepCopyInto(out *PodSpecWorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PodSpecWorkload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpecWorkloadList.
func (in *PodSpecWorkloadList) DeepCopy() *PodSpecWorkloadList {
	if in == nil {
		return nil
	}
	out := new(PodSpecWorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PodSpecWorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpecWorkloadSpec) DeepCopyInto(out *PodSpecWorkloadSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.PodSpec.DeepCopyInto(&out.PodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpecWorkloadSpec.
func (in *PodSpecWorkloadSpec) DeepCopy() *PodSpecWorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(PodSpecWorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSpecWorkloadStatus) DeepCopyInto(out *PodSpecWorkloadStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSpecWorkloadStatus.
func (in *PodSpecWorkloadStatus) DeepCopy() *PodSpecWorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(PodSpecWorkloadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rollout) DeepCopyInto(out *Rollout) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rollout.
func (in *Rollout) DeepCopy() *Rollout {
	if in == nil {
		return nil
	}
	out := new(Rollout)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Rollout) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutBatch) DeepCopyInto(out *RolloutBatch) {
	*out = *in
	out.Replicas = in.Replicas
	if in.PodList != nil {
		in, out := &in.PodList, &out.PodList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.InstanceInterval != nil {
		in, out := &in.InstanceInterval, &out.InstanceInterval
		*out = new(int32)
		**out = **in
	}
	if in.BatchRolloutWebhooks != nil {
		in, out := &in.BatchRolloutWebhooks, &out.BatchRolloutWebhooks
		*out = make([]RolloutWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CanaryMetric != nil {
		in, out := &in.CanaryMetric, &out.CanaryMetric
		*out = make([]CanaryMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutBatch.
func (in *RolloutBatch) DeepCopy() *RolloutBatch {
	if in == nil {
		return nil
	}
	out := new(RolloutBatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutList) DeepCopyInto(out *RolloutList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Rollout, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutList.
func (in *RolloutList) DeepCopy() *RolloutList {
	if in == nil {
		return nil
	}
	out := new(RolloutList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutPlan) DeepCopyInto(out *RolloutPlan) {
	*out = *in
	if in.TargetSize != nil {
		in, out := &in.TargetSize, &out.TargetSize
		*out = new(int32)
		**out = **in
	}
	if in.NumBatches != nil {
		in, out := &in.NumBatches, &out.NumBatches
		*out = new(int32)
		**out = **in
	}
	if in.RolloutBatches != nil {
		in, out := &in.RolloutBatches, &out.RolloutBatches
		*out = make([]RolloutBatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.BatchPartition != nil {
		in, out := &in.BatchPartition, &out.BatchPartition
		*out = new(int32)
		**out = **in
	}
	if in.RolloutWebhooks != nil {
		in, out := &in.RolloutWebhooks, &out.RolloutWebhooks
		*out = make([]RolloutWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CanaryMetric != nil {
		in, out := &in.CanaryMetric, &out.CanaryMetric
		*out = make([]CanaryMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutPlan.
func (in *RolloutPlan) DeepCopy() *RolloutPlan {
	if in == nil {
		return nil
	}
	out := new(RolloutPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutSpec) DeepCopyInto(out *RolloutSpec) {
	*out = *in
	in.RolloutPlan.DeepCopyInto(&out.RolloutPlan)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutSpec.
func (in *RolloutSpec) DeepCopy() *RolloutSpec {
	if in == nil {
		return nil
	}
	out := new(RolloutSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutStatus) DeepCopyInto(out *RolloutStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStatus.
func (in *RolloutStatus) DeepCopy() *RolloutStatus {
	if in == nil {
		return nil
	}
	out := new(RolloutStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutTrait) DeepCopyInto(out *RolloutTrait) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutTrait.
func (in *RolloutTrait) DeepCopy() *RolloutTrait {
	if in == nil {
		return nil
	}
	out := new(RolloutTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutTrait) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutTraitList) DeepCopyInto(out *RolloutTraitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RolloutTrait, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutTraitList.
func (in *RolloutTraitList) DeepCopy() *RolloutTraitList {
	if in == nil {
		return nil
	}
	out := new(RolloutTraitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutTraitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutTraitSpec) DeepCopyInto(out *RolloutTraitSpec) {
	*out = *in
	out.TargetRef = in.TargetRef
	if in.SourceRef != nil {
		in, out := &in.SourceRef, &out.SourceRef
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	in.RolloutPlan.DeepCopyInto(&out.RolloutPlan)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutTraitSpec.
func (in *RolloutTraitSpec) DeepCopy() *RolloutTraitSpec {
	if in == nil {
		return nil
	}
	out := new(RolloutTraitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutWebhook) DeepCopyInto(out *RolloutWebhook) {
	*out = *in
	if in.ExpectedStatus != nil {
		in, out := &in.ExpectedStatus, &out.ExpectedStatus
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutWebhook.
func (in *RolloutWebhook) DeepCopy() *RolloutWebhook {
	if in == nil {
		return nil
	}
	out := new(RolloutWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutWebhookPayload) DeepCopyInto(out *RolloutWebhookPayload) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutWebhookPayload.
func (in *RolloutWebhookPayload) DeepCopy() *RolloutWebhookPayload {
	if in == nil {
		return nil
	}
	out := new(RolloutWebhookPayload)
	in.DeepCopyInto(out)
	return out
}
