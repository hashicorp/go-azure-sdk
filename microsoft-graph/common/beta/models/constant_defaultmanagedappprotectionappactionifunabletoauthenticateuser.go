package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserblock DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Block"
	DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Warn"
	DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserblock),
		string(DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn),
		string(DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe),
	}
}

func parseDefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserblock,
		"warn":  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn,
		"wipe":  DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
