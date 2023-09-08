package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	ManagedAppProtectionAppActionIfUnableToAuthenticateUserblock ManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Block"
	ManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn  ManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Warn"
	ManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe  ManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Wipe"
)

func PossibleValuesForManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(ManagedAppProtectionAppActionIfUnableToAuthenticateUserblock),
		string(ManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn),
		string(ManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe),
	}
}

func parseManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*ManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]ManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": ManagedAppProtectionAppActionIfUnableToAuthenticateUserblock,
		"warn":  ManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn,
		"wipe":  ManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
