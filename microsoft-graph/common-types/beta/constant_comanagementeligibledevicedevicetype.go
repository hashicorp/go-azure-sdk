package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceDeviceType string

const (
	ComanagementEligibleDeviceDeviceType_Android           ComanagementEligibleDeviceDeviceType = "android"
	ComanagementEligibleDeviceDeviceType_AndroidEnterprise ComanagementEligibleDeviceDeviceType = "androidEnterprise"
	ComanagementEligibleDeviceDeviceType_AndroidForWork    ComanagementEligibleDeviceDeviceType = "androidForWork"
	ComanagementEligibleDeviceDeviceType_AndroidnGMS       ComanagementEligibleDeviceDeviceType = "androidnGMS"
	ComanagementEligibleDeviceDeviceType_Blackberry        ComanagementEligibleDeviceDeviceType = "blackberry"
	ComanagementEligibleDeviceDeviceType_ChromeOS          ComanagementEligibleDeviceDeviceType = "chromeOS"
	ComanagementEligibleDeviceDeviceType_CloudPC           ComanagementEligibleDeviceDeviceType = "cloudPC"
	ComanagementEligibleDeviceDeviceType_Desktop           ComanagementEligibleDeviceDeviceType = "desktop"
	ComanagementEligibleDeviceDeviceType_HoloLens          ComanagementEligibleDeviceDeviceType = "holoLens"
	ComanagementEligibleDeviceDeviceType_IPad              ComanagementEligibleDeviceDeviceType = "iPad"
	ComanagementEligibleDeviceDeviceType_IPhone            ComanagementEligibleDeviceDeviceType = "iPhone"
	ComanagementEligibleDeviceDeviceType_IPod              ComanagementEligibleDeviceDeviceType = "iPod"
	ComanagementEligibleDeviceDeviceType_ISocConsumer      ComanagementEligibleDeviceDeviceType = "iSocConsumer"
	ComanagementEligibleDeviceDeviceType_Linux             ComanagementEligibleDeviceDeviceType = "linux"
	ComanagementEligibleDeviceDeviceType_Mac               ComanagementEligibleDeviceDeviceType = "mac"
	ComanagementEligibleDeviceDeviceType_MacMDM            ComanagementEligibleDeviceDeviceType = "macMDM"
	ComanagementEligibleDeviceDeviceType_Nokia             ComanagementEligibleDeviceDeviceType = "nokia"
	ComanagementEligibleDeviceDeviceType_Palm              ComanagementEligibleDeviceDeviceType = "palm"
	ComanagementEligibleDeviceDeviceType_SurfaceHub        ComanagementEligibleDeviceDeviceType = "surfaceHub"
	ComanagementEligibleDeviceDeviceType_Unix              ComanagementEligibleDeviceDeviceType = "unix"
	ComanagementEligibleDeviceDeviceType_Unknown           ComanagementEligibleDeviceDeviceType = "unknown"
	ComanagementEligibleDeviceDeviceType_WinCE             ComanagementEligibleDeviceDeviceType = "winCE"
	ComanagementEligibleDeviceDeviceType_WinEmbedded       ComanagementEligibleDeviceDeviceType = "winEmbedded"
	ComanagementEligibleDeviceDeviceType_WinMO6            ComanagementEligibleDeviceDeviceType = "winMO6"
	ComanagementEligibleDeviceDeviceType_Windows10x        ComanagementEligibleDeviceDeviceType = "windows10x"
	ComanagementEligibleDeviceDeviceType_WindowsPhone      ComanagementEligibleDeviceDeviceType = "windowsPhone"
	ComanagementEligibleDeviceDeviceType_WindowsRT         ComanagementEligibleDeviceDeviceType = "windowsRT"
)

