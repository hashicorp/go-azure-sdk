package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetireScheduledManagedDeviceDeviceType string

const (
	RetireScheduledManagedDeviceDeviceTypeandroid           RetireScheduledManagedDeviceDeviceType = "Android"
	RetireScheduledManagedDeviceDeviceTypeandroidEnterprise RetireScheduledManagedDeviceDeviceType = "AndroidEnterprise"
	RetireScheduledManagedDeviceDeviceTypeandroidForWork    RetireScheduledManagedDeviceDeviceType = "AndroidForWork"
	RetireScheduledManagedDeviceDeviceTypeandroidnGMS       RetireScheduledManagedDeviceDeviceType = "AndroidnGMS"
	RetireScheduledManagedDeviceDeviceTypeblackberry        RetireScheduledManagedDeviceDeviceType = "Blackberry"
	RetireScheduledManagedDeviceDeviceTypechromeOS          RetireScheduledManagedDeviceDeviceType = "ChromeOS"
	RetireScheduledManagedDeviceDeviceTypecloudPC           RetireScheduledManagedDeviceDeviceType = "CloudPC"
	RetireScheduledManagedDeviceDeviceTypedesktop           RetireScheduledManagedDeviceDeviceType = "Desktop"
	RetireScheduledManagedDeviceDeviceTypeholoLens          RetireScheduledManagedDeviceDeviceType = "HoloLens"
	RetireScheduledManagedDeviceDeviceTypeiPad              RetireScheduledManagedDeviceDeviceType = "IPad"
	RetireScheduledManagedDeviceDeviceTypeiPhone            RetireScheduledManagedDeviceDeviceType = "IPhone"
	RetireScheduledManagedDeviceDeviceTypeiPod              RetireScheduledManagedDeviceDeviceType = "IPod"
	RetireScheduledManagedDeviceDeviceTypeiSocConsumer      RetireScheduledManagedDeviceDeviceType = "ISocConsumer"
	RetireScheduledManagedDeviceDeviceTypelinux             RetireScheduledManagedDeviceDeviceType = "Linux"
	RetireScheduledManagedDeviceDeviceTypemac               RetireScheduledManagedDeviceDeviceType = "Mac"
	RetireScheduledManagedDeviceDeviceTypemacMDM            RetireScheduledManagedDeviceDeviceType = "MacMDM"
	RetireScheduledManagedDeviceDeviceTypenokia             RetireScheduledManagedDeviceDeviceType = "Nokia"
	RetireScheduledManagedDeviceDeviceTypepalm              RetireScheduledManagedDeviceDeviceType = "Palm"
	RetireScheduledManagedDeviceDeviceTypesurfaceHub        RetireScheduledManagedDeviceDeviceType = "SurfaceHub"
	RetireScheduledManagedDeviceDeviceTypeunix              RetireScheduledManagedDeviceDeviceType = "Unix"
	RetireScheduledManagedDeviceDeviceTypeunknown           RetireScheduledManagedDeviceDeviceType = "Unknown"
	RetireScheduledManagedDeviceDeviceTypewinCE             RetireScheduledManagedDeviceDeviceType = "WinCE"
	RetireScheduledManagedDeviceDeviceTypewinEmbedded       RetireScheduledManagedDeviceDeviceType = "WinEmbedded"
	RetireScheduledManagedDeviceDeviceTypewinMO6            RetireScheduledManagedDeviceDeviceType = "WinMO6"
	RetireScheduledManagedDeviceDeviceTypewindows10x        RetireScheduledManagedDeviceDeviceType = "Windows10x"
	RetireScheduledManagedDeviceDeviceTypewindowsPhone      RetireScheduledManagedDeviceDeviceType = "WindowsPhone"
	RetireScheduledManagedDeviceDeviceTypewindowsRT         RetireScheduledManagedDeviceDeviceType = "WindowsRT"
)

