package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceDeviceType string

const (
	RetireScheduledManagedDeviceDeviceType_Android           RetireScheduledManagedDeviceDeviceType = "android"
	RetireScheduledManagedDeviceDeviceType_AndroidEnterprise RetireScheduledManagedDeviceDeviceType = "androidEnterprise"
	RetireScheduledManagedDeviceDeviceType_AndroidForWork    RetireScheduledManagedDeviceDeviceType = "androidForWork"
	RetireScheduledManagedDeviceDeviceType_AndroidnGMS       RetireScheduledManagedDeviceDeviceType = "androidnGMS"
	RetireScheduledManagedDeviceDeviceType_Blackberry        RetireScheduledManagedDeviceDeviceType = "blackberry"
	RetireScheduledManagedDeviceDeviceType_ChromeOS          RetireScheduledManagedDeviceDeviceType = "chromeOS"
	RetireScheduledManagedDeviceDeviceType_CloudPC           RetireScheduledManagedDeviceDeviceType = "cloudPC"
	RetireScheduledManagedDeviceDeviceType_Desktop           RetireScheduledManagedDeviceDeviceType = "desktop"
	RetireScheduledManagedDeviceDeviceType_HoloLens          RetireScheduledManagedDeviceDeviceType = "holoLens"
	RetireScheduledManagedDeviceDeviceType_IPad              RetireScheduledManagedDeviceDeviceType = "iPad"
	RetireScheduledManagedDeviceDeviceType_IPhone            RetireScheduledManagedDeviceDeviceType = "iPhone"
	RetireScheduledManagedDeviceDeviceType_IPod              RetireScheduledManagedDeviceDeviceType = "iPod"
	RetireScheduledManagedDeviceDeviceType_ISocConsumer      RetireScheduledManagedDeviceDeviceType = "iSocConsumer"
	RetireScheduledManagedDeviceDeviceType_Linux             RetireScheduledManagedDeviceDeviceType = "linux"
	RetireScheduledManagedDeviceDeviceType_Mac               RetireScheduledManagedDeviceDeviceType = "mac"
	RetireScheduledManagedDeviceDeviceType_MacMDM            RetireScheduledManagedDeviceDeviceType = "macMDM"
	RetireScheduledManagedDeviceDeviceType_Nokia             RetireScheduledManagedDeviceDeviceType = "nokia"
	RetireScheduledManagedDeviceDeviceType_Palm              RetireScheduledManagedDeviceDeviceType = "palm"
	RetireScheduledManagedDeviceDeviceType_SurfaceHub        RetireScheduledManagedDeviceDeviceType = "surfaceHub"
	RetireScheduledManagedDeviceDeviceType_Unix              RetireScheduledManagedDeviceDeviceType = "unix"
	RetireScheduledManagedDeviceDeviceType_Unknown           RetireScheduledManagedDeviceDeviceType = "unknown"
	RetireScheduledManagedDeviceDeviceType_WinCE             RetireScheduledManagedDeviceDeviceType = "winCE"
	RetireScheduledManagedDeviceDeviceType_WinEmbedded       RetireScheduledManagedDeviceDeviceType = "winEmbedded"
	RetireScheduledManagedDeviceDeviceType_WinMO6            RetireScheduledManagedDeviceDeviceType = "winMO6"
	RetireScheduledManagedDeviceDeviceType_Windows10x        RetireScheduledManagedDeviceDeviceType = "windows10x"
	RetireScheduledManagedDeviceDeviceType_WindowsPhone      RetireScheduledManagedDeviceDeviceType = "windowsPhone"
	RetireScheduledManagedDeviceDeviceType_WindowsRT         RetireScheduledManagedDeviceDeviceType = "windowsRT"
)

