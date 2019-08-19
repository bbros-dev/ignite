// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package runtime

import (
	pkgruntime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIType) DeepCopyInto(out *APIType) {
	*out = *in
	if in.TypeMeta != nil {
		in, out := &in.TypeMeta, &out.TypeMeta
		*out = new(TypeMeta)
		**out = **in
	}
	if in.ObjectMeta != nil {
		in, out := &in.ObjectMeta, &out.ObjectMeta
		*out = new(ObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIType.
func (in *APIType) DeepCopy() *APIType {
	if in == nil {
		return nil
	}
	out := new(APIType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new pkgruntime.Object.
func (in *APIType) DeepCopyObject() pkgruntime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in APITypeList) DeepCopyInto(out *APITypeList) {
	{
		in := &in
		*out = make(APITypeList, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(APIType)
				(*in).DeepCopyInto(*out)
			}
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APITypeList.
func (in APITypeList) DeepCopy() APITypeList {
	if in == nil {
		return nil
	}
	out := new(APITypeList)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectMeta) DeepCopyInto(out *ObjectMeta) {
	*out = *in
	in.Created.DeepCopyInto(&out.Created)
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectMeta.
func (in *ObjectMeta) DeepCopy() *ObjectMeta {
	if in == nil {
		return nil
	}
	out := new(ObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Time) DeepCopyInto(out *Time) {
	*out = *in
	in.Time.DeepCopyInto(&out.Time)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Time.
func (in *Time) DeepCopy() *Time {
	if in == nil {
		return nil
	}
	out := new(Time)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TypeMeta) DeepCopyInto(out *TypeMeta) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TypeMeta.
func (in *TypeMeta) DeepCopy() *TypeMeta {
	if in == nil {
		return nil
	}
	out := new(TypeMeta)
	in.DeepCopyInto(out)
	return out
}
