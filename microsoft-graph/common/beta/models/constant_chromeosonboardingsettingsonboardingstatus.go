package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChromeOSOnboardingSettingsOnboardingStatus string

const (
	ChromeOSOnboardingSettingsOnboardingStatusfailed      ChromeOSOnboardingSettingsOnboardingStatus = "Failed"
	ChromeOSOnboardingSettingsOnboardingStatusinprogress  ChromeOSOnboardingSettingsOnboardingStatus = "Inprogress"
	ChromeOSOnboardingSettingsOnboardingStatusoffboarding ChromeOSOnboardingSettingsOnboardingStatus = "Offboarding"
	ChromeOSOnboardingSettingsOnboardingStatusonboarded   ChromeOSOnboardingSettingsOnboardingStatus = "Onboarded"
	ChromeOSOnboardingSettingsOnboardingStatusunknown     ChromeOSOnboardingSettingsOnboardingStatus = "Unknown"
)

func PossibleValuesForChromeOSOnboardingSettingsOnboardingStatus() []string {
	return []string{
		string(ChromeOSOnboardingSettingsOnboardingStatusfailed),
		string(ChromeOSOnboardingSettingsOnboardingStatusinprogress),
		string(ChromeOSOnboardingSettingsOnboardingStatusoffboarding),
		string(ChromeOSOnboardingSettingsOnboardingStatusonboarded),
		string(ChromeOSOnboardingSettingsOnboardingStatusunknown),
	}
}

func parseChromeOSOnboardingSettingsOnboardingStatus(input string) (*ChromeOSOnboardingSettingsOnboardingStatus, error) {
	vals := map[string]ChromeOSOnboardingSettingsOnboardingStatus{
		"failed":      ChromeOSOnboardingSettingsOnboardingStatusfailed,
		"inprogress":  ChromeOSOnboardingSettingsOnboardingStatusinprogress,
		"offboarding": ChromeOSOnboardingSettingsOnboardingStatusoffboarding,
		"onboarded":   ChromeOSOnboardingSettingsOnboardingStatusonboarded,
		"unknown":     ChromeOSOnboardingSettingsOnboardingStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChromeOSOnboardingSettingsOnboardingStatus(input)
	return &out, nil
}
