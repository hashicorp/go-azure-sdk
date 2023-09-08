package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupportedDeviceTypeType string

const (
	MobileAppSupportedDeviceTypeTypeandroid           MobileAppSupportedDeviceTypeType = "Android"
	MobileAppSupportedDeviceTypeTypeandroidEnterprise MobileAppSupportedDeviceTypeType = "AndroidEnterprise"
	MobileAppSupportedDeviceTypeTypeandroidForWork    MobileAppSupportedDeviceTypeType = "AndroidForWork"
	MobileAppSupportedDeviceTypeTypeandroidnGMS       MobileAppSupportedDeviceTypeType = "AndroidnGMS"
	MobileAppSupportedDeviceTypeTypeblackberry        MobileAppSupportedDeviceTypeType = "Blackberry"
	MobileAppSupportedDeviceTypeTypechromeOS          MobileAppSupportedDeviceTypeType = "ChromeOS"
	MobileAppSupportedDeviceTypeTypecloudPC           MobileAppSupportedDeviceTypeType = "CloudPC"
	MobileAppSupportedDeviceTypeTypedesktop           MobileAppSupportedDeviceTypeType = "Desktop"
	MobileAppSupportedDeviceTypeTypeholoLens          MobileAppSupportedDeviceTypeType = "HoloLens"
	MobileAppSupportedDeviceTypeTypeiPad              MobileAppSupportedDeviceTypeType = "IPad"
	MobileAppSupportedDeviceTypeTypeiPhone            MobileAppSupportedDeviceTypeType = "IPhone"
	MobileAppSupportedDeviceTypeTypeiPod              MobileAppSupportedDeviceTypeType = "IPod"
	MobileAppSupportedDeviceTypeTypeiSocConsumer      MobileAppSupportedDeviceTypeType = "ISocConsumer"
	MobileAppSupportedDeviceTypeTypelinux             MobileAppSupportedDeviceTypeType = "Linux"
	MobileAppSupportedDeviceTypeTypemac               MobileAppSupportedDeviceTypeType = "Mac"
	MobileAppSupportedDeviceTypeTypemacMDM            MobileAppSupportedDeviceTypeType = "MacMDM"
	MobileAppSupportedDeviceTypeTypenokia             MobileAppSupportedDeviceTypeType = "Nokia"
	MobileAppSupportedDeviceTypeTypepalm              MobileAppSupportedDeviceTypeType = "Palm"
	MobileAppSupportedDeviceTypeTypesurfaceHub        MobileAppSupportedDeviceTypeType = "SurfaceHub"
	MobileAppSupportedDeviceTypeTypeunix              MobileAppSupportedDeviceTypeType = "Unix"
	MobileAppSupportedDeviceTypeTypeunknown           MobileAppSupportedDeviceTypeType = "Unknown"
	MobileAppSupportedDeviceTypeTypewinCE             MobileAppSupportedDeviceTypeType = "WinCE"
	MobileAppSupportedDeviceTypeTypewinEmbedded       MobileAppSupportedDeviceTypeType = "WinEmbedded"
	MobileAppSupportedDeviceTypeTypewinMO6            MobileAppSupportedDeviceTypeType = "WinMO6"
	MobileAppSupportedDeviceTypeTypewindows10x        MobileAppSupportedDeviceTypeType = "Windows10x"
	MobileAppSupportedDeviceTypeTypewindowsPhone      MobileAppSupportedDeviceTypeType = "WindowsPhone"
	MobileAppSupportedDeviceTypeTypewindowsRT         MobileAppSupportedDeviceTypeType = "WindowsRT"
)

