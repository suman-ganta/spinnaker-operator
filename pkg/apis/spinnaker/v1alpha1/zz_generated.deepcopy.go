// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerDeploymentStatus) DeepCopyInto(out *SpinnakerDeploymentStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerDeploymentStatus.
func (in *SpinnakerDeploymentStatus) DeepCopy() *SpinnakerDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(SpinnakerDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerFileSource) DeepCopyInto(out *SpinnakerFileSource) {
	*out = *in
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(SpinnakerFileSourceReference)
		**out = **in
	}
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		*out = new(SpinnakerFileSourceReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerFileSource.
func (in *SpinnakerFileSource) DeepCopy() *SpinnakerFileSource {
	if in == nil {
		return nil
	}
	out := new(SpinnakerFileSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerFileSourceReference) DeepCopyInto(out *SpinnakerFileSourceReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerFileSourceReference.
func (in *SpinnakerFileSourceReference) DeepCopy() *SpinnakerFileSourceReference {
	if in == nil {
		return nil
	}
	out := new(SpinnakerFileSourceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerService) DeepCopyInto(out *SpinnakerService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerService.
func (in *SpinnakerService) DeepCopy() *SpinnakerService {
	if in == nil {
		return nil
	}
	out := new(SpinnakerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpinnakerService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerServiceList) DeepCopyInto(out *SpinnakerServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpinnakerService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerServiceList.
func (in *SpinnakerServiceList) DeepCopy() *SpinnakerServiceList {
	if in == nil {
		return nil
	}
	out := new(SpinnakerServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpinnakerServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerServiceSpec) DeepCopyInto(out *SpinnakerServiceSpec) {
	*out = *in
	in.HalConfig.DeepCopyInto(&out.HalConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerServiceSpec.
func (in *SpinnakerServiceSpec) DeepCopy() *SpinnakerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(SpinnakerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerServiceStatus) DeepCopyInto(out *SpinnakerServiceStatus) {
	*out = *in
	in.LastConfigurationTime.DeepCopyInto(&out.LastConfigurationTime)
	in.HalConfig.DeepCopyInto(&out.HalConfig)
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = make([]SpinnakerDeploymentStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerServiceStatus.
func (in *SpinnakerServiceStatus) DeepCopy() *SpinnakerServiceStatus {
	if in == nil {
		return nil
	}
	out := new(SpinnakerServiceStatus)
	in.DeepCopyInto(out)
	return out
}
