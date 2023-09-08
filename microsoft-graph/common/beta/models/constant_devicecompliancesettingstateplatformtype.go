package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceSettingStatePlatformType string

const (
	DeviceComplianceSettingStatePlatformTypeandroid           DeviceComplianceSettingStatePlatformType = "Android"
	DeviceComplianceSettingStatePlatformTypeandroidEnterprise DeviceComplianceSettingStatePlatformType = "AndroidEnterprise"
	DeviceComplianceSettingStatePlatformTypeandroidForWork    DeviceComplianceSettingStatePlatformType = "AndroidForWork"
	DeviceComplianceSettingStatePlatformTypeandroidnGMS       DeviceComplianceSettingStatePlatformType = "AndroidnGMS"
	DeviceComplianceSettingStatePlatformTypeblackberry        DeviceComplianceSettingStatePlatformType = "Blackberry"
	DeviceComplianceSettingStatePlatformTypechromeOS          DeviceComplianceSettingStatePlatformType = "ChromeOS"
	DeviceComplianceSettingStatePlatformTypecloudPC           DeviceComplianceSettingStatePlatformType = "CloudPC"
	DeviceComplianceSettingStatePlatformTypedesktop           DeviceComplianceSettingStatePlatformType = "Desktop"
	DeviceComplianceSettingStatePlatformTypeholoLens          DeviceComplianceSettingStatePlatformType = "HoloLens"
	DeviceComplianceSettingStatePlatformTypeiPad              DeviceComplianceSettingStatePlatformType = "IPad"
	DeviceComplianceSettingStatePlatformTypeiPhone            DeviceComplianceSettingStatePlatformType = "IPhone"
	DeviceComplianceSettingStatePlatformTypeiPod              DeviceComplianceSettingStatePlatformType = "IPod"
	DeviceComplianceSettingStatePlatformTypeiSocConsumer      DeviceComplianceSettingStatePlatformType = "ISocConsumer"
	DeviceComplianceSettingStatePlatformTypelinux             DeviceComplianceSettingStatePlatformType = "Linux"
	DeviceComplianceSettingStatePlatformTypemac               DeviceComplianceSettingStatePlatformType = "Mac"
	DeviceComplianceSettingStatePlatformTypemacMDM            DeviceComplianceSettingStatePlatformType = "MacMDM"
	DeviceComplianceSettingStatePlatformTypenokia             DeviceComplianceSettingStatePlatformType = "Nokia"
	DeviceComplianceSettingStatePlatformTypepalm              DeviceComplianceSettingStatePlatformType = "Palm"
	DeviceComplianceSettingStatePlatformTypesurfaceHub        DeviceComplianceSettingStatePlatformType = "SurfaceHub"
	DeviceComplianceSettingStatePlatformTypeunix              DeviceComplianceSettingStatePlatformType = "Unix"
	DeviceComplianceSettingStatePlatformTypeunknown           DeviceComplianceSettingStatePlatformType = "Unknown"
	DeviceComplianceSettingStatePlatformTypewinCE             DeviceComplianceSettingStatePlatformType = "WinCE"
	DeviceComplianceSettingStatePlatformTypewinEmbedded       DeviceComplianceSettingStatePlatformType = "WinEmbedded"
	DeviceComplianceSettingStatePlatformTypewinMO6            DeviceComplianceSettingStatePlatformType = "WinMO6"
	DeviceComplianceSettingStatePlatformTypewindows10x        DeviceComplianceSettingStatePlatformType = "Windows10x"
	DeviceComplianceSettingStatePlatformTypewindowsPhone      DeviceComplianceSettingStatePlatformType = "WindowsPhone"
	DeviceComplianceSettingStatePlatformTypewindowsRT         DeviceComplianceSettingStatePlatformType = "WindowsRT"
)

