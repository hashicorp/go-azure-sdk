package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceLinkDeviceVendor string

const (
	NetworkaccessDeviceLinkDeviceVendor_BarracudaNetworks  NetworkaccessDeviceLinkDeviceVendor = "barracudaNetworks"
	NetworkaccessDeviceLinkDeviceVendor_CheckPoint         NetworkaccessDeviceLinkDeviceVendor = "checkPoint"
	NetworkaccessDeviceLinkDeviceVendor_CiscoMeraki        NetworkaccessDeviceLinkDeviceVendor = "ciscoMeraki"
	NetworkaccessDeviceLinkDeviceVendor_Citrix             NetworkaccessDeviceLinkDeviceVendor = "citrix"
	NetworkaccessDeviceLinkDeviceVendor_Fortinet           NetworkaccessDeviceLinkDeviceVendor = "fortinet"
	NetworkaccessDeviceLinkDeviceVendor_HpeAruba           NetworkaccessDeviceLinkDeviceVendor = "hpeAruba"
	NetworkaccessDeviceLinkDeviceVendor_NetFoundry         NetworkaccessDeviceLinkDeviceVendor = "netFoundry"
	NetworkaccessDeviceLinkDeviceVendor_Nuage              NetworkaccessDeviceLinkDeviceVendor = "nuage"
	NetworkaccessDeviceLinkDeviceVendor_OpenSystems        NetworkaccessDeviceLinkDeviceVendor = "openSystems"
	NetworkaccessDeviceLinkDeviceVendor_Other              NetworkaccessDeviceLinkDeviceVendor = "other"
	NetworkaccessDeviceLinkDeviceVendor_PaloAltoNetworks   NetworkaccessDeviceLinkDeviceVendor = "paloAltoNetworks"
	NetworkaccessDeviceLinkDeviceVendor_RiverbedTechnology NetworkaccessDeviceLinkDeviceVendor = "riverbedTechnology"
	NetworkaccessDeviceLinkDeviceVendor_SilverPeak         NetworkaccessDeviceLinkDeviceVendor = "silverPeak"
	NetworkaccessDeviceLinkDeviceVendor_Versa              NetworkaccessDeviceLinkDeviceVendor = "versa"
	NetworkaccessDeviceLinkDeviceVendor_VmWareSdWan        NetworkaccessDeviceLinkDeviceVendor = "vmWareSdWan"
)

func PossibleValuesForNetworkaccessDeviceLinkDeviceVendor() []string {
	return []string{
		string(NetworkaccessDeviceLinkDeviceVendor_BarracudaNetworks),
		string(NetworkaccessDeviceLinkDeviceVendor_CheckPoint),
		string(NetworkaccessDeviceLinkDeviceVendor_CiscoMeraki),
		string(NetworkaccessDeviceLinkDeviceVendor_Citrix),
		string(NetworkaccessDeviceLinkDeviceVendor_Fortinet),
		string(NetworkaccessDeviceLinkDeviceVendor_HpeAruba),
		string(NetworkaccessDeviceLinkDeviceVendor_NetFoundry),
		string(NetworkaccessDeviceLinkDeviceVendor_Nuage),
		string(NetworkaccessDeviceLinkDeviceVendor_OpenSystems),
		string(NetworkaccessDeviceLinkDeviceVendor_Other),
		string(NetworkaccessDeviceLinkDeviceVendor_PaloAltoNetworks),
		string(NetworkaccessDeviceLinkDeviceVendor_RiverbedTechnology),
		string(NetworkaccessDeviceLinkDeviceVendor_SilverPeak),
		string(NetworkaccessDeviceLinkDeviceVendor_Versa),
		string(NetworkaccessDeviceLinkDeviceVendor_VmWareSdWan),
	}
}

func (s *NetworkaccessDeviceLinkDeviceVendor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDeviceLinkDeviceVendor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDeviceLinkDeviceVendor(input string) (*NetworkaccessDeviceLinkDeviceVendor, error) {
	vals := map[string]NetworkaccessDeviceLinkDeviceVendor{
		"barracudanetworks":  NetworkaccessDeviceLinkDeviceVendor_BarracudaNetworks,
		"checkpoint":         NetworkaccessDeviceLinkDeviceVendor_CheckPoint,
		"ciscomeraki":        NetworkaccessDeviceLinkDeviceVendor_CiscoMeraki,
		"citrix":             NetworkaccessDeviceLinkDeviceVendor_Citrix,
		"fortinet":           NetworkaccessDeviceLinkDeviceVendor_Fortinet,
		"hpearuba":           NetworkaccessDeviceLinkDeviceVendor_HpeAruba,
		"netfoundry":         NetworkaccessDeviceLinkDeviceVendor_NetFoundry,
		"nuage":              NetworkaccessDeviceLinkDeviceVendor_Nuage,
		"opensystems":        NetworkaccessDeviceLinkDeviceVendor_OpenSystems,
		"other":              NetworkaccessDeviceLinkDeviceVendor_Other,
		"paloaltonetworks":   NetworkaccessDeviceLinkDeviceVendor_PaloAltoNetworks,
		"riverbedtechnology": NetworkaccessDeviceLinkDeviceVendor_RiverbedTechnology,
		"silverpeak":         NetworkaccessDeviceLinkDeviceVendor_SilverPeak,
		"versa":              NetworkaccessDeviceLinkDeviceVendor_Versa,
		"vmwaresdwan":        NetworkaccessDeviceLinkDeviceVendor_VmWareSdWan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDeviceLinkDeviceVendor(input)
	return &out, nil
}
