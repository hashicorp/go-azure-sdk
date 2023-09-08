package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed string

const (
	IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedblock IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "Block"
	IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwarn  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "Warn"
	IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwipe  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "Wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedblock),
		string(IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwarn),
		string(IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwipe),
	}
}

func parseIosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input string) (*IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed{
		"block": IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedblock,
		"warn":  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwarn,
		"wipe":  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowedwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input)
	return &out, nil
}
