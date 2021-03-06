// +build windows

package api

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comutil"
)

// NewIADsContainer returns a new instance of the IADsContainer
// component object model interface.
func NewIADsContainer(server string, clsid *ole.GUID) (*IADsContainer, error) {
	p, err := comutil.CreateRemoteObject(server, clsid, IID_IADsContainer)
	return (*IADsContainer)(unsafe.Pointer(p)), err
}

// NewEnum retrieves an enumerator interface that provides access to the objects
// within the container.
//
// See https://msdn.microsoft.com/library/aa705990
func (v *IADsContainer) NewEnum() (enum *ole.IUnknown, err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().NewEnum),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(&enum)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}

// Filter retrieves the filter for the container.
func (v *IADsContainer) Filter() (variant *ole.VARIANT, err error) {
	variant = new(ole.VARIANT)
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().Filter),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(variant)),
		0)
	if hr != 0 {
		return nil, convertHresultToError(hr)
	}
	return
}

// SetFilter sets the filter for the container.
func (v *IADsContainer) SetFilter(variant *ole.VARIANT) (err error) {
	hr, _, _ := syscall.Syscall(
		uintptr(v.VTable().SetFilter),
		2,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(variant)),
		0)
	if hr != 0 {
		return convertHresultToError(hr)
	}
	return nil
}
