package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateDeviceType string

const (
	ManagedDeviceEncryptionStateDeviceType_Android           ManagedDeviceEncryptionStateDeviceType = "android"
	ManagedDeviceEncryptionStateDeviceType_AndroidEnterprise ManagedDeviceEncryptionStateDeviceType = "androidEnterprise"
	ManagedDeviceEncryptionStateDeviceType_AndroidForWork    ManagedDeviceEncryptionStateDeviceType = "androidForWork"
	ManagedDeviceEncryptionStateDeviceType_Blackberry        ManagedDeviceEncryptionStateDeviceType = "blackberry"
	ManagedDeviceEncryptionStateDeviceType_Desktop           ManagedDeviceEncryptionStateDeviceType = "desktop"
	ManagedDeviceEncryptionStateDeviceType_HoloLens          ManagedDeviceEncryptionStateDeviceType = "holoLens"
	ManagedDeviceEncryptionStateDeviceType_IPad              ManagedDeviceEncryptionStateDeviceType = "iPad"
	ManagedDeviceEncryptionStateDeviceType_IPhone            ManagedDeviceEncryptionStateDeviceType = "iPhone"
	ManagedDeviceEncryptionStateDeviceType_IPod              ManagedDeviceEncryptionStateDeviceType = "iPod"
	ManagedDeviceEncryptionStateDeviceType_ISocConsumer      ManagedDeviceEncryptionStateDeviceType = "iSocConsumer"
	ManagedDeviceEncryptionStateDeviceType_Mac               ManagedDeviceEncryptionStateDeviceType = "mac"
	ManagedDeviceEncryptionStateDeviceType_MacMDM            ManagedDeviceEncryptionStateDeviceType = "macMDM"
	ManagedDeviceEncryptionStateDeviceType_Nokia             ManagedDeviceEncryptionStateDeviceType = "nokia"
	ManagedDeviceEncryptionStateDeviceType_Palm              ManagedDeviceEncryptionStateDeviceType = "palm"
	ManagedDeviceEncryptionStateDeviceType_SurfaceHub        ManagedDeviceEncryptionStateDeviceType = "surfaceHub"
	ManagedDeviceEncryptionStateDeviceType_Unix              ManagedDeviceEncryptionStateDeviceType = "unix"
	ManagedDeviceEncryptionStateDeviceType_Unknown           ManagedDeviceEncryptionStateDeviceType = "unknown"
	ManagedDeviceEncryptionStateDeviceType_WinCE             ManagedDeviceEncryptionStateDeviceType = "winCE"
	ManagedDeviceEncryptionStateDeviceType_WinEmbedded       ManagedDeviceEncryptionStateDeviceType = "winEmbedded"
	ManagedDeviceEncryptionStateDeviceType_WinMO6            ManagedDeviceEncryptionStateDeviceType = "winMO6"
	ManagedDeviceEncryptionStateDeviceType_WindowsPhone      ManagedDeviceEncryptionStateDeviceType = "windowsPhone"
	ManagedDeviceEncryptionStateDeviceType_WindowsRT         ManagedDeviceEncryptionStateDeviceType = "windowsRT"
)

func PossibleValuesForManagedDeviceEncryptionStateDeviceType() []string {
	return []string{
		string(ManagedDeviceEncryptionStateDeviceType_Android),
		string(ManagedDeviceEncryptionStateDeviceType_AndroidEnterprise),
		string(ManagedDeviceEncryptionStateDeviceType_AndroidForWork),
		string(ManagedDeviceEncryptionStateDeviceType_Blackberry),
		string(ManagedDeviceEncryptionStateDeviceType_Desktop),
		string(ManagedDeviceEncryptionStateDeviceType_HoloLens),
		string(ManagedDeviceEncryptionStateDeviceType_IPad),
		string(ManagedDeviceEncryptionStateDeviceType_IPhone),
		string(ManagedDeviceEncryptionStateDeviceType_IPod),
		string(ManagedDeviceEncryptionStateDeviceType_ISocConsumer),
		string(ManagedDeviceEncryptionStateDeviceType_Mac),
		string(ManagedDeviceEncryptionStateDeviceType_MacMDM),
		string(ManagedDeviceEncryptionStateDeviceType_Nokia),
		string(ManagedDeviceEncryptionStateDeviceType_Palm),
		string(ManagedDeviceEncryptionStateDeviceType_SurfaceHub),
		string(ManagedDeviceEncryptionStateDeviceType_Unix),
		string(ManagedDeviceEncryptionStateDeviceType_Unknown),
		string(ManagedDeviceEncryptionStateDeviceType_WinCE),
		string(ManagedDeviceEncryptionStateDeviceType_WinEmbedded),
		string(ManagedDeviceEncryptionStateDeviceType_WinMO6),
		string(ManagedDeviceEncryptionStateDeviceType_WindowsPhone),
		string(ManagedDeviceEncryptionStateDeviceType_WindowsRT),
	}
}

func (s *ManagedDeviceEncryptionStateDeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceEncryptionStateDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceEncryptionStateDeviceType(input string) (*ManagedDeviceEncryptionStateDeviceType, error) {
	vals := map[string]ManagedDeviceEncryptionStateDeviceType{
		"android":           ManagedDeviceEncryptionStateDeviceType_Android,
		"androidenterprise": ManagedDeviceEncryptionStateDeviceType_AndroidEnterprise,
		"androidforwork":    ManagedDeviceEncryptionStateDeviceType_AndroidForWork,
		"blackberry":        ManagedDeviceEncryptionStateDeviceType_Blackberry,
		"desktop":           ManagedDeviceEncryptionStateDeviceType_Desktop,
		"hololens":          ManagedDeviceEncryptionStateDeviceType_HoloLens,
		"ipad":              ManagedDeviceEncryptionStateDeviceType_IPad,
		"iphone":            ManagedDeviceEncryptionStateDeviceType_IPhone,
		"ipod":              ManagedDeviceEncryptionStateDeviceType_IPod,
		"isocconsumer":      ManagedDeviceEncryptionStateDeviceType_ISocConsumer,
		"mac":               ManagedDeviceEncryptionStateDeviceType_Mac,
		"macmdm":            ManagedDeviceEncryptionStateDeviceType_MacMDM,
		"nokia":             ManagedDeviceEncryptionStateDeviceType_Nokia,
		"palm":              ManagedDeviceEncryptionStateDeviceType_Palm,
		"surfacehub":        ManagedDeviceEncryptionStateDeviceType_SurfaceHub,
		"unix":              ManagedDeviceEncryptionStateDeviceType_Unix,
		"unknown":           ManagedDeviceEncryptionStateDeviceType_Unknown,
		"wince":             ManagedDeviceEncryptionStateDeviceType_WinCE,
		"winembedded":       ManagedDeviceEncryptionStateDeviceType_WinEmbedded,
		"winmo6":            ManagedDeviceEncryptionStateDeviceType_WinMO6,
		"windowsphone":      ManagedDeviceEncryptionStateDeviceType_WindowsPhone,
		"windowsrt":         ManagedDeviceEncryptionStateDeviceType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateDeviceType(input)
	return &out, nil
}
