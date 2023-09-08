package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserblock TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Block"
	TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Warn"
	TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Wipe"
)

func PossibleValuesForTargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserblock),
		string(TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn),
		string(TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe),
	}
}

func parseTargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserblock,
		"warn":  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn,
		"wipe":  TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
