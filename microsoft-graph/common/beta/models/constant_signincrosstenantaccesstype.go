package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInCrossTenantAccessType string

const (
	SignInCrossTenantAccessTypeb2bCollaboration SignInCrossTenantAccessType = "B2bCollaboration"
	SignInCrossTenantAccessTypeb2bDirectConnect SignInCrossTenantAccessType = "B2bDirectConnect"
	SignInCrossTenantAccessTypemicrosoftSupport SignInCrossTenantAccessType = "MicrosoftSupport"
	SignInCrossTenantAccessTypenone             SignInCrossTenantAccessType = "None"
	SignInCrossTenantAccessTypeserviceProvider  SignInCrossTenantAccessType = "ServiceProvider"
)

func PossibleValuesForSignInCrossTenantAccessType() []string {
	return []string{
		string(SignInCrossTenantAccessTypeb2bCollaboration),
		string(SignInCrossTenantAccessTypeb2bDirectConnect),
		string(SignInCrossTenantAccessTypemicrosoftSupport),
		string(SignInCrossTenantAccessTypenone),
		string(SignInCrossTenantAccessTypeserviceProvider),
	}
}

func parseSignInCrossTenantAccessType(input string) (*SignInCrossTenantAccessType, error) {
	vals := map[string]SignInCrossTenantAccessType{
		"b2bcollaboration": SignInCrossTenantAccessTypeb2bCollaboration,
		"b2bdirectconnect": SignInCrossTenantAccessTypeb2bDirectConnect,
		"microsoftsupport": SignInCrossTenantAccessTypemicrosoftSupport,
		"none":             SignInCrossTenantAccessTypenone,
		"serviceprovider":  SignInCrossTenantAccessTypeserviceProvider,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInCrossTenantAccessType(input)
	return &out, nil
}
