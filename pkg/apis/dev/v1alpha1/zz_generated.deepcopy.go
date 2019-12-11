// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentPoolSpec) DeepCopyInto(out *AgentPoolSpec) {
	*out = *in
	if in.PoolSpec != nil {
		in, out := &in.PoolSpec, &out.PoolSpec
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentPoolSpec.
func (in *AgentPoolSpec) DeepCopy() *AgentPoolSpec {
	if in == nil {
		return nil
	}
	out := new(AgentPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePipelinePool) DeepCopyInto(out *AzurePipelinePool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePipelinePool.
func (in *AzurePipelinePool) DeepCopy() *AzurePipelinePool {
	if in == nil {
		return nil
	}
	out := new(AzurePipelinePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzurePipelinePool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePipelinePoolList) DeepCopyInto(out *AzurePipelinePoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzurePipelinePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePipelinePoolList.
func (in *AzurePipelinePoolList) DeepCopy() *AzurePipelinePoolList {
	if in == nil {
		return nil
	}
	out := new(AzurePipelinePoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzurePipelinePoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePipelinePoolSpec) DeepCopyInto(out *AzurePipelinePoolSpec) {
	*out = *in
	if in.AgentPools != nil {
		in, out := &in.AgentPools, &out.AgentPools
		*out = make([]AgentPoolSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePipelinePoolSpec.
func (in *AzurePipelinePoolSpec) DeepCopy() *AzurePipelinePoolSpec {
	if in == nil {
		return nil
	}
	out := new(AzurePipelinePoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzurePipelinePoolStatus) DeepCopyInto(out *AzurePipelinePoolStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzurePipelinePoolStatus.
func (in *AzurePipelinePoolStatus) DeepCopy() *AzurePipelinePoolStatus {
	if in == nil {
		return nil
	}
	out := new(AzurePipelinePoolStatus)
	in.DeepCopyInto(out)
	return out
}