func PossibleValuesForMobileAppSupportedDeviceTypeType() []string {
	return []string{
		string(MobileAppSupportedDeviceTypeTypeandroid),
		string(MobileAppSupportedDeviceTypeTypeandroidEnterprise),
		string(MobileAppSupportedDeviceTypeTypeandroidForWork),
		string(MobileAppSupportedDeviceTypeTypeandroidnGMS),
		string(MobileAppSupportedDeviceTypeTypeblackberry),
		string(MobileAppSupportedDeviceTypeTypechromeOS),
		string(MobileAppSupportedDeviceTypeTypecloudPC),
		string(MobileAppSupportedDeviceTypeTypedesktop),
		string(MobileAppSupportedDeviceTypeTypeholoLens),
		string(MobileAppSupportedDeviceTypeTypeiPad),
		string(MobileAppSupportedDeviceTypeTypeiPhone),
		string(MobileAppSupportedDeviceTypeTypeiPod),
		string(MobileAppSupportedDeviceTypeTypeiSocConsumer),
		string(MobileAppSupportedDeviceTypeTypelinux),
		string(MobileAppSupportedDeviceTypeTypemac),
		string(MobileAppSupportedDeviceTypeTypemacMDM),
		string(MobileAppSupportedDeviceTypeTypenokia),
		string(MobileAppSupportedDeviceTypeTypepalm),
		string(MobileAppSupportedDeviceTypeTypesurfaceHub),
		string(MobileAppSupportedDeviceTypeTypeunix),
		string(MobileAppSupportedDeviceTypeTypeunknown),
		string(MobileAppSupportedDeviceTypeTypewinCE),
		string(MobileAppSupportedDeviceTypeTypewinEmbedded),
		string(MobileAppSupportedDeviceTypeTypewinMO6),
		string(MobileAppSupportedDeviceTypeTypewindows10x),
		string(MobileAppSupportedDeviceTypeTypewindowsPhone),
		string(MobileAppSupportedDeviceTypeTypewindowsRT),
	}
}

func parseMobileAppSupportedDeviceTypeType(input string) (*MobileAppSupportedDeviceTypeType, error) {
	vals := map[string]MobileAppSupportedDeviceTypeType{
		"android":           MobileAppSupportedDeviceTypeTypeandroid,
		"androidenterprise": MobileAppSupportedDeviceTypeTypeandroidEnterprise,
		"androidforwork":    MobileAppSupportedDeviceTypeTypeandroidForWork,
		"androidngms":       MobileAppSupportedDeviceTypeTypeandroidnGMS,
		"blackberry":        MobileAppSupportedDeviceTypeTypeblackberry,
		"chromeos":          MobileAppSupportedDeviceTypeTypechromeOS,
		"cloudpc":           MobileAppSupportedDeviceTypeTypecloudPC,
		"desktop":           MobileAppSupportedDeviceTypeTypedesktop,
		"hololens":          MobileAppSupportedDeviceTypeTypeholoLens,
		"ipad":              MobileAppSupportedDeviceTypeTypeiPad,
		"iphone":            MobileAppSupportedDeviceTypeTypeiPhone,
		"ipod":              MobileAppSupportedDeviceTypeTypeiPod,
		"isocconsumer":      MobileAppSupportedDeviceTypeTypeiSocConsumer,
		"linux":             MobileAppSupportedDeviceTypeTypelinux,
		"mac":               MobileAppSupportedDeviceTypeTypemac,
		"macmdm":            MobileAppSupportedDeviceTypeTypemacMDM,
		"nokia":             MobileAppSupportedDeviceTypeTypenokia,
		"palm":              MobileAppSupportedDeviceTypeTypepalm,
		"surfacehub":        MobileAppSupportedDeviceTypeTypesurfaceHub,
		"unix":              MobileAppSupportedDeviceTypeTypeunix,
		"unknown":           MobileAppSupportedDeviceTypeTypeunknown,
		"wince":             MobileAppSupportedDeviceTypeTypewinCE,
		"winembedded":       MobileAppSupportedDeviceTypeTypewinEmbedded,
		"winmo6":            MobileAppSupportedDeviceTypeTypewinMO6,
		"windows10x":        MobileAppSupportedDeviceTypeTypewindows10x,
		"windowsphone":      MobileAppSupportedDeviceTypeTypewindowsPhone,
		"windowsrt":         MobileAppSupportedDeviceTypeTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupportedDeviceTypeType(input)
	return &out, nil
}
