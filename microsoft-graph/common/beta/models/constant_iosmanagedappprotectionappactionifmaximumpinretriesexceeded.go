package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Block"
	IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Warn"
	IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "Wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock),
		string(IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn),
		string(IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe),
	}
}

func parseIosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededblock,
		"warn":  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededwarn,
		"wipe":  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceededwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
