package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceSettingStatePlatformType string

const (
	DeviceComplianceSettingStatePlatformType_Android           DeviceComplianceSettingStatePlatformType = "android"
	DeviceComplianceSettingStatePlatformType_AndroidEnterprise DeviceComplianceSettingStatePlatformType = "androidEnterprise"
	DeviceComplianceSettingStatePlatformType_AndroidForWork    DeviceComplianceSettingStatePlatformType = "androidForWork"
	DeviceComplianceSettingStatePlatformType_AndroidnGMS       DeviceComplianceSettingStatePlatformType = "androidnGMS"
	DeviceComplianceSettingStatePlatformType_Blackberry        DeviceComplianceSettingStatePlatformType = "blackberry"
	DeviceComplianceSettingStatePlatformType_ChromeOS          DeviceComplianceSettingStatePlatformType = "chromeOS"
	DeviceComplianceSettingStatePlatformType_CloudPC           DeviceComplianceSettingStatePlatformType = "cloudPC"
	DeviceComplianceSettingStatePlatformType_Desktop           DeviceComplianceSettingStatePlatformType = "desktop"
	DeviceComplianceSettingStatePlatformType_HoloLens          DeviceComplianceSettingStatePlatformType = "holoLens"
	DeviceComplianceSettingStatePlatformType_IPad              DeviceComplianceSettingStatePlatformType = "iPad"
	DeviceComplianceSettingStatePlatformType_IPhone            DeviceComplianceSettingStatePlatformType = "iPhone"
	DeviceComplianceSettingStatePlatformType_IPod              DeviceComplianceSettingStatePlatformType = "iPod"
	DeviceComplianceSettingStatePlatformType_ISocConsumer      DeviceComplianceSettingStatePlatformType = "iSocConsumer"
	DeviceComplianceSettingStatePlatformType_Linux             DeviceComplianceSettingStatePlatformType = "linux"
	DeviceComplianceSettingStatePlatformType_Mac               DeviceComplianceSettingStatePlatformType = "mac"
	DeviceComplianceSettingStatePlatformType_MacMDM            DeviceComplianceSettingStatePlatformType = "macMDM"
	DeviceComplianceSettingStatePlatformType_Nokia             DeviceComplianceSettingStatePlatformType = "nokia"
	DeviceComplianceSettingStatePlatformType_Palm              DeviceComplianceSettingStatePlatformType = "palm"
	DeviceComplianceSettingStatePlatformType_SurfaceHub        DeviceComplianceSettingStatePlatformType = "surfaceHub"
	DeviceComplianceSettingStatePlatformType_Unix              DeviceComplianceSettingStatePlatformType = "unix"
	DeviceComplianceSettingStatePlatformType_Unknown           DeviceComplianceSettingStatePlatformType = "unknown"
	DeviceComplianceSettingStatePlatformType_WinCE             DeviceComplianceSettingStatePlatformType = "winCE"
	DeviceComplianceSettingStatePlatformType_WinEmbedded       DeviceComplianceSettingStatePlatformType = "winEmbedded"
	DeviceComplianceSettingStatePlatformType_WinMO6            DeviceComplianceSettingStatePlatformType = "winMO6"
	DeviceComplianceSettingStatePlatformType_Windows10x        DeviceComplianceSettingStatePlatformType = "windows10x"
	DeviceComplianceSettingStatePlatformType_WindowsPhone      DeviceComplianceSettingStatePlatformType = "windowsPhone"
	DeviceComplianceSettingStatePlatformType_WindowsRT         DeviceComplianceSettingStatePlatformType = "windowsRT"
)

