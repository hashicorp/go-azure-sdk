package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceType string

const (
	ManagedDeviceDeviceType_Android           ManagedDeviceDeviceType = "android"
	ManagedDeviceDeviceType_AndroidEnterprise ManagedDeviceDeviceType = "androidEnterprise"
	ManagedDeviceDeviceType_AndroidForWork    ManagedDeviceDeviceType = "androidForWork"
	ManagedDeviceDeviceType_AndroidnGMS       ManagedDeviceDeviceType = "androidnGMS"
	ManagedDeviceDeviceType_Blackberry        ManagedDeviceDeviceType = "blackberry"
	ManagedDeviceDeviceType_ChromeOS          ManagedDeviceDeviceType = "chromeOS"
	ManagedDeviceDeviceType_CloudPC           ManagedDeviceDeviceType = "cloudPC"
	ManagedDeviceDeviceType_Desktop           ManagedDeviceDeviceType = "desktop"
	ManagedDeviceDeviceType_HoloLens          ManagedDeviceDeviceType = "holoLens"
	ManagedDeviceDeviceType_IPad              ManagedDeviceDeviceType = "iPad"
	ManagedDeviceDeviceType_IPhone            ManagedDeviceDeviceType = "iPhone"
	ManagedDeviceDeviceType_IPod              ManagedDeviceDeviceType = "iPod"
	ManagedDeviceDeviceType_ISocConsumer      ManagedDeviceDeviceType = "iSocConsumer"
	ManagedDeviceDeviceType_Linux             ManagedDeviceDeviceType = "linux"
	ManagedDeviceDeviceType_Mac               ManagedDeviceDeviceType = "mac"
	ManagedDeviceDeviceType_MacMDM            ManagedDeviceDeviceType = "macMDM"
	ManagedDeviceDeviceType_Nokia             ManagedDeviceDeviceType = "nokia"
	ManagedDeviceDeviceType_Palm              ManagedDeviceDeviceType = "palm"
	ManagedDeviceDeviceType_SurfaceHub        ManagedDeviceDeviceType = "surfaceHub"
	ManagedDeviceDeviceType_Unix              ManagedDeviceDeviceType = "unix"
	ManagedDeviceDeviceType_Unknown           ManagedDeviceDeviceType = "unknown"
	ManagedDeviceDeviceType_WinCE             ManagedDeviceDeviceType = "winCE"
	ManagedDeviceDeviceType_WinEmbedded       ManagedDeviceDeviceType = "winEmbedded"
	ManagedDeviceDeviceType_WinMO6            ManagedDeviceDeviceType = "winMO6"
	ManagedDeviceDeviceType_Windows10x        ManagedDeviceDeviceType = "windows10x"
	ManagedDeviceDeviceType_WindowsPhone      ManagedDeviceDeviceType = "windowsPhone"
	ManagedDeviceDeviceType_WindowsRT         ManagedDeviceDeviceType = "windowsRT"
)

func PossibleValuesForManagedDeviceDeviceType() []string {
	return []string{
		string(ManagedDeviceDeviceType_Android),
		string(ManagedDeviceDeviceType_AndroidEnterprise),
		string(ManagedDeviceDeviceType_AndroidForWork),
		string(ManagedDeviceDeviceType_AndroidnGMS),
		string(ManagedDeviceDeviceType_Blackberry),
		string(ManagedDeviceDeviceType_ChromeOS),
		string(ManagedDeviceDeviceType_CloudPC),
		string(ManagedDeviceDeviceType_Desktop),
		string(ManagedDeviceDeviceType_HoloLens),
		string(ManagedDeviceDeviceType_IPad),
		string(ManagedDeviceDeviceType_IPhone),
		string(ManagedDeviceDeviceType_IPod),
		string(ManagedDeviceDeviceType_ISocConsumer),
		string(ManagedDeviceDeviceType_Linux),
		string(ManagedDeviceDeviceType_Mac),
		string(ManagedDeviceDeviceType_MacMDM),
		string(ManagedDeviceDeviceType_Nokia),
		string(ManagedDeviceDeviceType_Palm),
		string(ManagedDeviceDeviceType_SurfaceHub),
		string(ManagedDeviceDeviceType_Unix),
		string(ManagedDeviceDeviceType_Unknown),
		string(ManagedDeviceDeviceType_WinCE),
		string(ManagedDeviceDeviceType_WinEmbedded),
		string(ManagedDeviceDeviceType_WinMO6),
		string(ManagedDeviceDeviceType_Windows10x),
		string(ManagedDeviceDeviceType_WindowsPhone),
		string(ManagedDeviceDeviceType_WindowsRT),
	}
}

func (s *ManagedDeviceDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceDeviceType(input string) (*ManagedDeviceDeviceType, error) {
	vals := map[string]ManagedDeviceDeviceType{
		"android":           ManagedDeviceDeviceType_Android,
		"androidenterprise": ManagedDeviceDeviceType_AndroidEnterprise,
		"androidforwork":    ManagedDeviceDeviceType_AndroidForWork,
		"androidngms":       ManagedDeviceDeviceType_AndroidnGMS,
		"blackberry":        ManagedDeviceDeviceType_Blackberry,
		"chromeos":          ManagedDeviceDeviceType_ChromeOS,
		"cloudpc":           ManagedDeviceDeviceType_CloudPC,
		"desktop":           ManagedDeviceDeviceType_Desktop,
		"hololens":          ManagedDeviceDeviceType_HoloLens,
		"ipad":              ManagedDeviceDeviceType_IPad,
		"iphone":            ManagedDeviceDeviceType_IPhone,
		"ipod":              ManagedDeviceDeviceType_IPod,
		"isocconsumer":      ManagedDeviceDeviceType_ISocConsumer,
		"linux":             ManagedDeviceDeviceType_Linux,
		"mac":               ManagedDeviceDeviceType_Mac,
		"macmdm":            ManagedDeviceDeviceType_MacMDM,
		"nokia":             ManagedDeviceDeviceType_Nokia,
		"palm":              ManagedDeviceDeviceType_Palm,
		"surfacehub":        ManagedDeviceDeviceType_SurfaceHub,
		"unix":              ManagedDeviceDeviceType_Unix,
		"unknown":           ManagedDeviceDeviceType_Unknown,
		"wince":             ManagedDeviceDeviceType_WinCE,
		"winembedded":       ManagedDeviceDeviceType_WinEmbedded,
		"winmo6":            ManagedDeviceDeviceType_WinMO6,
		"windows10x":        ManagedDeviceDeviceType_Windows10x,
		"windowsphone":      ManagedDeviceDeviceType_WindowsPhone,
		"windowsrt":         ManagedDeviceDeviceType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceType(input)
	return &out, nil
}
