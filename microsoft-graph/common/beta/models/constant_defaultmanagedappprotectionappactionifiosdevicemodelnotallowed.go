package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed string

const (
	DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedblock DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "Block"
	DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwarn  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "Warn"
	DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwipe  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedblock),
		string(DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwarn),
		string(DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input string) (*DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed{
		"block": DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedblock,
		"warn":  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input)
	return &out, nil
}