func PossibleValuesForDeviceComplianceSettingStatePlatformType() []string {
	return []string{
		string(DeviceComplianceSettingStatePlatformType_Android),
		string(DeviceComplianceSettingStatePlatformType_AndroidEnterprise),
		string(DeviceComplianceSettingStatePlatformType_AndroidForWork),
		string(DeviceComplianceSettingStatePlatformType_AndroidnGMS),
		string(DeviceComplianceSettingStatePlatformType_Blackberry),
		string(DeviceComplianceSettingStatePlatformType_ChromeOS),
		string(DeviceComplianceSettingStatePlatformType_CloudPC),
		string(DeviceComplianceSettingStatePlatformType_Desktop),
		string(DeviceComplianceSettingStatePlatformType_HoloLens),
		string(DeviceComplianceSettingStatePlatformType_IPad),
		string(DeviceComplianceSettingStatePlatformType_IPhone),
		string(DeviceComplianceSettingStatePlatformType_IPod),
		string(DeviceComplianceSettingStatePlatformType_ISocConsumer),
		string(DeviceComplianceSettingStatePlatformType_Linux),
		string(DeviceComplianceSettingStatePlatformType_Mac),
		string(DeviceComplianceSettingStatePlatformType_MacMDM),
		string(DeviceComplianceSettingStatePlatformType_Nokia),
		string(DeviceComplianceSettingStatePlatformType_Palm),
		string(DeviceComplianceSettingStatePlatformType_SurfaceHub),
		string(DeviceComplianceSettingStatePlatformType_Unix),
		string(DeviceComplianceSettingStatePlatformType_Unknown),
		string(DeviceComplianceSettingStatePlatformType_WinCE),
		string(DeviceComplianceSettingStatePlatformType_WinEmbedded),
		string(DeviceComplianceSettingStatePlatformType_WinMO6),
		string(DeviceComplianceSettingStatePlatformType_Windows10x),
		string(DeviceComplianceSettingStatePlatformType_WindowsPhone),
		string(DeviceComplianceSettingStatePlatformType_WindowsRT),
	}
}

func (s *DeviceComplianceSettingStatePlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceComplianceSettingStatePlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceComplianceSettingStatePlatformType(input string) (*DeviceComplianceSettingStatePlatformType, error) {
	vals := map[string]DeviceComplianceSettingStatePlatformType{
		"android":           DeviceComplianceSettingStatePlatformType_Android,
		"androidenterprise": DeviceComplianceSettingStatePlatformType_AndroidEnterprise,
		"androidforwork":    DeviceComplianceSettingStatePlatformType_AndroidForWork,
		"androidngms":       DeviceComplianceSettingStatePlatformType_AndroidnGMS,
		"blackberry":        DeviceComplianceSettingStatePlatformType_Blackberry,
		"chromeos":          DeviceComplianceSettingStatePlatformType_ChromeOS,
		"cloudpc":           DeviceComplianceSettingStatePlatformType_CloudPC,
		"desktop":           DeviceComplianceSettingStatePlatformType_Desktop,
		"hololens":          DeviceComplianceSettingStatePlatformType_HoloLens,
		"ipad":              DeviceComplianceSettingStatePlatformType_IPad,
		"iphone":            DeviceComplianceSettingStatePlatformType_IPhone,
		"ipod":              DeviceComplianceSettingStatePlatformType_IPod,
		"isocconsumer":      DeviceComplianceSettingStatePlatformType_ISocConsumer,
		"linux":             DeviceComplianceSettingStatePlatformType_Linux,
		"mac":               DeviceComplianceSettingStatePlatformType_Mac,
		"macmdm":            DeviceComplianceSettingStatePlatformType_MacMDM,
		"nokia":             DeviceComplianceSettingStatePlatformType_Nokia,
		"palm":              DeviceComplianceSettingStatePlatformType_Palm,
		"surfacehub":        DeviceComplianceSettingStatePlatformType_SurfaceHub,
		"unix":              DeviceComplianceSettingStatePlatformType_Unix,
		"unknown":           DeviceComplianceSettingStatePlatformType_Unknown,
		"wince":             DeviceComplianceSettingStatePlatformType_WinCE,
		"winembedded":       DeviceComplianceSettingStatePlatformType_WinEmbedded,
		"winmo6":            DeviceComplianceSettingStatePlatformType_WinMO6,
		"windows10x":        DeviceComplianceSettingStatePlatformType_Windows10x,
		"windowsphone":      DeviceComplianceSettingStatePlatformType_WindowsPhone,
		"windowsrt":         DeviceComplianceSettingStatePlatformType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceSettingStatePlatformType(input)
	return &out, nil
}
