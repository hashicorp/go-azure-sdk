package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupportedDeviceTypeType string

const (
	MobileAppSupportedDeviceTypeType_Android           MobileAppSupportedDeviceTypeType = "android"
	MobileAppSupportedDeviceTypeType_AndroidEnterprise MobileAppSupportedDeviceTypeType = "androidEnterprise"
	MobileAppSupportedDeviceTypeType_AndroidForWork    MobileAppSupportedDeviceTypeType = "androidForWork"
	MobileAppSupportedDeviceTypeType_AndroidnGMS       MobileAppSupportedDeviceTypeType = "androidnGMS"
	MobileAppSupportedDeviceTypeType_Blackberry        MobileAppSupportedDeviceTypeType = "blackberry"
	MobileAppSupportedDeviceTypeType_ChromeOS          MobileAppSupportedDeviceTypeType = "chromeOS"
	MobileAppSupportedDeviceTypeType_CloudPC           MobileAppSupportedDeviceTypeType = "cloudPC"
	MobileAppSupportedDeviceTypeType_Desktop           MobileAppSupportedDeviceTypeType = "desktop"
	MobileAppSupportedDeviceTypeType_HoloLens          MobileAppSupportedDeviceTypeType = "holoLens"
	MobileAppSupportedDeviceTypeType_IPad              MobileAppSupportedDeviceTypeType = "iPad"
	MobileAppSupportedDeviceTypeType_IPhone            MobileAppSupportedDeviceTypeType = "iPhone"
	MobileAppSupportedDeviceTypeType_IPod              MobileAppSupportedDeviceTypeType = "iPod"
	MobileAppSupportedDeviceTypeType_ISocConsumer      MobileAppSupportedDeviceTypeType = "iSocConsumer"
	MobileAppSupportedDeviceTypeType_Linux             MobileAppSupportedDeviceTypeType = "linux"
	MobileAppSupportedDeviceTypeType_Mac               MobileAppSupportedDeviceTypeType = "mac"
	MobileAppSupportedDeviceTypeType_MacMDM            MobileAppSupportedDeviceTypeType = "macMDM"
	MobileAppSupportedDeviceTypeType_Nokia             MobileAppSupportedDeviceTypeType = "nokia"
	MobileAppSupportedDeviceTypeType_Palm              MobileAppSupportedDeviceTypeType = "palm"
	MobileAppSupportedDeviceTypeType_SurfaceHub        MobileAppSupportedDeviceTypeType = "surfaceHub"
	MobileAppSupportedDeviceTypeType_Unix              MobileAppSupportedDeviceTypeType = "unix"
	MobileAppSupportedDeviceTypeType_Unknown           MobileAppSupportedDeviceTypeType = "unknown"
	MobileAppSupportedDeviceTypeType_WinCE             MobileAppSupportedDeviceTypeType = "winCE"
	MobileAppSupportedDeviceTypeType_WinEmbedded       MobileAppSupportedDeviceTypeType = "winEmbedded"
	MobileAppSupportedDeviceTypeType_WinMO6            MobileAppSupportedDeviceTypeType = "winMO6"
	MobileAppSupportedDeviceTypeType_Windows10x        MobileAppSupportedDeviceTypeType = "windows10x"
	MobileAppSupportedDeviceTypeType_WindowsPhone      MobileAppSupportedDeviceTypeType = "windowsPhone"
	MobileAppSupportedDeviceTypeType_WindowsRT         MobileAppSupportedDeviceTypeType = "windowsRT"
)

