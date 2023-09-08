package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceDeviceType string

const (
	ComanagementEligibleDeviceDeviceTypeandroid           ComanagementEligibleDeviceDeviceType = "Android"
	ComanagementEligibleDeviceDeviceTypeandroidEnterprise ComanagementEligibleDeviceDeviceType = "AndroidEnterprise"
	ComanagementEligibleDeviceDeviceTypeandroidForWork    ComanagementEligibleDeviceDeviceType = "AndroidForWork"
	ComanagementEligibleDeviceDeviceTypeandroidnGMS       ComanagementEligibleDeviceDeviceType = "AndroidnGMS"
	ComanagementEligibleDeviceDeviceTypeblackberry        ComanagementEligibleDeviceDeviceType = "Blackberry"
	ComanagementEligibleDeviceDeviceTypechromeOS          ComanagementEligibleDeviceDeviceType = "ChromeOS"
	ComanagementEligibleDeviceDeviceTypecloudPC           ComanagementEligibleDeviceDeviceType = "CloudPC"
	ComanagementEligibleDeviceDeviceTypedesktop           ComanagementEligibleDeviceDeviceType = "Desktop"
	ComanagementEligibleDeviceDeviceTypeholoLens          ComanagementEligibleDeviceDeviceType = "HoloLens"
	ComanagementEligibleDeviceDeviceTypeiPad              ComanagementEligibleDeviceDeviceType = "IPad"
	ComanagementEligibleDeviceDeviceTypeiPhone            ComanagementEligibleDeviceDeviceType = "IPhone"
	ComanagementEligibleDeviceDeviceTypeiPod              ComanagementEligibleDeviceDeviceType = "IPod"
	ComanagementEligibleDeviceDeviceTypeiSocConsumer      ComanagementEligibleDeviceDeviceType = "ISocConsumer"
	ComanagementEligibleDeviceDeviceTypelinux             ComanagementEligibleDeviceDeviceType = "Linux"
	ComanagementEligibleDeviceDeviceTypemac               ComanagementEligibleDeviceDeviceType = "Mac"
	ComanagementEligibleDeviceDeviceTypemacMDM            ComanagementEligibleDeviceDeviceType = "MacMDM"
	ComanagementEligibleDeviceDeviceTypenokia             ComanagementEligibleDeviceDeviceType = "Nokia"
	ComanagementEligibleDeviceDeviceTypepalm              ComanagementEligibleDeviceDeviceType = "Palm"
	ComanagementEligibleDeviceDeviceTypesurfaceHub        ComanagementEligibleDeviceDeviceType = "SurfaceHub"
	ComanagementEligibleDeviceDeviceTypeunix              ComanagementEligibleDeviceDeviceType = "Unix"
	ComanagementEligibleDeviceDeviceTypeunknown           ComanagementEligibleDeviceDeviceType = "Unknown"
	ComanagementEligibleDeviceDeviceTypewinCE             ComanagementEligibleDeviceDeviceType = "WinCE"
	ComanagementEligibleDeviceDeviceTypewinEmbedded       ComanagementEligibleDeviceDeviceType = "WinEmbedded"
	ComanagementEligibleDeviceDeviceTypewinMO6            ComanagementEligibleDeviceDeviceType = "WinMO6"
	ComanagementEligibleDeviceDeviceTypewindows10x        ComanagementEligibleDeviceDeviceType = "Windows10x"
	ComanagementEligibleDeviceDeviceTypewindowsPhone      ComanagementEligibleDeviceDeviceType = "WindowsPhone"
	ComanagementEligibleDeviceDeviceTypewindowsRT         ComanagementEligibleDeviceDeviceType = "WindowsRT"
)