func PossibleValuesForComanagementEligibleDeviceDeviceType() []string {
	return []string{
		string(ComanagementEligibleDeviceDeviceType_Android),
		string(ComanagementEligibleDeviceDeviceType_AndroidEnterprise),
		string(ComanagementEligibleDeviceDeviceType_AndroidForWork),
		string(ComanagementEligibleDeviceDeviceType_AndroidnGMS),
		string(ComanagementEligibleDeviceDeviceType_Blackberry),
		string(ComanagementEligibleDeviceDeviceType_ChromeOS),
		string(ComanagementEligibleDeviceDeviceType_CloudPC),
		string(ComanagementEligibleDeviceDeviceType_Desktop),
		string(ComanagementEligibleDeviceDeviceType_HoloLens),
		string(ComanagementEligibleDeviceDeviceType_IPad),
		string(ComanagementEligibleDeviceDeviceType_IPhone),
		string(ComanagementEligibleDeviceDeviceType_IPod),
		string(ComanagementEligibleDeviceDeviceType_ISocConsumer),
		string(ComanagementEligibleDeviceDeviceType_Linux),
		string(ComanagementEligibleDeviceDeviceType_Mac),
		string(ComanagementEligibleDeviceDeviceType_MacMDM),
		string(ComanagementEligibleDeviceDeviceType_Nokia),
		string(ComanagementEligibleDeviceDeviceType_Palm),
		string(ComanagementEligibleDeviceDeviceType_SurfaceHub),
		string(ComanagementEligibleDeviceDeviceType_Unix),
		string(ComanagementEligibleDeviceDeviceType_Unknown),
		string(ComanagementEligibleDeviceDeviceType_WinCE),
		string(ComanagementEligibleDeviceDeviceType_WinEmbedded),
		string(ComanagementEligibleDeviceDeviceType_WinMO6),
		string(ComanagementEligibleDeviceDeviceType_Windows10x),
		string(ComanagementEligibleDeviceDeviceType_WindowsPhone),
		string(ComanagementEligibleDeviceDeviceType_WindowsRT),
	}
}

func (s *ComanagementEligibleDeviceDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseComanagementEligibleDeviceDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseComanagementEligibleDeviceDeviceType(input string) (*ComanagementEligibleDeviceDeviceType, error) {
	vals := map[string]ComanagementEligibleDeviceDeviceType{
		"android":           ComanagementEligibleDeviceDeviceType_Android,
		"androidenterprise": ComanagementEligibleDeviceDeviceType_AndroidEnterprise,
		"androidforwork":    ComanagementEligibleDeviceDeviceType_AndroidForWork,
		"androidngms":       ComanagementEligibleDeviceDeviceType_AndroidnGMS,
		"blackberry":        ComanagementEligibleDeviceDeviceType_Blackberry,
		"chromeos":          ComanagementEligibleDeviceDeviceType_ChromeOS,
		"cloudpc":           ComanagementEligibleDeviceDeviceType_CloudPC,
		"desktop":           ComanagementEligibleDeviceDeviceType_Desktop,
		"hololens":          ComanagementEligibleDeviceDeviceType_HoloLens,
		"ipad":              ComanagementEligibleDeviceDeviceType_IPad,
		"iphone":            ComanagementEligibleDeviceDeviceType_IPhone,
		"ipod":              ComanagementEligibleDeviceDeviceType_IPod,
		"isocconsumer":      ComanagementEligibleDeviceDeviceType_ISocConsumer,
		"linux":             ComanagementEligibleDeviceDeviceType_Linux,
		"mac":               ComanagementEligibleDeviceDeviceType_Mac,
		"macmdm":            ComanagementEligibleDeviceDeviceType_MacMDM,
		"nokia":             ComanagementEligibleDeviceDeviceType_Nokia,
		"palm":              ComanagementEligibleDeviceDeviceType_Palm,
		"surfacehub":        ComanagementEligibleDeviceDeviceType_SurfaceHub,
		"unix":              ComanagementEligibleDeviceDeviceType_Unix,
		"unknown":           ComanagementEligibleDeviceDeviceType_Unknown,
		"wince":             ComanagementEligibleDeviceDeviceType_WinCE,
		"winembedded":       ComanagementEligibleDeviceDeviceType_WinEmbedded,
		"winmo6":            ComanagementEligibleDeviceDeviceType_WinMO6,
		"windows10x":        ComanagementEligibleDeviceDeviceType_Windows10x,
		"windowsphone":      ComanagementEligibleDeviceDeviceType_WindowsPhone,
		"windowsrt":         ComanagementEligibleDeviceDeviceType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceDeviceType(input)
	return &out, nil
}