func PossibleValuesForDeviceComplianceSettingStatePlatformType() []string {
	return []string{
		string(DeviceComplianceSettingStatePlatformTypeandroid),
		string(DeviceComplianceSettingStatePlatformTypeandroidEnterprise),
		string(DeviceComplianceSettingStatePlatformTypeandroidForWork),
		string(DeviceComplianceSettingStatePlatformTypeandroidnGMS),
		string(DeviceComplianceSettingStatePlatformTypeblackberry),
		string(DeviceComplianceSettingStatePlatformTypechromeOS),
		string(DeviceComplianceSettingStatePlatformTypecloudPC),
		string(DeviceComplianceSettingStatePlatformTypedesktop),
		string(DeviceComplianceSettingStatePlatformTypeholoLens),
		string(DeviceComplianceSettingStatePlatformTypeiPad),
		string(DeviceComplianceSettingStatePlatformTypeiPhone),
		string(DeviceComplianceSettingStatePlatformTypeiPod),
		string(DeviceComplianceSettingStatePlatformTypeiSocConsumer),
		string(DeviceComplianceSettingStatePlatformTypelinux),
		string(DeviceComplianceSettingStatePlatformTypemac),
		string(DeviceComplianceSettingStatePlatformTypemacMDM),
		string(DeviceComplianceSettingStatePlatformTypenokia),
		string(DeviceComplianceSettingStatePlatformTypepalm),
		string(DeviceComplianceSettingStatePlatformTypesurfaceHub),
		string(DeviceComplianceSettingStatePlatformTypeunix),
		string(DeviceComplianceSettingStatePlatformTypeunknown),
		string(DeviceComplianceSettingStatePlatformTypewinCE),
		string(DeviceComplianceSettingStatePlatformTypewinEmbedded),
		string(DeviceComplianceSettingStatePlatformTypewinMO6),
		string(DeviceComplianceSettingStatePlatformTypewindows10x),
		string(DeviceComplianceSettingStatePlatformTypewindowsPhone),
		string(DeviceComplianceSettingStatePlatformTypewindowsRT),
	}
}

func parseDeviceComplianceSettingStatePlatformType(input string) (*DeviceComplianceSettingStatePlatformType, error) {
	vals := map[string]DeviceComplianceSettingStatePlatformType{
		"android":           DeviceComplianceSettingStatePlatformTypeandroid,
		"androidenterprise": DeviceComplianceSettingStatePlatformTypeandroidEnterprise,
		"androidforwork":    DeviceComplianceSettingStatePlatformTypeandroidForWork,
		"androidngms":       DeviceComplianceSettingStatePlatformTypeandroidnGMS,
		"blackberry":        DeviceComplianceSettingStatePlatformTypeblackberry,
		"chromeos":          DeviceComplianceSettingStatePlatformTypechromeOS,
		"cloudpc":           DeviceComplianceSettingStatePlatformTypecloudPC,
		"desktop":           DeviceComplianceSettingStatePlatformTypedesktop,
		"hololens":          DeviceComplianceSettingStatePlatformTypeholoLens,
		"ipad":              DeviceComplianceSettingStatePlatformTypeiPad,
		"iphone":            DeviceComplianceSettingStatePlatformTypeiPhone,
		"ipod":              DeviceComplianceSettingStatePlatformTypeiPod,
		"isocconsumer":      DeviceComplianceSettingStatePlatformTypeiSocConsumer,
		"linux":             DeviceComplianceSettingStatePlatformTypelinux,
		"mac":               DeviceComplianceSettingStatePlatformTypemac,
		"macmdm":            DeviceComplianceSettingStatePlatformTypemacMDM,
		"nokia":             DeviceComplianceSettingStatePlatformTypenokia,
		"palm":              DeviceComplianceSettingStatePlatformTypepalm,
		"surfacehub":        DeviceComplianceSettingStatePlatformTypesurfaceHub,
		"unix":              DeviceComplianceSettingStatePlatformTypeunix,
		"unknown":           DeviceComplianceSettingStatePlatformTypeunknown,
		"wince":             DeviceComplianceSettingStatePlatformTypewinCE,
		"winembedded":       DeviceComplianceSettingStatePlatformTypewinEmbedded,
		"winmo6":            DeviceComplianceSettingStatePlatformTypewinMO6,
		"windows10x":        DeviceComplianceSettingStatePlatformTypewindows10x,
		"windowsphone":      DeviceComplianceSettingStatePlatformTypewindowsPhone,
		"windowsrt":         DeviceComplianceSettingStatePlatformTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceSettingStatePlatformType(input)
	return &out, nil
}