func PossibleValuesForRetireScheduledManagedDeviceDeviceType() []string {
	return []string{
		string(RetireScheduledManagedDeviceDeviceTypeandroid),
		string(RetireScheduledManagedDeviceDeviceTypeandroidEnterprise),
		string(RetireScheduledManagedDeviceDeviceTypeandroidForWork),
		string(RetireScheduledManagedDeviceDeviceTypeandroidnGMS),
		string(RetireScheduledManagedDeviceDeviceTypeblackberry),
		string(RetireScheduledManagedDeviceDeviceTypechromeOS),
		string(RetireScheduledManagedDeviceDeviceTypecloudPC),
		string(RetireScheduledManagedDeviceDeviceTypedesktop),
		string(RetireScheduledManagedDeviceDeviceTypeholoLens),
		string(RetireScheduledManagedDeviceDeviceTypeiPad),
		string(RetireScheduledManagedDeviceDeviceTypeiPhone),
		string(RetireScheduledManagedDeviceDeviceTypeiPod),
		string(RetireScheduledManagedDeviceDeviceTypeiSocConsumer),
		string(RetireScheduledManagedDeviceDeviceTypelinux),
		string(RetireScheduledManagedDeviceDeviceTypemac),
		string(RetireScheduledManagedDeviceDeviceTypemacMDM),
		string(RetireScheduledManagedDeviceDeviceTypenokia),
		string(RetireScheduledManagedDeviceDeviceTypepalm),
		string(RetireScheduledManagedDeviceDeviceTypesurfaceHub),
		string(RetireScheduledManagedDeviceDeviceTypeunix),
		string(RetireScheduledManagedDeviceDeviceTypeunknown),
		string(RetireScheduledManagedDeviceDeviceTypewinCE),
		string(RetireScheduledManagedDeviceDeviceTypewinEmbedded),
		string(RetireScheduledManagedDeviceDeviceTypewinMO6),
		string(RetireScheduledManagedDeviceDeviceTypewindows10x),
		string(RetireScheduledManagedDeviceDeviceTypewindowsPhone),
		string(RetireScheduledManagedDeviceDeviceTypewindowsRT),
	}
}

func parseRetireScheduledManagedDeviceDeviceType(input string) (*RetireScheduledManagedDeviceDeviceType, error) {
	vals := map[string]RetireScheduledManagedDeviceDeviceType{
		"android":           RetireScheduledManagedDeviceDeviceTypeandroid,
		"androidenterprise": RetireScheduledManagedDeviceDeviceTypeandroidEnterprise,
		"androidforwork":    RetireScheduledManagedDeviceDeviceTypeandroidForWork,
		"androidngms":       RetireScheduledManagedDeviceDeviceTypeandroidnGMS,
		"blackberry":        RetireScheduledManagedDeviceDeviceTypeblackberry,
		"chromeos":          RetireScheduledManagedDeviceDeviceTypechromeOS,
		"cloudpc":           RetireScheduledManagedDeviceDeviceTypecloudPC,
		"desktop":           RetireScheduledManagedDeviceDeviceTypedesktop,
		"hololens":          RetireScheduledManagedDeviceDeviceTypeholoLens,
		"ipad":              RetireScheduledManagedDeviceDeviceTypeiPad,
		"iphone":            RetireScheduledManagedDeviceDeviceTypeiPhone,
		"ipod":              RetireScheduledManagedDeviceDeviceTypeiPod,
		"isocconsumer":      RetireScheduledManagedDeviceDeviceTypeiSocConsumer,
		"linux":             RetireScheduledManagedDeviceDeviceTypelinux,
		"mac":               RetireScheduledManagedDeviceDeviceTypemac,
		"macmdm":            RetireScheduledManagedDeviceDeviceTypemacMDM,
		"nokia":             RetireScheduledManagedDeviceDeviceTypenokia,
		"palm":              RetireScheduledManagedDeviceDeviceTypepalm,
		"surfacehub":        RetireScheduledManagedDeviceDeviceTypesurfaceHub,
		"unix":              RetireScheduledManagedDeviceDeviceTypeunix,
		"unknown":           RetireScheduledManagedDeviceDeviceTypeunknown,
		"wince":             RetireScheduledManagedDeviceDeviceTypewinCE,
		"winembedded":       RetireScheduledManagedDeviceDeviceTypewinEmbedded,
		"winmo6":            RetireScheduledManagedDeviceDeviceTypewinMO6,
		"windows10x":        RetireScheduledManagedDeviceDeviceTypewindows10x,
		"windowsphone":      RetireScheduledManagedDeviceDeviceTypewindowsPhone,
		"windowsrt":         RetireScheduledManagedDeviceDeviceTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetireScheduledManagedDeviceDeviceType(input)
	return &out, nil
}
