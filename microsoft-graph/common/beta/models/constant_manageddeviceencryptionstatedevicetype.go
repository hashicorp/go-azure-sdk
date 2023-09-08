package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateDeviceType string

const (
	ManagedDeviceEncryptionStateDeviceTypeandroid           ManagedDeviceEncryptionStateDeviceType = "Android"
	ManagedDeviceEncryptionStateDeviceTypeandroidEnterprise ManagedDeviceEncryptionStateDeviceType = "AndroidEnterprise"
	ManagedDeviceEncryptionStateDeviceTypeandroidForWork    ManagedDeviceEncryptionStateDeviceType = "AndroidForWork"
	ManagedDeviceEncryptionStateDeviceTypeblackberry        ManagedDeviceEncryptionStateDeviceType = "Blackberry"
	ManagedDeviceEncryptionStateDeviceTypedesktop           ManagedDeviceEncryptionStateDeviceType = "Desktop"
	ManagedDeviceEncryptionStateDeviceTypeholoLens          ManagedDeviceEncryptionStateDeviceType = "HoloLens"
	ManagedDeviceEncryptionStateDeviceTypeiPad              ManagedDeviceEncryptionStateDeviceType = "IPad"
	ManagedDeviceEncryptionStateDeviceTypeiPhone            ManagedDeviceEncryptionStateDeviceType = "IPhone"
	ManagedDeviceEncryptionStateDeviceTypeiPod              ManagedDeviceEncryptionStateDeviceType = "IPod"
	ManagedDeviceEncryptionStateDeviceTypeiSocConsumer      ManagedDeviceEncryptionStateDeviceType = "ISocConsumer"
	ManagedDeviceEncryptionStateDeviceTypemac               ManagedDeviceEncryptionStateDeviceType = "Mac"
	ManagedDeviceEncryptionStateDeviceTypemacMDM            ManagedDeviceEncryptionStateDeviceType = "MacMDM"
	ManagedDeviceEncryptionStateDeviceTypenokia             ManagedDeviceEncryptionStateDeviceType = "Nokia"
	ManagedDeviceEncryptionStateDeviceTypepalm              ManagedDeviceEncryptionStateDeviceType = "Palm"
	ManagedDeviceEncryptionStateDeviceTypesurfaceHub        ManagedDeviceEncryptionStateDeviceType = "SurfaceHub"
	ManagedDeviceEncryptionStateDeviceTypeunix              ManagedDeviceEncryptionStateDeviceType = "Unix"
	ManagedDeviceEncryptionStateDeviceTypeunknown           ManagedDeviceEncryptionStateDeviceType = "Unknown"
	ManagedDeviceEncryptionStateDeviceTypewinCE             ManagedDeviceEncryptionStateDeviceType = "WinCE"
	ManagedDeviceEncryptionStateDeviceTypewinEmbedded       ManagedDeviceEncryptionStateDeviceType = "WinEmbedded"
	ManagedDeviceEncryptionStateDeviceTypewinMO6            ManagedDeviceEncryptionStateDeviceType = "WinMO6"
	ManagedDeviceEncryptionStateDeviceTypewindowsPhone      ManagedDeviceEncryptionStateDeviceType = "WindowsPhone"
	ManagedDeviceEncryptionStateDeviceTypewindowsRT         ManagedDeviceEncryptionStateDeviceType = "WindowsRT"
)

func PossibleValuesForManagedDeviceEncryptionStateDeviceType() []string {
	return []string{
		string(ManagedDeviceEncryptionStateDeviceTypeandroid),
		string(ManagedDeviceEncryptionStateDeviceTypeandroidEnterprise),
		string(ManagedDeviceEncryptionStateDeviceTypeandroidForWork),
		string(ManagedDeviceEncryptionStateDeviceTypeblackberry),
		string(ManagedDeviceEncryptionStateDeviceTypedesktop),
		string(ManagedDeviceEncryptionStateDeviceTypeholoLens),
		string(ManagedDeviceEncryptionStateDeviceTypeiPad),
		string(ManagedDeviceEncryptionStateDeviceTypeiPhone),
		string(ManagedDeviceEncryptionStateDeviceTypeiPod),
		string(ManagedDeviceEncryptionStateDeviceTypeiSocConsumer),
		string(ManagedDeviceEncryptionStateDeviceTypemac),
		string(ManagedDeviceEncryptionStateDeviceTypemacMDM),
		string(ManagedDeviceEncryptionStateDeviceTypenokia),
		string(ManagedDeviceEncryptionStateDeviceTypepalm),
		string(ManagedDeviceEncryptionStateDeviceTypesurfaceHub),
		string(ManagedDeviceEncryptionStateDeviceTypeunix),
		string(ManagedDeviceEncryptionStateDeviceTypeunknown),
		string(ManagedDeviceEncryptionStateDeviceTypewinCE),
		string(ManagedDeviceEncryptionStateDeviceTypewinEmbedded),
		string(ManagedDeviceEncryptionStateDeviceTypewinMO6),
		string(ManagedDeviceEncryptionStateDeviceTypewindowsPhone),
		string(ManagedDeviceEncryptionStateDeviceTypewindowsRT),
	}
}

func parseManagedDeviceEncryptionStateDeviceType(input string) (*ManagedDeviceEncryptionStateDeviceType, error) {
	vals := map[string]ManagedDeviceEncryptionStateDeviceType{
		"android":           ManagedDeviceEncryptionStateDeviceTypeandroid,
		"androidenterprise": ManagedDeviceEncryptionStateDeviceTypeandroidEnterprise,
		"androidforwork":    ManagedDeviceEncryptionStateDeviceTypeandroidForWork,
		"blackberry":        ManagedDeviceEncryptionStateDeviceTypeblackberry,
		"desktop":           ManagedDeviceEncryptionStateDeviceTypedesktop,
		"hololens":          ManagedDeviceEncryptionStateDeviceTypeholoLens,
		"ipad":              ManagedDeviceEncryptionStateDeviceTypeiPad,
		"iphone":            ManagedDeviceEncryptionStateDeviceTypeiPhone,
		"ipod":              ManagedDeviceEncryptionStateDeviceTypeiPod,
		"isocconsumer":      ManagedDeviceEncryptionStateDeviceTypeiSocConsumer,
		"mac":               ManagedDeviceEncryptionStateDeviceTypemac,
		"macmdm":            ManagedDeviceEncryptionStateDeviceTypemacMDM,
		"nokia":             ManagedDeviceEncryptionStateDeviceTypenokia,
		"palm":              ManagedDeviceEncryptionStateDeviceTypepalm,
		"surfacehub":        ManagedDeviceEncryptionStateDeviceTypesurfaceHub,
		"unix":              ManagedDeviceEncryptionStateDeviceTypeunix,
		"unknown":           ManagedDeviceEncryptionStateDeviceTypeunknown,
		"wince":             ManagedDeviceEncryptionStateDeviceTypewinCE,
		"winembedded":       ManagedDeviceEncryptionStateDeviceTypewinEmbedded,
		"winmo6":            ManagedDeviceEncryptionStateDeviceTypewinMO6,
		"windowsphone":      ManagedDeviceEncryptionStateDeviceTypewindowsPhone,
		"windowsrt":         ManagedDeviceEncryptionStateDeviceTypewindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateDeviceType(input)
	return &out, nil
}
