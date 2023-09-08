package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceLinkDeviceVendor string

const (
	NetworkaccessDeviceLinkDeviceVendorbarracudaNetworks  NetworkaccessDeviceLinkDeviceVendor = "BarracudaNetworks"
	NetworkaccessDeviceLinkDeviceVendorcheckPoint         NetworkaccessDeviceLinkDeviceVendor = "CheckPoint"
	NetworkaccessDeviceLinkDeviceVendorciscoMeraki        NetworkaccessDeviceLinkDeviceVendor = "CiscoMeraki"
	NetworkaccessDeviceLinkDeviceVendorcitrix             NetworkaccessDeviceLinkDeviceVendor = "Citrix"
	NetworkaccessDeviceLinkDeviceVendorfortinet           NetworkaccessDeviceLinkDeviceVendor = "Fortinet"
	NetworkaccessDeviceLinkDeviceVendorhpeAruba           NetworkaccessDeviceLinkDeviceVendor = "HpeAruba"
	NetworkaccessDeviceLinkDeviceVendornetFoundry         NetworkaccessDeviceLinkDeviceVendor = "NetFoundry"
	NetworkaccessDeviceLinkDeviceVendornuage              NetworkaccessDeviceLinkDeviceVendor = "Nuage"
	NetworkaccessDeviceLinkDeviceVendoropenSystems        NetworkaccessDeviceLinkDeviceVendor = "OpenSystems"
	NetworkaccessDeviceLinkDeviceVendorother              NetworkaccessDeviceLinkDeviceVendor = "Other"
	NetworkaccessDeviceLinkDeviceVendorpaloAltoNetworks   NetworkaccessDeviceLinkDeviceVendor = "PaloAltoNetworks"
	NetworkaccessDeviceLinkDeviceVendorriverbedTechnology NetworkaccessDeviceLinkDeviceVendor = "RiverbedTechnology"
	NetworkaccessDeviceLinkDeviceVendorsilverPeak         NetworkaccessDeviceLinkDeviceVendor = "SilverPeak"
	NetworkaccessDeviceLinkDeviceVendorversa              NetworkaccessDeviceLinkDeviceVendor = "Versa"
	NetworkaccessDeviceLinkDeviceVendorvmWareSdWan        NetworkaccessDeviceLinkDeviceVendor = "VmWareSdWan"
)

func PossibleValuesForNetworkaccessDeviceLinkDeviceVendor() []string {
	return []string{
		string(NetworkaccessDeviceLinkDeviceVendorbarracudaNetworks),
		string(NetworkaccessDeviceLinkDeviceVendorcheckPoint),
		string(NetworkaccessDeviceLinkDeviceVendorciscoMeraki),
		string(NetworkaccessDeviceLinkDeviceVendorcitrix),
		string(NetworkaccessDeviceLinkDeviceVendorfortinet),
		string(NetworkaccessDeviceLinkDeviceVendorhpeAruba),
		string(NetworkaccessDeviceLinkDeviceVendornetFoundry),
		string(NetworkaccessDeviceLinkDeviceVendornuage),
		string(NetworkaccessDeviceLinkDeviceVendoropenSystems),
		string(NetworkaccessDeviceLinkDeviceVendorother),
		string(NetworkaccessDeviceLinkDeviceVendorpaloAltoNetworks),
		string(NetworkaccessDeviceLinkDeviceVendorriverbedTechnology),
		string(NetworkaccessDeviceLinkDeviceVendorsilverPeak),
		string(NetworkaccessDeviceLinkDeviceVendorversa),
		string(NetworkaccessDeviceLinkDeviceVendorvmWareSdWan),
	}
}

func parseNetworkaccessDeviceLinkDeviceVendor(input string) (*NetworkaccessDeviceLinkDeviceVendor, error) {
	vals := map[string]NetworkaccessDeviceLinkDeviceVendor{
		"barracudanetworks":  NetworkaccessDeviceLinkDeviceVendorbarracudaNetworks,
		"checkpoint":         NetworkaccessDeviceLinkDeviceVendorcheckPoint,
		"ciscomeraki":        NetworkaccessDeviceLinkDeviceVendorciscoMeraki,
		"citrix":             NetworkaccessDeviceLinkDeviceVendorcitrix,
		"fortinet":           NetworkaccessDeviceLinkDeviceVendorfortinet,
		"hpearuba":           NetworkaccessDeviceLinkDeviceVendorhpeAruba,
		"netfoundry":         NetworkaccessDeviceLinkDeviceVendornetFoundry,
		"nuage":              NetworkaccessDeviceLinkDeviceVendornuage,
		"opensystems":        NetworkaccessDeviceLinkDeviceVendoropenSystems,
		"other":              NetworkaccessDeviceLinkDeviceVendorother,
		"paloaltonetworks":   NetworkaccessDeviceLinkDeviceVendorpaloAltoNetworks,
		"riverbedtechnology": NetworkaccessDeviceLinkDeviceVendorriverbedTechnology,
		"silverpeak":         NetworkaccessDeviceLinkDeviceVendorsilverPeak,
		"versa":              NetworkaccessDeviceLinkDeviceVendorversa,
		"vmwaresdwan":        NetworkaccessDeviceLinkDeviceVendorvmWareSdWan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDeviceLinkDeviceVendor(input)
	return &out, nil
}
