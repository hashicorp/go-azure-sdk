package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceDeviceType string

const (
	WindowsManagedDeviceDeviceType_Android           WindowsManagedDeviceDeviceType = "android"
	WindowsManagedDeviceDeviceType_AndroidEnterprise WindowsManagedDeviceDeviceType = "androidEnterprise"
	WindowsManagedDeviceDeviceType_AndroidForWork    WindowsManagedDeviceDeviceType = "androidForWork"
	WindowsManagedDeviceDeviceType_AndroidnGMS       WindowsManagedDeviceDeviceType = "androidnGMS"
	WindowsManagedDeviceDeviceType_Blackberry        WindowsManagedDeviceDeviceType = "blackberry"
	WindowsManagedDeviceDeviceType_ChromeOS          WindowsManagedDeviceDeviceType = "chromeOS"
	WindowsManagedDeviceDeviceType_CloudPC           WindowsManagedDeviceDeviceType = "cloudPC"
	WindowsManagedDeviceDeviceType_Desktop           WindowsManagedDeviceDeviceType = "desktop"
	WindowsManagedDeviceDeviceType_HoloLens          WindowsManagedDeviceDeviceType = "holoLens"
	WindowsManagedDeviceDeviceType_IPad              WindowsManagedDeviceDeviceType = "iPad"
	WindowsManagedDeviceDeviceType_IPhone            WindowsManagedDeviceDeviceType = "iPhone"
	WindowsManagedDeviceDeviceType_IPod              WindowsManagedDeviceDeviceType = "iPod"
	WindowsManagedDeviceDeviceType_ISocConsumer      WindowsManagedDeviceDeviceType = "iSocConsumer"
	WindowsManagedDeviceDeviceType_Linux             WindowsManagedDeviceDeviceType = "linux"
	WindowsManagedDeviceDeviceType_Mac               WindowsManagedDeviceDeviceType = "mac"
	WindowsManagedDeviceDeviceType_MacMDM            WindowsManagedDeviceDeviceType = "macMDM"
	WindowsManagedDeviceDeviceType_Nokia             WindowsManagedDeviceDeviceType = "nokia"
	WindowsManagedDeviceDeviceType_Palm              WindowsManagedDeviceDeviceType = "palm"
	WindowsManagedDeviceDeviceType_SurfaceHub        WindowsManagedDeviceDeviceType = "surfaceHub"
	WindowsManagedDeviceDeviceType_Unix              WindowsManagedDeviceDeviceType = "unix"
	WindowsManagedDeviceDeviceType_Unknown           WindowsManagedDeviceDeviceType = "unknown"
	WindowsManagedDeviceDeviceType_WinCE             WindowsManagedDeviceDeviceType = "winCE"
	WindowsManagedDeviceDeviceType_WinEmbedded       WindowsManagedDeviceDeviceType = "winEmbedded"
	WindowsManagedDeviceDeviceType_WinMO6            WindowsManagedDeviceDeviceType = "winMO6"
	WindowsManagedDeviceDeviceType_Windows10x        WindowsManagedDeviceDeviceType = "windows10x"
	WindowsManagedDeviceDeviceType_WindowsPhone      WindowsManagedDeviceDeviceType = "windowsPhone"
	WindowsManagedDeviceDeviceType_WindowsRT         WindowsManagedDeviceDeviceType = "windowsRT"
)

func PossibleValuesForWindowsManagedDeviceDeviceType() []string {
	return []string{
		string(WindowsManagedDeviceDeviceType_Android),
		string(WindowsManagedDeviceDeviceType_AndroidEnterprise),
		string(WindowsManagedDeviceDeviceType_AndroidForWork),
		string(WindowsManagedDeviceDeviceType_AndroidnGMS),
		string(WindowsManagedDeviceDeviceType_Blackberry),
		string(WindowsManagedDeviceDeviceType_ChromeOS),
		string(WindowsManagedDeviceDeviceType_CloudPC),
		string(WindowsManagedDeviceDeviceType_Desktop),
		string(WindowsManagedDeviceDeviceType_HoloLens),
		string(WindowsManagedDeviceDeviceType_IPad),
		string(WindowsManagedDeviceDeviceType_IPhone),
		string(WindowsManagedDeviceDeviceType_IPod),
		string(WindowsManagedDeviceDeviceType_ISocConsumer),
		string(WindowsManagedDeviceDeviceType_Linux),
		string(WindowsManagedDeviceDeviceType_Mac),
		string(WindowsManagedDeviceDeviceType_MacMDM),
		string(WindowsManagedDeviceDeviceType_Nokia),
		string(WindowsManagedDeviceDeviceType_Palm),
		string(WindowsManagedDeviceDeviceType_SurfaceHub),
		string(WindowsManagedDeviceDeviceType_Unix),
		string(WindowsManagedDeviceDeviceType_Unknown),
		string(WindowsManagedDeviceDeviceType_WinCE),
		string(WindowsManagedDeviceDeviceType_WinEmbedded),
		string(WindowsManagedDeviceDeviceType_WinMO6),
		string(WindowsManagedDeviceDeviceType_Windows10x),
		string(WindowsManagedDeviceDeviceType_WindowsPhone),
		string(WindowsManagedDeviceDeviceType_WindowsRT),
	}
}

func (s *WindowsManagedDeviceDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceDeviceType(input string) (*WindowsManagedDeviceDeviceType, error) {
	vals := map[string]WindowsManagedDeviceDeviceType{
		"android":           WindowsManagedDeviceDeviceType_Android,
		"androidenterprise": WindowsManagedDeviceDeviceType_AndroidEnterprise,
		"androidforwork":    WindowsManagedDeviceDeviceType_AndroidForWork,
		"androidngms":       WindowsManagedDeviceDeviceType_AndroidnGMS,
		"blackberry":        WindowsManagedDeviceDeviceType_Blackberry,
		"chromeos":          WindowsManagedDeviceDeviceType_ChromeOS,
		"cloudpc":           WindowsManagedDeviceDeviceType_CloudPC,
		"desktop":           WindowsManagedDeviceDeviceType_Desktop,
		"hololens":          WindowsManagedDeviceDeviceType_HoloLens,
		"ipad":              WindowsManagedDeviceDeviceType_IPad,
		"iphone":            WindowsManagedDeviceDeviceType_IPhone,
		"ipod":              WindowsManagedDeviceDeviceType_IPod,
		"isocconsumer":      WindowsManagedDeviceDeviceType_ISocConsumer,
		"linux":             WindowsManagedDeviceDeviceType_Linux,
		"mac":               WindowsManagedDeviceDeviceType_Mac,
		"macmdm":            WindowsManagedDeviceDeviceType_MacMDM,
		"nokia":             WindowsManagedDeviceDeviceType_Nokia,
		"palm":              WindowsManagedDeviceDeviceType_Palm,
		"surfacehub":        WindowsManagedDeviceDeviceType_SurfaceHub,
		"unix":              WindowsManagedDeviceDeviceType_Unix,
		"unknown":           WindowsManagedDeviceDeviceType_Unknown,
		"wince":             WindowsManagedDeviceDeviceType_WinCE,
		"winembedded":       WindowsManagedDeviceDeviceType_WinEmbedded,
		"winmo6":            WindowsManagedDeviceDeviceType_WinMO6,
		"windows10x":        WindowsManagedDeviceDeviceType_Windows10x,
		"windowsphone":      WindowsManagedDeviceDeviceType_WindowsPhone,
		"windowsrt":         WindowsManagedDeviceDeviceType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceDeviceType(input)
	return &out, nil
}