func PossibleValuesForMobileAppSupportedDeviceTypeType() []string {
	return []string{
		string(MobileAppSupportedDeviceTypeType_Android),
		string(MobileAppSupportedDeviceTypeType_AndroidEnterprise),
		string(MobileAppSupportedDeviceTypeType_AndroidForWork),
		string(MobileAppSupportedDeviceTypeType_AndroidnGMS),
		string(MobileAppSupportedDeviceTypeType_Blackberry),
		string(MobileAppSupportedDeviceTypeType_ChromeOS),
		string(MobileAppSupportedDeviceTypeType_CloudPC),
		string(MobileAppSupportedDeviceTypeType_Desktop),
		string(MobileAppSupportedDeviceTypeType_HoloLens),
		string(MobileAppSupportedDeviceTypeType_IPad),
		string(MobileAppSupportedDeviceTypeType_IPhone),
		string(MobileAppSupportedDeviceTypeType_IPod),
		string(MobileAppSupportedDeviceTypeType_ISocConsumer),
		string(MobileAppSupportedDeviceTypeType_Linux),
		string(MobileAppSupportedDeviceTypeType_Mac),
		string(MobileAppSupportedDeviceTypeType_MacMDM),
		string(MobileAppSupportedDeviceTypeType_Nokia),
		string(MobileAppSupportedDeviceTypeType_Palm),
		string(MobileAppSupportedDeviceTypeType_SurfaceHub),
		string(MobileAppSupportedDeviceTypeType_Unix),
		string(MobileAppSupportedDeviceTypeType_Unknown),
		string(MobileAppSupportedDeviceTypeType_WinCE),
		string(MobileAppSupportedDeviceTypeType_WinEmbedded),
		string(MobileAppSupportedDeviceTypeType_WinMO6),
		string(MobileAppSupportedDeviceTypeType_Windows10x),
		string(MobileAppSupportedDeviceTypeType_WindowsPhone),
		string(MobileAppSupportedDeviceTypeType_WindowsRT),
	}
}

func (s *MobileAppSupportedDeviceTypeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppSupportedDeviceTypeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppSupportedDeviceTypeType(input string) (*MobileAppSupportedDeviceTypeType, error) {
	vals := map[string]MobileAppSupportedDeviceTypeType{
		"android":           MobileAppSupportedDeviceTypeType_Android,
		"androidenterprise": MobileAppSupportedDeviceTypeType_AndroidEnterprise,
		"androidforwork":    MobileAppSupportedDeviceTypeType_AndroidForWork,
		"androidngms":       MobileAppSupportedDeviceTypeType_AndroidnGMS,
		"blackberry":        MobileAppSupportedDeviceTypeType_Blackberry,
		"chromeos":          MobileAppSupportedDeviceTypeType_ChromeOS,
		"cloudpc":           MobileAppSupportedDeviceTypeType_CloudPC,
		"desktop":           MobileAppSupportedDeviceTypeType_Desktop,
		"hololens":          MobileAppSupportedDeviceTypeType_HoloLens,
		"ipad":              MobileAppSupportedDeviceTypeType_IPad,
		"iphone":            MobileAppSupportedDeviceTypeType_IPhone,
		"ipod":              MobileAppSupportedDeviceTypeType_IPod,
		"isocconsumer":      MobileAppSupportedDeviceTypeType_ISocConsumer,
		"linux":             MobileAppSupportedDeviceTypeType_Linux,
		"mac":               MobileAppSupportedDeviceTypeType_Mac,
		"macmdm":            MobileAppSupportedDeviceTypeType_MacMDM,
		"nokia":             MobileAppSupportedDeviceTypeType_Nokia,
		"palm":              MobileAppSupportedDeviceTypeType_Palm,
		"surfacehub":        MobileAppSupportedDeviceTypeType_SurfaceHub,
		"unix":              MobileAppSupportedDeviceTypeType_Unix,
		"unknown":           MobileAppSupportedDeviceTypeType_Unknown,
		"wince":             MobileAppSupportedDeviceTypeType_WinCE,
		"winembedded":       MobileAppSupportedDeviceTypeType_WinEmbedded,
		"winmo6":            MobileAppSupportedDeviceTypeType_WinMO6,
		"windows10x":        MobileAppSupportedDeviceTypeType_Windows10x,
		"windowsphone":      MobileAppSupportedDeviceTypeType_WindowsPhone,
		"windowsrt":         MobileAppSupportedDeviceTypeType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupportedDeviceTypeType(input)
	return &out, nil
}
