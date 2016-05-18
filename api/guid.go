package api

import (
	ole "github.com/go-ole/go-ole"
)

// See: https://github.com/blueardour/pxe-server/blob/master/0.winxp-drivers-via-binlsrv/inf/activeds.inf

var (
	// {228D9A81-C302-11CF-9AA4-00AA004A5691}
	CLSID_LDAP = &ole.GUID{0x228D9A81, 0xC302, 0x11CF, [8]byte{0x9A, 0xA4, 0x00, 0xAA, 0x00, 0x4A, 0x56, 0x91}}

	// {228D9A82-C302-11CF-9AA4-00AA004A5691}
	CLSID_LDAPNamespace = &ole.GUID{0x228D9A82, 0xC302, 0x11CF, [8]byte{0x9A, 0xA4, 0x00, 0xAA, 0x00, 0x4A, 0x56, 0x91}}

	// {250E91A0-0367-11CF-ABC4-02608C9E7553}
	CLSID_WinNTNamespace = &ole.GUID{0x250E91A0, 0x0367, 0x11CF, [8]byte{0xAB, 0xC4, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}}

	// CLSID_ADs is the component object model identifier of the Active
	// Directory Services class.
	//
	// {4753DA60-5B71-11CF-B035-00AA006E0975}
	CLSID_ADs = &ole.GUID{0x4753DA60, 0x5B71, 0x11CF, [8]byte{0xB0, 0x35, 0x00, 0xAA, 0x00, 0x6E, 0x09, 0x75}}

	// {549365D0-EC26-11CF-8310-00AA00B505DB}
	CLSID_ADsDSOObject = &ole.GUID{0x549365D0, 0xEC26, 0x11CF, [8]byte{0x83, 0x10, 0x00, 0xAA, 0x00, 0xB5, 0x05, 0xDB}}

	// {233664B0-0367-11CF-ABC4-02608C9E7553}
	CLSID_ADsNamespaces = &ole.GUID{0x233664B0, 0x0367, 0x11CF, [8]byte{0xAB, 0xC4, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}}

	// {50B6327F-AFD1-11D2-9CB9-0000F87A369E}
	CLSID_ADSystemInfo = &ole.GUID{0x50B6327F, 0xAFD1, 0x11D2, [8]byte{0x9C, 0xB9, 0x00, 0x00, 0xF8, 0x7A, 0x36, 0x9E}}

	// {E0FA581D-2188-11D2-A739-00C04FA377A1}
	CLSID_ADsOLEDB = &ole.GUID{0xE0FA581D, 0x2188, 0x11D2, [8]byte{0xA7, 0x39, 0x00, 0xC0, 0x4F, 0xA3, 0x77, 0xA1}}

	// IID_IADs is the component object model identifier of the
	// IADs class.
	//
	// {FD8256D0-FD15-11CE-ABC4-02608C9E7553}
	IID_IADs = &ole.GUID{0xFD8256D0, 0xFD15, 0x11CE, [8]byte{0xAB, 0xC4, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}}

	// IID_IADsOpenDSObject is the component object model identifier of the
	// IADsOpenDSObject class.
	//
	// {DDF2891E-0F9C-11D0-8AD4-00C04FD8D503}
	IID_IADsOpenDSObject = &ole.GUID{0xDDF2891E, 0x0F9C, 0x11D0, [8]byte{0x8A, 0xD4, 0x00, 0xC0, 0x4F, 0xD8, 0xD5, 0x03}}

	// IID_IADsContainer is the component object model identifier of the
	// IADsContainer class.
	//
	// {001677D0-FD16-11CE-ABC4-02608C9E7553}
	IID_IADsContainer = &ole.GUID{0x001677D0, 0xFD16, 0x11CE, [8]byte{0xAB, 0xC4, 0x02, 0x60, 0x8C, 0x9E, 0x75, 0x53}}
)
