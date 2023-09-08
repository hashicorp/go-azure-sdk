package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfUnableToAuthenticateUser string

const (
	IosManagedAppProtectionAppActionIfUnableToAuthenticateUserblock IosManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Block"
	IosManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn  IosManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Warn"
	IosManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe  IosManagedAppProtectionAppActionIfUnableToAuthenticateUser = "Wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfUnableToAuthenticateUser() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfUnableToAuthenticateUserblock),
		string(IosManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn),
		string(IosManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe),
	}
}

func parseIosManagedAppProtectionAppActionIfUnableToAuthenticateUser(input string) (*IosManagedAppProtectionAppActionIfUnableToAuthenticateUser, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfUnableToAuthenticateUser{
		"block": IosManagedAppProtectionAppActionIfUnableToAuthenticateUserblock,
		"warn":  IosManagedAppProtectionAppActionIfUnableToAuthenticateUserwarn,
		"wipe":  IosManagedAppProtectionAppActionIfUnableToAuthenticateUserwipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfUnableToAuthenticateUser(input)
	return &out, nil
}
