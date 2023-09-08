package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType string

const (
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroid           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Android"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidEnterprise AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "AndroidEnterprise"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidForWork    AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "AndroidForWork"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidnGMS       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "AndroidnGMS"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeblackberry        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Blackberry"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypechromeOS          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "ChromeOS"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypecloudPC           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "CloudPC"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypedesktop           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Desktop"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeholoLens          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "HoloLens"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPad              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "IPad"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPhone            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "IPhone"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPod              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "IPod"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiSocConsumer      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "ISocConsumer"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypelinux             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Linux"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypemac               AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Mac"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypemacMDM            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "MacMDM"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypenokia             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Nokia"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypepalm              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Palm"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypesurfaceHub        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "SurfaceHub"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeunix              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Unix"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeunknown           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Unknown"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinCE             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "WinCE"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinEmbedded       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "WinEmbedded"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinMO6            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "WinMO6"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindows10x        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "Windows10x"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindowsPhone      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "WindowsPhone"
	AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindowsRT         AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType = "WindowsRT"
)

func PossibleValuesForAdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType() []string {
	return []string{
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroid),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidEnterprise),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidForWork),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidnGMS),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeblackberry),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypechromeOS),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypecloudPC),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypedesktop),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeholoLens),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPad),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPhone),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPod),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiSocConsumer),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypelinux),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypemac),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypemacMDM),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypenokia),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypepalm),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypesurfaceHub),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeunix),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeunknown),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinCE),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinEmbedded),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinMO6),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindows10x),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindowsPhone),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindowsRT),
	}
}

func parseAdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType(input string) (*AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType, error) {
	vals := map[string]AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType{
		"android":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroid,
		"androidenterprise": AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidEnterprise,
		"androidforwork":    AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidForWork,
		"androidngms":       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeandroidnGMS,
		"blackberry":        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeblackberry,
		"chromeos":          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypechromeOS,
		"cloudpc":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypecloudPC,
		"desktop":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypedesktop,
		"hololens":          AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeholoLens,
		"ipad":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPad,
		"iphone":            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPhone,
		"ipod":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiPod,
		"isocconsumer":      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeiSocConsumer,
		"linux":             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypelinux,
		"mac":               AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypemac,
		"macmdm":            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypemacMDM,
		"nokia":             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypenokia,
		"palm":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypepalm,
		"surfacehub":        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypesurfaceHub,
		"unix":              AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeunix,
		"unknown":           AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypeunknown,
		"wince":             AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinCE,
		"winembedded":       AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinEmbedded,
		"winmo6":            AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewinMO6,
		"windows10x":        AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindows10x,
		"windowsphone":      AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindowsPhone,
		"windowsrt":         AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedThreatProtectionOnboardingDeviceSettingStatePlatformType(input)
	return &out, nil
}
