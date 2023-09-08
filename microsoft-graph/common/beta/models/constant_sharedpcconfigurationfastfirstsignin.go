package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationFastFirstSignIn string

const (
	SharedPCConfigurationFastFirstSignIndisabled      SharedPCConfigurationFastFirstSignIn = "Disabled"
	SharedPCConfigurationFastFirstSignInenabled       SharedPCConfigurationFastFirstSignIn = "Enabled"
	SharedPCConfigurationFastFirstSignInnotConfigured SharedPCConfigurationFastFirstSignIn = "NotConfigured"
)

func PossibleValuesForSharedPCConfigurationFastFirstSignIn() []string {
	return []string{
		string(SharedPCConfigurationFastFirstSignIndisabled),
		string(SharedPCConfigurationFastFirstSignInenabled),
		string(SharedPCConfigurationFastFirstSignInnotConfigured),
	}
}

func parseSharedPCConfigurationFastFirstSignIn(input string) (*SharedPCConfigurationFastFirstSignIn, error) {
	vals := map[string]SharedPCConfigurationFastFirstSignIn{
		"disabled":      SharedPCConfigurationFastFirstSignIndisabled,
		"enabled":       SharedPCConfigurationFastFirstSignInenabled,
		"notconfigured": SharedPCConfigurationFastFirstSignInnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationFastFirstSignIn(input)
	return &out, nil
}
