package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChromeOSOnboardingSettingsOnboardingStatus string

const (
	ChromeOSOnboardingSettingsOnboardingStatus_Failed      ChromeOSOnboardingSettingsOnboardingStatus = "failed"
	ChromeOSOnboardingSettingsOnboardingStatus_Inprogress  ChromeOSOnboardingSettingsOnboardingStatus = "inprogress"
	ChromeOSOnboardingSettingsOnboardingStatus_Offboarding ChromeOSOnboardingSettingsOnboardingStatus = "offboarding"
	ChromeOSOnboardingSettingsOnboardingStatus_Onboarded   ChromeOSOnboardingSettingsOnboardingStatus = "onboarded"
	ChromeOSOnboardingSettingsOnboardingStatus_Unknown     ChromeOSOnboardingSettingsOnboardingStatus = "unknown"
)

func PossibleValuesForChromeOSOnboardingSettingsOnboardingStatus() []string {
	return []string{
		string(ChromeOSOnboardingSettingsOnboardingStatus_Failed),
		string(ChromeOSOnboardingSettingsOnboardingStatus_Inprogress),
		string(ChromeOSOnboardingSettingsOnboardingStatus_Offboarding),
		string(ChromeOSOnboardingSettingsOnboardingStatus_Onboarded),
		string(ChromeOSOnboardingSettingsOnboardingStatus_Unknown),
	}
}

func (s *ChromeOSOnboardingSettingsOnboardingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChromeOSOnboardingSettingsOnboardingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChromeOSOnboardingSettingsOnboardingStatus(input string) (*ChromeOSOnboardingSettingsOnboardingStatus, error) {
	vals := map[string]ChromeOSOnboardingSettingsOnboardingStatus{
		"failed":      ChromeOSOnboardingSettingsOnboardingStatus_Failed,
		"inprogress":  ChromeOSOnboardingSettingsOnboardingStatus_Inprogress,
		"offboarding": ChromeOSOnboardingSettingsOnboardingStatus_Offboarding,
		"onboarded":   ChromeOSOnboardingSettingsOnboardingStatus_Onboarded,
		"unknown":     ChromeOSOnboardingSettingsOnboardingStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChromeOSOnboardingSettingsOnboardingStatus(input)
	return &out, nil
}
