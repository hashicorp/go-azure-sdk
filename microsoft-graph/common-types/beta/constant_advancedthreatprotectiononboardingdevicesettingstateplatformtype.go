package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType string

const (
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Android           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "android"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidEnterprise AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "androidEnterprise"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidForWork    AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "androidForWork"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidnGMS       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "androidnGMS"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Blackberry        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "blackberry"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_ChromeOS          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "chromeOS"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_CloudPC           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "cloudPC"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Desktop           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "desktop"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_HoloLens          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "holoLens"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPad              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "iPad"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPhone            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "iPhone"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPod              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "iPod"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_ISocConsumer      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "iSocConsumer"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Linux             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "linux"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Mac               AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "mac"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_MacMDM            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "macMDM"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Nokia             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "nokia"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Palm              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "palm"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_SurfaceHub        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "surfaceHub"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Unix              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "unix"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Unknown           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "unknown"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinCE             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "winCE"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinEmbedded       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "winEmbedded"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinMO6            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "winMO6"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Windows10x        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "windows10x"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WindowsPhone      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "windowsPhone"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WindowsRT         AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "windowsRT"
)

func PossibleValuesForAdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType() []string {
	return []string{
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Android),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidEnterprise),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidForWork),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidnGMS),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Blackberry),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_ChromeOS),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_CloudPC),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Desktop),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_HoloLens),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPad),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPhone),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPod),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_ISocConsumer),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Linux),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Mac),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_MacMDM),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Nokia),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Palm),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_SurfaceHub),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Unix),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Unknown),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinCE),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinEmbedded),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinMO6),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Windows10x),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WindowsPhone),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WindowsRT),
	}
}

func (s *AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType(input string) (*AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType, error) {
	vals := map[string]AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType{
		"android":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Android,
		"androidenterprise": AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidEnterprise,
		"androidforwork":    AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidForWork,
		"androidngms":       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_AndroidnGMS,
		"blackberry":        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Blackberry,
		"chromeos":          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_ChromeOS,
		"cloudpc":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_CloudPC,
		"desktop":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Desktop,
		"hololens":          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_HoloLens,
		"ipad":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPad,
		"iphone":            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPhone,
		"ipod":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_IPod,
		"isocconsumer":      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_ISocConsumer,
		"linux":             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Linux,
		"mac":               AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Mac,
		"macmdm":            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_MacMDM,
		"nokia":             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Nokia,
		"palm":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Palm,
		"surfacehub":        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_SurfaceHub,
		"unix":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Unix,
		"unknown":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Unknown,
		"wince":             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinCE,
		"winembedded":       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinEmbedded,
		"winmo6":            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WinMO6,
		"windows10x":        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_Windows10x,
		"windowsphone":      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WindowsPhone,
		"windowsrt":         AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType(input)
	return &out, nil
}
