package adsi

import (
	"io"
	"sync"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/scjalliance/comshim"
	"github.com/scjalliance/comutil"
	"gopkg.in/adsi.v0/api"
)

// Container provides access to Active Directory container objects.
type Container struct {
	m     sync.RWMutex
	iface *api.IADsContainer
}

// NewContainer returns a container that manages the given COM interface.
func NewContainer(iface *api.IADsContainer) *Container {
	comshim.Add(1)
	return &Container{iface: iface}
}

func (c *Container) closed() bool {
	return (c.iface == nil)
}

// Close will release resources consumed by the container. It should be
// called when the container is no longer needed.
func (c *Container) Close() {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		c.iface.Release()
		return nil
	})
	// FIXME: What happens if the run returns an error?
	c.iface = nil
}

// Children returns an object iterator that provides access to the immediate
// children of the container.
func (c *Container) Children() (iter *ObjectIter, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		iunknown, err := c.iface.NewEnum()
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

// Filter returns the current filter of the container.
func (c *Container) Filter() (filter []string, err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		variant, err := c.iface.Filter()
		if err != nil {
			return err
		}
		defer variant.Clear()
		filter = variant.ToArray().ToStringArray()
		return nil
	})
	return
}

// SetFilter set the filter for the container.
func (c *Container) SetFilter(filter ...string) (err error) {
	c.m.Lock()
	defer c.m.Unlock()
	if c.closed() {
		return ErrClosed
	}
	err = run(func() error {
		safeByteArray := comutil.SafeArrayFromStringSlice(filter)
		variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_BSTR, int64(uintptr(unsafe.Pointer(safeByteArray))))
		v := &variant
		defer v.Clear()
		return c.iface.SetFilter(v)
	})
	return
}

// ObjectIter provides an iterator for a set of objects.
type ObjectIter struct {
	m     sync.RWMutex
	iface *ole.IEnumVARIANT
}

// NewObjectIter returns an object iterator that provides access to the objects
// contained in the given enumerator.
func NewObjectIter(enumerator *ole.IEnumVARIANT) *ObjectIter {
	comshim.Add(1)
	return &ObjectIter{iface: enumerator}
}

// Next moves the iterator to the next object and returns a pointer to it. If it
// has reached the end of the set it will return io.EOF. It the iterator has
// already been closed it will return ErrClosed.
//
// FIXME: Make sure that io.EOF is being returned as expected. We might have
// to intercept an internal error.
func (iter *ObjectIter) Next() (obj *Object, err error) {
	iter.m.Lock()
	defer iter.m.Unlock()
	if iter.closed() {
		return nil, ErrClosed
	}
	err = run(func() error {
		// See https://msdn.microsoft.com/library/aa705990
		array, length, err := iter.iface.Next(1)
		if err != nil {
			return err
		}
		defer array.Clear()
		if length == 0 {
			return io.EOF
		}

		idispatch := array.ToIDispatch()
		if idispatch == nil {
			return ErrNonDispatchVariant
		}
		// Note: Do *not* call idispatch.Release() here, as it will be called
		//       automatically by array.Clear()

		iresult, err := idispatch.QueryInterface(api.IID_IADs)
		if err != nil {
			return err
		}
		iface := (*api.IADs)(unsafe.Pointer(iresult))
		obj = NewObject(iface)
		return nil
	})
	return
}

func (iter *ObjectIter) closed() bool {
	return (iter.iface == nil)
}

// Close will release resources consumed by the iterator. It should be
// called when the iterator is no longer needed.
func (iter *ObjectIter) Close() {
	iter.m.Lock()
	defer iter.m.Unlock()
	if iter.closed() {
		return
	}
	defer comshim.Done()
	run(func() error {
		iter.iface.Release()
		return nil
	})
	// FIXME: What happens if the run returns an error?
	iter.iface = nil
}
