package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConfigurationTaskApplicablePlatform string

const (
	SecurityConfigurationTaskApplicablePlatformmacOS                     SecurityConfigurationTaskApplicablePlatform = "MacOS"
	SecurityConfigurationTaskApplicablePlatformunknown                   SecurityConfigurationTaskApplicablePlatform = "Unknown"
	SecurityConfigurationTaskApplicablePlatformwindows10AndLater         SecurityConfigurationTaskApplicablePlatform = "Windows10AndLater"
	SecurityConfigurationTaskApplicablePlatformwindows10AndWindowsServer SecurityConfigurationTaskApplicablePlatform = "Windows10AndWindowsServer"
)

func PossibleValuesForSecurityConfigurationTaskApplicablePlatform() []string {
	return []string{
		string(SecurityConfigurationTaskApplicablePlatformmacOS),
		string(SecurityConfigurationTaskApplicablePlatformunknown),
		string(SecurityConfigurationTaskApplicablePlatformwindows10AndLater),
		string(SecurityConfigurationTaskApplicablePlatformwindows10AndWindowsServer),
	}
}

func parseSecurityConfigurationTaskApplicablePlatform(input string) (*SecurityConfigurationTaskApplicablePlatform, error) {
	vals := map[string]SecurityConfigurationTaskApplicablePlatform{
		"macos":                     SecurityConfigurationTaskApplicablePlatformmacOS,
		"unknown":                   SecurityConfigurationTaskApplicablePlatformunknown,
		"windows10andlater":         SecurityConfigurationTaskApplicablePlatformwindows10AndLater,
		"windows10andwindowsserver": SecurityConfigurationTaskApplicablePlatformwindows10AndWindowsServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityConfigurationTaskApplicablePlatform(input)
	return &out, nil
}