func PossibleValuesForComanagementEligibleDeviceDeviceType() []string {
	return []string{
		string(ComanagementEligibleDeviceDeviceTypeandroid),
		string(ComanagementEligibleDeviceDeviceTypeandroidEnterprise),
		string(ComanagementEligibleDeviceDeviceTypeandroidForWork),
		string(ComanagementEligibleDeviceDeviceTypeandroidnGMS),
		string(ComanagementEligibleDeviceDeviceTypeblackberry),
		string(ComanagementEligibleDeviceDeviceTypechromeOS),
		string(ComanagementEligibleDeviceDeviceTypecloudPC),
		string(ComanagementEligibleDeviceDeviceTypedesktop),
		string(ComanagementEligibleDeviceDeviceTypeholoLens),
		string(ComanagementEligibleDeviceDeviceTypeiPad),
		string(ComanagementEligibleDeviceDeviceTypeiPhone),
		string(ComanagementEligibleDeviceDeviceTypeiPod),
		string(ComanagementEligibleDeviceDeviceTypeiSocConsumer),
		string(ComanagementEligibleDeviceDeviceTypelinux),
		string(ComanagementEligibleDeviceDeviceTypemac),
		string(ComanagementEligibleDeviceDeviceTypemacMDM),
		string(ComanagementEligibleDeviceDeviceTypenokia),
		string(ComanagementEligibleDeviceDeviceTypepalm),
		string(ComanagementEligibleDeviceDeviceTypesurfaceHub),
		string(ComanagementEligibleDeviceDeviceTypeunix),
		string(ComanagementEligibleDeviceDeviceTypeunknown),
		string(ComanagementEligibleDeviceDeviceTypewinCE),
		string(ComanagementEligibleDeviceDeviceTypewinEmbedded),
		string(ComanagementEligibleDeviceDeviceTypewinMO6),
		string(ComanagementEligibleDeviceDeviceTypewindows10x),
		string(ComanagementEligibleDeviceDeviceTypewindowsPhone),
		string(ComanagementEligibleDeviceDeviceTypewindowsRT),
	}
}

func parseComanagementEligibleDeviceDeviceType(input string) (*ComanagementEligibleDeviceDeviceType, error) {
	vals := map[string]ComanagementEligibleDeviceDeviceType{
		"android":           ComanagementEligibleDeviceDeviceTypeandroid,
		"androidenterprise": ComanagementEligibleDeviceDeviceTypeandroidEnterprise,
		"androidforwork":    ComanagementEligibleDeviceDeviceTypeandroidForWork,
		"androidngms":       ComanagementEligibleDeviceDeviceTypeandroidnGMS,
		"blackberry":        ComanagementEligibleDeviceDeviceTypeblackberry,
		"chromeos":          ComanagementEligibleDeviceDeviceTypechromeOS,
		"cloudpc":           ComanagementEligibleDeviceDeviceTypecloudPC,
		"desktop":           ComanagementEligibleDeviceDeviceTypedesktop,
		"hololens":          ComanagementEligibleDeviceDeviceTypeholoLens,
		"ipad":              ComanagementEligibleDeviceDeviceTypeiPad,
		"iphone":            ComanagementEligibleDeviceDeviceTypeiPhone,
		"ipod":              ComanagementEligibleDeviceDeviceTypeiPod,
		"isocconsumer":      ComanagementEligibleDeviceDeviceTypeiSocConsumer,
		"linux":             ComanagementEligibleDeviceDeviceTypelinux,
		"mac":               ComanagementEligibleDeviceDeviceTypemac,
		"macmdm":            ComanagementEligibleDeviceDeviceTypemacMDM,
		"nokia":             ComanagementEligibleDeviceDeviceTypenokia,
		"palm":              ComanagementEligibleDeviceDeviceTypepalm,
		"surfacehub":        ComanagementEligibleDeviceDeviceTypesurfaceHub,
		"unix":              ComanagementEligibleDeviceDeviceTypeunix,
		"unknown":           ComanagementEligibleDeviceDeviceTypeunknown,
		"wince":             ComanagementEligibleDeviceDeviceTypewinCE,
		"winembedded":       ComanagementEligibleDeviceDeviceTypewinEmbedded,
		"winmo6":            ComanagementEligibleDeviceDeviceTypewinMO6,
		"windows10x":        ComanagementEligibleDeviceDeviceTypewindows10x,
		"windowsphone":      ComanagementEligibleDeviceDeviceTypewindowsPhone,
		"windowsrt":         ComanagementEligibleDeviceDeviceTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceDeviceType(input)
	return &out, nil
}
