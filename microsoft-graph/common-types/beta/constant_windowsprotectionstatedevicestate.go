package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsProtectionStateDeviceState string

const (
	WindowsProtectionStateDeviceState_Clean              WindowsProtectionStateDeviceState = "clean"
	WindowsProtectionStateDeviceState_Critical           WindowsProtectionStateDeviceState = "critical"
	WindowsProtectionStateDeviceState_FullScanPending    WindowsProtectionStateDeviceState = "fullScanPending"
	WindowsProtectionStateDeviceState_ManualStepsPending WindowsProtectionStateDeviceState = "manualStepsPending"
	WindowsProtectionStateDeviceState_OfflineScanPending WindowsProtectionStateDeviceState = "offlineScanPending"
	WindowsProtectionStateDeviceState_RebootPending      WindowsProtectionStateDeviceState = "rebootPending"
)

func PossibleValuesForWindowsProtectionStateDeviceState() []string {
	return []string{
		string(WindowsProtectionStateDeviceState_Clean),
		string(WindowsProtectionStateDeviceState_Critical),
		string(WindowsProtectionStateDeviceState_FullScanPending),
		string(WindowsProtectionStateDeviceState_ManualStepsPending),
		string(WindowsProtectionStateDeviceState_OfflineScanPending),
		string(WindowsProtectionStateDeviceState_RebootPending),
	}
}

func (s *WindowsProtectionStateDeviceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsProtectionStateDeviceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsProtectionStateDeviceState(input string) (*WindowsProtectionStateDeviceState, error) {
	vals := map[string]WindowsProtectionStateDeviceState{
		"clean":              WindowsProtectionStateDeviceState_Clean,
		"critical":           WindowsProtectionStateDeviceState_Critical,
		"fullscanpending":    WindowsProtectionStateDeviceState_FullScanPending,
		"manualstepspending": WindowsProtectionStateDeviceState_ManualStepsPending,
		"offlinescanpending": WindowsProtectionStateDeviceState_OfflineScanPending,
		"rebootpending":      WindowsProtectionStateDeviceState_RebootPending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsProtectionStateDeviceState(input)
	return &out, nil
}
