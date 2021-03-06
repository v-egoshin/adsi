package adsi

import (
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
	"github.com/scjalliance/comutil"

	"gopkg.in/adsi.v0/api"
)

// Members provides access to group membership.
type Members struct {
	m     sync.RWMutex
	iface *api.IADsMembers
}

// NewMembers returns a membership that manages the given COM
// interface.
func NewMembers(iface *api.IADsMembers) *Members {
	comshim.Add(1)
	return &Members{iface: iface}
}

func (m *Members) closed() bool {
	return (m.iface == nil)
}

// Close will release resources consumed by the membership. It should be
// called when the membership is no longer needed.
func (m *Members) Close() {
	m.m.Lock()
	defer m.m.Unlock()
	if m.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		m.iface.Release()
		return nil
	})
	m.iface = nil
}

// Iter returns an object iterator that provides access to the members
// of the group.
func (m *Members) Iter() (iter *ObjectIter, err error) {
	m.m.Lock()
	defer m.m.Unlock()
	if m.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		iunknown, err := m.iface.NewEnum()
		if err != nil {
			return err
		}
		defer iunknown.Release()
		idispatch, err := iunknown.QueryInterface(ole.IID_IEnumVariant)
		if err != nil {
			return err
		}
		iface := (*ole.IEnumVARIANT)(unsafe.Pointer(idispatch))
		iter = NewObjectIter(iface)
		return nil
	})
	return
}

// Filter returns the current filter of the membership.
func (m *Members) Filter() (filter []string, err error) {
	m.m.Lock()
	defer m.m.Unlock()
	if m.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		variant, err := m.iface.Filter()
		if err != nil {
			return err
		}
		defer variant.Clear()
		filter = variant.ToArray().ToStringArray()
		return nil
	})
	return
}

// SetFilter set the filter for the mebership.
func (m *Members) SetFilter(filter ...string) (err error) {
	m.m.Lock()
	defer m.m.Unlock()
	if m.closed() {
		return ErrClosed
	}
	err = run(func() error {
		safeByteArray := comutil.SafeArrayFromStringSlice(filter)
		variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_BSTR, int64(uintptr(unsafe.Pointer(safeByteArray))))
		v := &variant
		defer v.Clear()
		return m.iface.SetFilter(v)
	})
	return
}
