// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package api

import (
	net "net"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileMapping) DeepCopyInto(out *FileMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileMapping.
func (in *FileMapping) DeepCopy() *FileMapping {
	if in == nil {
		return nil
	}
	out := new(FileMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageClaim) DeepCopyInto(out *ImageClaim) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageClaim.
func (in *ImageClaim) DeepCopy() *ImageClaim {
	if in == nil {
		return nil
	}
	out := new(ImageClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSource) DeepCopyInto(out *ImageSource) {
	*out = *in
	out.Size = in.Size
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSource.
func (in *ImageSource) DeepCopy() *ImageSource {
	if in == nil {
		return nil
	}
	out := new(ImageSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	out.Source = in.Source
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kernel) DeepCopyInto(out *Kernel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kernel.
func (in *Kernel) DeepCopy() *Kernel {
	if in == nil {
		return nil
	}
	out := new(Kernel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kernel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelClaim) DeepCopyInto(out *KernelClaim) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelClaim.
func (in *KernelClaim) DeepCopy() *KernelClaim {
	if in == nil {
		return nil
	}
	out := new(KernelClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelSpec) DeepCopyInto(out *KernelSpec) {
	*out = *in
	out.Source = in.Source
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelSpec.
func (in *KernelSpec) DeepCopy() *KernelSpec {
	if in == nil {
		return nil
	}
	out := new(KernelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pool) DeepCopyInto(out *Pool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pool.
func (in *Pool) DeepCopy() *Pool {
	if in == nil {
		return nil
	}
	out := new(Pool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolDevice) DeepCopyInto(out *PoolDevice) {
	*out = *in
	out.Size = in.Size
	out.Parent = in.Parent
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolDevice.
func (in *PoolDevice) DeepCopy() *PoolDevice {
	if in == nil {
		return nil
	}
	out := new(PoolDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpec) DeepCopyInto(out *PoolSpec) {
	*out = *in
	out.MetadataSize = in.MetadataSize
	out.DataSize = in.DataSize
	out.AllocationSize = in.AllocationSize
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpec.
func (in *PoolSpec) DeepCopy() *PoolSpec {
	if in == nil {
		return nil
	}
	out := new(PoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolStatus) DeepCopyInto(out *PoolStatus) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*PoolDevice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PoolDevice)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolStatus.
func (in *PoolStatus) DeepCopy() *PoolStatus {
	if in == nil {
		return nil
	}
	out := new(PoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortMapping) DeepCopyInto(out *PortMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortMapping.
func (in *PortMapping) DeepCopy() *PortMapping {
	if in == nil {
		return nil
	}
	out := new(PortMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSH) DeepCopyInto(out *SSH) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSH.
func (in *SSH) DeepCopy() *SSH {
	if in == nil {
		return nil
	}
	out := new(SSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VM) DeepCopyInto(out *VM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VM.
func (in *VM) DeepCopy() *VM {
	if in == nil {
		return nil
	}
	out := new(VM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSpec) DeepCopyInto(out *VMSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageClaim)
		**out = **in
	}
	if in.Kernel != nil {
		in, out := &in.Kernel, &out.Kernel
		*out = new(KernelClaim)
		**out = **in
	}
	out.Memory = in.Memory
	out.DiskSize = in.DiskSize
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]PortMapping, len(*in))
		copy(*out, *in)
	}
	if in.CopyFiles != nil {
		in, out := &in.CopyFiles, &out.CopyFiles
		*out = make([]FileMapping, len(*in))
		copy(*out, *in)
	}
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		*out = new(SSH)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSpec.
func (in *VMSpec) DeepCopy() *VMSpec {
	if in == nil {
		return nil
	}
	out := new(VMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMStatus) DeepCopyInto(out *VMStatus) {
	*out = *in
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make([]net.IP, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(net.IP, len(*in))
				copy(*out, *in)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMStatus.
func (in *VMStatus) DeepCopy() *VMStatus {
	if in == nil {
		return nil
	}
	out := new(VMStatus)
	in.DeepCopyInto(out)
	return out
}