func PossibleValuesForRetireScheduledManagedDeviceDeviceType() []string {
	return []string{
		string(RetireScheduledManagedDeviceDeviceType_Android),
		string(RetireScheduledManagedDeviceDeviceType_AndroidEnterprise),
		string(RetireScheduledManagedDeviceDeviceType_AndroidForWork),
		string(RetireScheduledManagedDeviceDeviceType_AndroidnGMS),
		string(RetireScheduledManagedDeviceDeviceType_Blackberry),
		string(RetireScheduledManagedDeviceDeviceType_ChromeOS),
		string(RetireScheduledManagedDeviceDeviceType_CloudPC),
		string(RetireScheduledManagedDeviceDeviceType_Desktop),
		string(RetireScheduledManagedDeviceDeviceType_HoloLens),
		string(RetireScheduledManagedDeviceDeviceType_IPad),
		string(RetireScheduledManagedDeviceDeviceType_IPhone),
		string(RetireScheduledManagedDeviceDeviceType_IPod),
		string(RetireScheduledManagedDeviceDeviceType_ISocConsumer),
		string(RetireScheduledManagedDeviceDeviceType_Linux),
		string(RetireScheduledManagedDeviceDeviceType_Mac),
		string(RetireScheduledManagedDeviceDeviceType_MacMDM),
		string(RetireScheduledManagedDeviceDeviceType_Nokia),
		string(RetireScheduledManagedDeviceDeviceType_Palm),
		string(RetireScheduledManagedDeviceDeviceType_SurfaceHub),
		string(RetireScheduledManagedDeviceDeviceType_Unix),
		string(RetireScheduledManagedDeviceDeviceType_Unknown),
		string(RetireScheduledManagedDeviceDeviceType_WinCE),
		string(RetireScheduledManagedDeviceDeviceType_WinEmbedded),
		string(RetireScheduledManagedDeviceDeviceType_WinMO6),
		string(RetireScheduledManagedDeviceDeviceType_Windows10x),
		string(RetireScheduledManagedDeviceDeviceType_WindowsPhone),
		string(RetireScheduledManagedDeviceDeviceType_WindowsRT),
	}
}

func (s *RetireScheduledManagedDeviceDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetireScheduledManagedDeviceDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetireScheduledManagedDeviceDeviceType(input string) (*RetireScheduledManagedDeviceDeviceType, error) {
	vals := map[string]RetireScheduledManagedDeviceDeviceType{
		"android":           RetireScheduledManagedDeviceDeviceType_Android,
		"androidenterprise": RetireScheduledManagedDeviceDeviceType_AndroidEnterprise,
		"androidforwork":    RetireScheduledManagedDeviceDeviceType_AndroidForWork,
		"androidngms":       RetireScheduledManagedDeviceDeviceType_AndroidnGMS,
		"blackberry":        RetireScheduledManagedDeviceDeviceType_Blackberry,
		"chromeos":          RetireScheduledManagedDeviceDeviceType_ChromeOS,
		"cloudpc":           RetireScheduledManagedDeviceDeviceType_CloudPC,
		"desktop":           RetireScheduledManagedDeviceDeviceType_Desktop,
		"hololens":          RetireScheduledManagedDeviceDeviceType_HoloLens,
		"ipad":              RetireScheduledManagedDeviceDeviceType_IPad,
		"iphone":            RetireScheduledManagedDeviceDeviceType_IPhone,
		"ipod":              RetireScheduledManagedDeviceDeviceType_IPod,
		"isocconsumer":      RetireScheduledManagedDeviceDeviceType_ISocConsumer,
		"linux":             RetireScheduledManagedDeviceDeviceType_Linux,
		"mac":               RetireScheduledManagedDeviceDeviceType_Mac,
		"macmdm":            RetireScheduledManagedDeviceDeviceType_MacMDM,
		"nokia":             RetireScheduledManagedDeviceDeviceType_Nokia,
		"palm":              RetireScheduledManagedDeviceDeviceType_Palm,
		"surfacehub":        RetireScheduledManagedDeviceDeviceType_SurfaceHub,
		"unix":              RetireScheduledManagedDeviceDeviceType_Unix,
		"unknown":           RetireScheduledManagedDeviceDeviceType_Unknown,
		"wince":             RetireScheduledManagedDeviceDeviceType_WinCE,
		"winembedded":       RetireScheduledManagedDeviceDeviceType_WinEmbedded,
		"winmo6":            RetireScheduledManagedDeviceDeviceType_WinMO6,
		"windows10x":        RetireScheduledManagedDeviceDeviceType_Windows10x,
		"windowsphone":      RetireScheduledManagedDeviceDeviceType_WindowsPhone,
		"windowsrt":         RetireScheduledManagedDeviceDeviceType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetireScheduledManagedDeviceDeviceType(input)
	return &out, nil
}
