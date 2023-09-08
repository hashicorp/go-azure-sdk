package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserblock AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Block"
	AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Warn"
	AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserblock),
		string(AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn),
		string(AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe),
	}
}

func parseAndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserblock,
		"warn":  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn,
		"wipe":  AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
