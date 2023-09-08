package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceDeviceType string

const (
	ManagedDeviceDeviceTypeandroid           ManagedDeviceDeviceType = "Android"
	ManagedDeviceDeviceTypeandroidEnterprise ManagedDeviceDeviceType = "AndroidEnterprise"
	ManagedDeviceDeviceTypeandroidForWork    ManagedDeviceDeviceType = "AndroidForWork"
	ManagedDeviceDeviceTypeandroidnGMS       ManagedDeviceDeviceType = "AndroidnGMS"
	ManagedDeviceDeviceTypeblackberry        ManagedDeviceDeviceType = "Blackberry"
	ManagedDeviceDeviceTypechromeOS          ManagedDeviceDeviceType = "ChromeOS"
	ManagedDeviceDeviceTypecloudPC           ManagedDeviceDeviceType = "CloudPC"
	ManagedDeviceDeviceTypedesktop           ManagedDeviceDeviceType = "Desktop"
	ManagedDeviceDeviceTypeholoLens          ManagedDeviceDeviceType = "HoloLens"
	ManagedDeviceDeviceTypeiPad              ManagedDeviceDeviceType = "IPad"
	ManagedDeviceDeviceTypeiPhone            ManagedDeviceDeviceType = "IPhone"
	ManagedDeviceDeviceTypeiPod              ManagedDeviceDeviceType = "IPod"
	ManagedDeviceDeviceTypeiSocConsumer      ManagedDeviceDeviceType = "ISocConsumer"
	ManagedDeviceDeviceTypelinux             ManagedDeviceDeviceType = "Linux"
	ManagedDeviceDeviceTypemac               ManagedDeviceDeviceType = "Mac"
	ManagedDeviceDeviceTypemacMDM            ManagedDeviceDeviceType = "MacMDM"
	ManagedDeviceDeviceTypenokia             ManagedDeviceDeviceType = "Nokia"
	ManagedDeviceDeviceTypepalm              ManagedDeviceDeviceType = "Palm"
	ManagedDeviceDeviceTypesurfaceHub        ManagedDeviceDeviceType = "SurfaceHub"
	ManagedDeviceDeviceTypeunix              ManagedDeviceDeviceType = "Unix"
	ManagedDeviceDeviceTypeunknown           ManagedDeviceDeviceType = "Unknown"
	ManagedDeviceDeviceTypewinCE             ManagedDeviceDeviceType = "WinCE"
	ManagedDeviceDeviceTypewinEmbedded       ManagedDeviceDeviceType = "WinEmbedded"
	ManagedDeviceDeviceTypewinMO6            ManagedDeviceDeviceType = "WinMO6"
	ManagedDeviceDeviceTypewindows10x        ManagedDeviceDeviceType = "Windows10x"
	ManagedDeviceDeviceTypewindowsPhone      ManagedDeviceDeviceType = "WindowsPhone"
	ManagedDeviceDeviceTypewindowsRT         ManagedDeviceDeviceType = "WindowsRT"
)

func PossibleValuesForManagedDeviceDeviceType() []string {
	return []string{
		string(ManagedDeviceDeviceTypeandroid),
		string(ManagedDeviceDeviceTypeandroidEnterprise),
		string(ManagedDeviceDeviceTypeandroidForWork),
		string(ManagedDeviceDeviceTypeandroidnGMS),
		string(ManagedDeviceDeviceTypeblackberry),
		string(ManagedDeviceDeviceTypechromeOS),
		string(ManagedDeviceDeviceTypecloudPC),
		string(ManagedDeviceDeviceTypedesktop),
		string(ManagedDeviceDeviceTypeholoLens),
		string(ManagedDeviceDeviceTypeiPad),
		string(ManagedDeviceDeviceTypeiPhone),
		string(ManagedDeviceDeviceTypeiPod),
		string(ManagedDeviceDeviceTypeiSocConsumer),
		string(ManagedDeviceDeviceTypelinux),
		string(ManagedDeviceDeviceTypemac),
		string(ManagedDeviceDeviceTypemacMDM),
		string(ManagedDeviceDeviceTypenokia),
		string(ManagedDeviceDeviceTypepalm),
		string(ManagedDeviceDeviceTypesurfaceHub),
		string(ManagedDeviceDeviceTypeunix),
		string(ManagedDeviceDeviceTypeunknown),
		string(ManagedDeviceDeviceTypewinCE),
		string(ManagedDeviceDeviceTypewinEmbedded),
		string(ManagedDeviceDeviceTypewinMO6),
		string(ManagedDeviceDeviceTypewindows10x),
		string(ManagedDeviceDeviceTypewindowsPhone),
		string(ManagedDeviceDeviceTypewindowsRT),
	}
}

func parseManagedDeviceDeviceType(input string) (*ManagedDeviceDeviceType, error) {
	vals := map[string]ManagedDeviceDeviceType{
		"android":           ManagedDeviceDeviceTypeandroid,
		"androidenterprise": ManagedDeviceDeviceTypeandroidEnterprise,
		"androidforwork":    ManagedDeviceDeviceTypeandroidForWork,
		"androidngms":       ManagedDeviceDeviceTypeandroidnGMS,
		"blackberry":        ManagedDeviceDeviceTypeblackberry,
		"chromeos":          ManagedDeviceDeviceTypechromeOS,
		"cloudpc":           ManagedDeviceDeviceTypecloudPC,
		"desktop":           ManagedDeviceDeviceTypedesktop,
		"hololens":          ManagedDeviceDeviceTypeholoLens,
		"ipad":              ManagedDeviceDeviceTypeiPad,
		"iphone":            ManagedDeviceDeviceTypeiPhone,
		"ipod":              ManagedDeviceDeviceTypeiPod,
		"isocconsumer":      ManagedDeviceDeviceTypeiSocConsumer,
		"linux":             ManagedDeviceDeviceTypelinux,
		"mac":               ManagedDeviceDeviceTypemac,
		"macmdm":            ManagedDeviceDeviceTypemacMDM,
		"nokia":             ManagedDeviceDeviceTypenokia,
		"palm":              ManagedDeviceDeviceTypepalm,
		"surfacehub":        ManagedDeviceDeviceTypesurfaceHub,
		"unix":              ManagedDeviceDeviceTypeunix,
		"unknown":           ManagedDeviceDeviceTypeunknown,
		"wince":             ManagedDeviceDeviceTypewinCE,
		"winembedded":       ManagedDeviceDeviceTypewinEmbedded,
		"winmo6":            ManagedDeviceDeviceTypewinMO6,
		"windows10x":        ManagedDeviceDeviceTypewindows10x,
		"windowsphone":      ManagedDeviceDeviceTypewindowsPhone,
		"windowsrt":         ManagedDeviceDeviceTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceDeviceType(input)
	return &out, nil
}
