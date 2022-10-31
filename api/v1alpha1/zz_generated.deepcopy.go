//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 The Kruise Authors.

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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowRunTime) DeepCopyInto(out *AllowRunTime) {
	*out = *in
	if in.TimeZone != nil {
		in, out := &in.TimeZone, &out.TimeZone
		*out = new(TimeZone)
		**out = **in
	}
	if in.TimeSlices != nil {
		in, out := &in.TimeSlices, &out.TimeSlices
		*out = make([]TimeSlice, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowRunTime.
func (in *AllowRunTime) DeepCopy() *AllowRunTime {
	if in == nil {
		return nil
	}
	out := new(AllowRunTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchRelease) DeepCopyInto(out *BatchRelease) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchRelease.
func (in *BatchRelease) DeepCopy() *BatchRelease {
	if in == nil {
		return nil
	}
	out := new(BatchRelease)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchRelease) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseCanaryStatus) DeepCopyInto(out *BatchReleaseCanaryStatus) {
	*out = *in
	if in.BatchReadyTime != nil {
		in, out := &in.BatchReadyTime, &out.BatchReadyTime
		*out = (*in).DeepCopy()
	}
	if in.NoNeedUpdateReplicas != nil {
		in, out := &in.NoNeedUpdateReplicas, &out.NoNeedUpdateReplicas
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseCanaryStatus.
func (in *BatchReleaseCanaryStatus) DeepCopy() *BatchReleaseCanaryStatus {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseCanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseList) DeepCopyInto(out *BatchReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BatchRelease, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseList.
func (in *BatchReleaseList) DeepCopy() *BatchReleaseList {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BatchReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseSpec) DeepCopyInto(out *BatchReleaseSpec) {
	*out = *in
	in.TargetRef.DeepCopyInto(&out.TargetRef)
	in.ReleasePlan.DeepCopyInto(&out.ReleasePlan)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseSpec.
func (in *BatchReleaseSpec) DeepCopy() *BatchReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BatchReleaseStatus) DeepCopyInto(out *BatchReleaseStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RolloutCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CanaryStatus.DeepCopyInto(&out.CanaryStatus)
	if in.CollisionCount != nil {
		in, out := &in.CollisionCount, &out.CollisionCount
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BatchReleaseStatus.
func (in *BatchReleaseStatus) DeepCopy() *BatchReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(BatchReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
	*out = *in
	if in.LastUpdateTime != nil {
		in, out := &in.LastUpdateTime, &out.LastUpdateTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStep) DeepCopyInto(out *CanaryStep) {
	*out = *in
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(int32)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(intstr.IntOrString)
		**out = **in
	}
	in.Pause.DeepCopyInto(&out.Pause)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStep.
func (in *CanaryStep) DeepCopy() *CanaryStep {
	if in == nil {
		return nil
	}
	out := new(CanaryStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStepInfo) DeepCopyInto(out *CanaryStepInfo) {
	*out = *in
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]Pod, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStepInfo.
func (in *CanaryStepInfo) DeepCopy() *CanaryStepInfo {
	if in == nil {
		return nil
	}
	out := new(CanaryStepInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStrategy) DeepCopyInto(out *CanaryStrategy) {
	*out = *in
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]CanaryStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TrafficRoutings != nil {
		in, out := &in.TrafficRoutings, &out.TrafficRoutings
		*out = make([]*TrafficRouting, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(TrafficRouting)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStrategy.
func (in *CanaryStrategy) DeepCopy() *CanaryStrategy {
	if in == nil {
		return nil
	}
	out := new(CanaryStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatewayTrafficRouting) DeepCopyInto(out *GatewayTrafficRouting) {
	*out = *in
	if in.HTTPRouteName != nil {
		in, out := &in.HTTPRouteName, &out.HTTPRouteName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatewayTrafficRouting.
func (in *GatewayTrafficRouting) DeepCopy() *GatewayTrafficRouting {
	if in == nil {
		return nil
	}
	out := new(GatewayTrafficRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPRouteInfo) DeepCopyInto(out *HTTPRouteInfo) {
	*out = *in
	in.NameAndSpecData.DeepCopyInto(&out.NameAndSpecData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPRouteInfo.
func (in *HTTPRouteInfo) DeepCopy() *HTTPRouteInfo {
	if in == nil {
		return nil
	}
	out := new(HTTPRouteInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressInfo) DeepCopyInto(out *IngressInfo) {
	*out = *in
	in.NameAndSpecData.DeepCopyInto(&out.NameAndSpecData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressInfo.
func (in *IngressInfo) DeepCopy() *IngressInfo {
	if in == nil {
		return nil
	}
	out := new(IngressInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressTrafficRouting) DeepCopyInto(out *IngressTrafficRouting) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressTrafficRouting.
func (in *IngressTrafficRouting) DeepCopy() *IngressTrafficRouting {
	if in == nil {
		return nil
	}
	out := new(IngressTrafficRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameAndSpecData) DeepCopyInto(out *NameAndSpecData) {
	*out = *in
	in.Data.DeepCopyInto(&out.Data)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameAndSpecData.
func (in *NameAndSpecData) DeepCopy() *NameAndSpecData {
	if in == nil {
		return nil
	}
	out := new(NameAndSpecData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
	if in.WorkloadRef != nil {
		in, out := &in.WorkloadRef, &out.WorkloadRef
		*out = new(WorkloadRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseBatch) DeepCopyInto(out *ReleaseBatch) {
	*out = *in
	out.CanaryReplicas = in.CanaryReplicas
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseBatch.
func (in *ReleaseBatch) DeepCopy() *ReleaseBatch {
	if in == nil {
		return nil
	}
	out := new(ReleaseBatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleasePlan) DeepCopyInto(out *ReleasePlan) {
	*out = *in
	if in.Batches != nil {
		in, out := &in.Batches, &out.Batches
		*out = make([]ReleaseBatch, len(*in))
		copy(*out, *in)
	}
	if in.BatchPartition != nil {
		in, out := &in.BatchPartition, &out.BatchPartition
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleasePlan.
func (in *ReleasePlan) DeepCopy() *ReleasePlan {
	if in == nil {
		return nil
	}
	out := new(ReleasePlan)
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
func (in *RolloutCondition) DeepCopyInto(out *RolloutCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutCondition.
func (in *RolloutCondition) DeepCopy() *RolloutCondition {
	if in == nil {
		return nil
	}
	out := new(RolloutCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutHistory) DeepCopyInto(out *RolloutHistory) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutHistory.
func (in *RolloutHistory) DeepCopy() *RolloutHistory {
	if in == nil {
		return nil
	}
	out := new(RolloutHistory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutHistory) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutHistoryList) DeepCopyInto(out *RolloutHistoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RolloutHistory, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutHistoryList.
func (in *RolloutHistoryList) DeepCopy() *RolloutHistoryList {
	if in == nil {
		return nil
	}
	out := new(RolloutHistoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RolloutHistoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutHistorySpec) DeepCopyInto(out *RolloutHistorySpec) {
	*out = *in
	in.Rollout.DeepCopyInto(&out.Rollout)
	in.Workload.DeepCopyInto(&out.Workload)
	in.Service.DeepCopyInto(&out.Service)
	in.TrafficRouting.DeepCopyInto(&out.TrafficRouting)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutHistorySpec.
func (in *RolloutHistorySpec) DeepCopy() *RolloutHistorySpec {
	if in == nil {
		return nil
	}
	out := new(RolloutHistorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutHistoryStatus) DeepCopyInto(out *RolloutHistoryStatus) {
	*out = *in
	if in.CanarySteps != nil {
		in, out := &in.CanarySteps, &out.CanarySteps
		*out = make([]CanaryStepInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutHistoryStatus.
func (in *RolloutHistoryStatus) DeepCopy() *RolloutHistoryStatus {
	if in == nil {
		return nil
	}
	out := new(RolloutHistoryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutInfo) DeepCopyInto(out *RolloutInfo) {
	*out = *in
	in.NameAndSpecData.DeepCopyInto(&out.NameAndSpecData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutInfo.
func (in *RolloutInfo) DeepCopy() *RolloutInfo {
	if in == nil {
		return nil
	}
	out := new(RolloutInfo)
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
func (in *RolloutPause) DeepCopyInto(out *RolloutPause) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutPause.
func (in *RolloutPause) DeepCopy() *RolloutPause {
	if in == nil {
		return nil
	}
	out := new(RolloutPause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RolloutSpec) DeepCopyInto(out *RolloutSpec) {
	*out = *in
	in.ObjectRef.DeepCopyInto(&out.ObjectRef)
	in.AllowRunTime.DeepCopyInto(&out.AllowRunTime)
	in.Strategy.DeepCopyInto(&out.Strategy)
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
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]RolloutCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CanaryStatus != nil {
		in, out := &in.CanaryStatus, &out.CanaryStatus
		*out = new(CanaryStatus)
		(*in).DeepCopyInto(*out)
	}
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
func (in *RolloutStrategy) DeepCopyInto(out *RolloutStrategy) {
	*out = *in
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(CanaryStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RolloutStrategy.
func (in *RolloutStrategy) DeepCopy() *RolloutStrategy {
	if in == nil {
		return nil
	}
	out := new(RolloutStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceInfo) DeepCopyInto(out *ServiceInfo) {
	*out = *in
	in.NameAndSpecData.DeepCopyInto(&out.NameAndSpecData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceInfo.
func (in *ServiceInfo) DeepCopy() *ServiceInfo {
	if in == nil {
		return nil
	}
	out := new(ServiceInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeSlice) DeepCopyInto(out *TimeSlice) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeSlice.
func (in *TimeSlice) DeepCopy() *TimeSlice {
	if in == nil {
		return nil
	}
	out := new(TimeSlice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeZone) DeepCopyInto(out *TimeZone) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeZone.
func (in *TimeZone) DeepCopy() *TimeZone {
	if in == nil {
		return nil
	}
	out := new(TimeZone)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficRouting) DeepCopyInto(out *TrafficRouting) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressTrafficRouting)
		**out = **in
	}
	if in.Gateway != nil {
		in, out := &in.Gateway, &out.Gateway
		*out = new(GatewayTrafficRouting)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficRouting.
func (in *TrafficRouting) DeepCopy() *TrafficRouting {
	if in == nil {
		return nil
	}
	out := new(TrafficRouting)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TrafficRoutingInfo) DeepCopyInto(out *TrafficRoutingInfo) {
	*out = *in
	if in.Ingress != nil {
		in, out := &in.Ingress, &out.Ingress
		*out = new(IngressInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPRoute != nil {
		in, out := &in.HTTPRoute, &out.HTTPRoute
		*out = new(HTTPRouteInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TrafficRoutingInfo.
func (in *TrafficRoutingInfo) DeepCopy() *TrafficRoutingInfo {
	if in == nil {
		return nil
	}
	out := new(TrafficRoutingInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadInfo) DeepCopyInto(out *WorkloadInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.NameAndSpecData.DeepCopyInto(&out.NameAndSpecData)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadInfo.
func (in *WorkloadInfo) DeepCopy() *WorkloadInfo {
	if in == nil {
		return nil
	}
	out := new(WorkloadInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRef) DeepCopyInto(out *WorkloadRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRef.
func (in *WorkloadRef) DeepCopy() *WorkloadRef {
	if in == nil {
		return nil
	}
	out := new(WorkloadRef)
	in.DeepCopyInto(out)
	return out
}
