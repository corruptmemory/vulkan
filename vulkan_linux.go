// +build linux,!android

package vulkan

/*
#cgo LDFLAGS: -ldl

#include "vulkan/vulkan.h"
#include <stdlib.h>
#include "vk_wrapper.h"
#include "vk_bridge.h"

static void write_bits(VkAccelerationStructureInstanceKHR *in,
                       uint32_t instanceCustomIndex,
                       uint32_t mask,
                       uint32_t instanceShaderBindingTableRecordOffset,
                       VkGeometryInstanceFlagsKHR flags)
{
	in->instanceCustomIndex = instanceCustomIndex;
	in->mask = mask;
	in->instanceShaderBindingTableRecordOffset = instanceShaderBindingTableRecordOffset;
	in->flags = flags;
}


static uint32_t get_instanceCustomIndex(VkAccelerationStructureInstanceKHR *in) {
	return in->instanceCustomIndex;
}

static uint32_t get_mask(VkAccelerationStructureInstanceKHR *in) {
	return in->mask;
}

static uint32_t get_instanceShaderBindingTableRecordOffset(VkAccelerationStructureInstanceKHR *in) {
	return in->instanceShaderBindingTableRecordOffset;
}

static uint32_t get_flags(VkAccelerationStructureInstanceKHR *in) {
	return in->flags;
}

*/
import "C"
import (
	"fmt"
	"unsafe"
)

// AccelerationStructureInstance as declared in https://www.khronos.org/registry/vulkan/specs/1.2-khr-extensions/html/vkspec.html#VkAccelerationStructureInstanceKHR
type AccelerationStructureInstance struct {
	Transform                              TransformMatrix
	InstanceCustomIndex                    uint32
	Mask                                   uint32
	InstanceShaderBindingTableRecordOffset uint32
	Flags                                  GeometryInstanceFlags
	AccelerationStructureReference         uint32
	refd62e4d1f                            *C.VkAccelerationStructureInstanceKHR
	allocsd62e4d1f                         interface{}
}

// allocAccelerationStructureInstanceMemory allocates memory for type C.VkAccelerationStructureInstanceKHR in C.
// The caller is responsible for freeing the this memory via C.free.
func allocAccelerationStructureInstanceMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfAccelerationStructureInstanceValue))
	if mem == nil {
		panic(fmt.Sprintln("memory alloc error: ", err))
	}
	return mem
}

const sizeOfAccelerationStructureInstanceValue = unsafe.Sizeof([1]C.VkAccelerationStructureInstanceKHR{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AccelerationStructureInstance) Ref() *C.VkAccelerationStructureInstanceKHR {
	if x == nil {
		return nil
	}
	return x.refd62e4d1f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AccelerationStructureInstance) Free() {
	if x != nil && x.allocsd62e4d1f != nil {
		x.allocsd62e4d1f.(*cgoAllocMap).Free()
		x.refd62e4d1f = nil
	}
}

// NewAccelerationStructureInstanceRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAccelerationStructureInstanceRef(ref unsafe.Pointer) *AccelerationStructureInstance {
	if ref == nil {
		return nil
	}
	obj := new(AccelerationStructureInstance)
	obj.refd62e4d1f = (*C.VkAccelerationStructureInstanceKHR)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AccelerationStructureInstance) PassRef() (*C.VkAccelerationStructureInstanceKHR, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd62e4d1f != nil {
		return x.refd62e4d1f, nil
	}
	memd62e4d1f := allocAccelerationStructureInstanceMemory(1)
	refd62e4d1f := (*C.VkAccelerationStructureInstanceKHR)(memd62e4d1f)
	allocsd62e4d1f := new(cgoAllocMap)
	allocsd62e4d1f.Add(memd62e4d1f)

	var ctransform_allocs *cgoAllocMap
	refd62e4d1f.transform, ctransform_allocs = x.Transform.PassValue()
	allocsd62e4d1f.Borrow(ctransform_allocs)

	C.write_bits(refd62e4d1f,
		(C.uint32_t)(x.InstanceCustomIndex),
		(C.uint32_t)(x.Mask),
		(C.uint32_t)(x.InstanceShaderBindingTableRecordOffset),
		(C.VkGeometryInstanceFlagsKHR)(x.Flags))

	var caccelerationStructureReference_allocs *cgoAllocMap
	refd62e4d1f.accelerationStructureReference, caccelerationStructureReference_allocs = (C.uint64_t)(x.AccelerationStructureReference), cgoAllocsUnknown
	allocsd62e4d1f.Borrow(caccelerationStructureReference_allocs)

	x.refd62e4d1f = refd62e4d1f
	x.allocsd62e4d1f = allocsd62e4d1f
	return refd62e4d1f, allocsd62e4d1f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AccelerationStructureInstance) PassValue() (C.VkAccelerationStructureInstanceKHR, *cgoAllocMap) {
	if x.refd62e4d1f != nil {
		return *x.refd62e4d1f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AccelerationStructureInstance) Deref() {
	if x.refd62e4d1f == nil {
		return
	}
	x.Transform = *NewTransformMatrixRef(unsafe.Pointer(&x.refd62e4d1f.transform))
	x.InstanceCustomIndex = uint32(C.get_instanceCustomIndex(x.refd62e4d1f))
	x.Mask = uint32(C.get_mask(x.refd62e4d1f))
	x.InstanceShaderBindingTableRecordOffset = uint32(C.get_instanceShaderBindingTableRecordOffset(x.refd62e4d1f))
	x.Flags = GeometryInstanceFlags(C.get_flags(x.refd62e4d1f))
	x.AccelerationStructureReference = (uint32)(x.refd62e4d1f.accelerationStructureReference)
}

type AccelerationStructureInstanceNV AccelerationStructureInstance

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *AccelerationStructureInstanceNV) Ref() *C.VkAccelerationStructureInstanceKHR {
	if x == nil {
		return nil
	}
	return x.refd62e4d1f
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *AccelerationStructureInstanceNV) Free() {
	if x != nil && x.allocsd62e4d1f != nil {
		x.allocsd62e4d1f.(*cgoAllocMap).Free()
		x.refd62e4d1f = nil
	}
}

// NewAccelerationStructureInstanceNVRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewAccelerationStructureInstanceNVRef(ref unsafe.Pointer) *AccelerationStructureInstanceNV {
	if ref == nil {
		return nil
	}
	obj := new(AccelerationStructureInstanceNV)
	obj.refd62e4d1f = (*C.VkAccelerationStructureInstanceKHR)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *AccelerationStructureInstanceNV) PassRef() (*C.VkAccelerationStructureInstanceKHR, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refd62e4d1f != nil {
		return x.refd62e4d1f, nil
	}
	memd62e4d1f := allocAccelerationStructureInstanceMemory(1)
	refd62e4d1f := (*C.VkAccelerationStructureInstanceKHR)(memd62e4d1f)
	allocsd62e4d1f := new(cgoAllocMap)
	allocsd62e4d1f.Add(memd62e4d1f)

	var ctransform_allocs *cgoAllocMap
	refd62e4d1f.transform, ctransform_allocs = x.Transform.PassValue()
	allocsd62e4d1f.Borrow(ctransform_allocs)

	C.write_bits(refd62e4d1f,
		(C.uint32_t)(x.InstanceCustomIndex),
		(C.uint32_t)(x.Mask),
		(C.uint32_t)(x.InstanceShaderBindingTableRecordOffset),
		(C.VkGeometryInstanceFlagsKHR)(x.Flags))

	var caccelerationStructureReference_allocs *cgoAllocMap
	refd62e4d1f.accelerationStructureReference, caccelerationStructureReference_allocs = (C.uint64_t)(x.AccelerationStructureReference), cgoAllocsUnknown
	allocsd62e4d1f.Borrow(caccelerationStructureReference_allocs)

	x.refd62e4d1f = refd62e4d1f
	x.allocsd62e4d1f = allocsd62e4d1f
	return refd62e4d1f, allocsd62e4d1f

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x AccelerationStructureInstanceNV) PassValue() (C.VkAccelerationStructureInstanceKHR, *cgoAllocMap) {
	if x.refd62e4d1f != nil {
		return *x.refd62e4d1f, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *AccelerationStructureInstanceNV) Deref() {
	if x.refd62e4d1f == nil {
		return
	}
	x.Transform = *NewTransformMatrixRef(unsafe.Pointer(&x.refd62e4d1f.transform))
	x.InstanceCustomIndex = uint32(C.get_instanceCustomIndex(x.refd62e4d1f))
	x.Mask = uint32(C.get_mask(x.refd62e4d1f))
	x.InstanceShaderBindingTableRecordOffset = uint32(C.get_instanceShaderBindingTableRecordOffset(x.refd62e4d1f))
	x.Flags = GeometryInstanceFlags(C.get_flags(x.refd62e4d1f))
	x.AccelerationStructureReference = (uint32)(x.refd62e4d1f.accelerationStructureReference)
}
