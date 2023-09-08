package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationCountRegistrationStatus string

const (
	UserRegistrationCountRegistrationStatuscapable       UserRegistrationCountRegistrationStatus = "Capable"
	UserRegistrationCountRegistrationStatusenabled       UserRegistrationCountRegistrationStatus = "Enabled"
	UserRegistrationCountRegistrationStatusmfaRegistered UserRegistrationCountRegistrationStatus = "MfaRegistered"
	UserRegistrationCountRegistrationStatusregistered    UserRegistrationCountRegistrationStatus = "Registered"
)

func PossibleValuesForUserRegistrationCountRegistrationStatus() []string {
	return []string{
		string(UserRegistrationCountRegistrationStatuscapable),
		string(UserRegistrationCountRegistrationStatusenabled),
		string(UserRegistrationCountRegistrationStatusmfaRegistered),
		string(UserRegistrationCountRegistrationStatusregistered),
	}
}

func parseUserRegistrationCountRegistrationStatus(input string) (*UserRegistrationCountRegistrationStatus, error) {
	vals := map[string]UserRegistrationCountRegistrationStatus{
		"capable":       UserRegistrationCountRegistrationStatuscapable,
		"enabled":       UserRegistrationCountRegistrationStatusenabled,
		"mfaregistered": UserRegistrationCountRegistrationStatusmfaRegistered,
		"registered":    UserRegistrationCountRegistrationStatusregistered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationCountRegistrationStatus(input)
	return &out, nil
}
