package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSignInOnResume string

const (
	SharedPCConfigurationSignInOnResumedisabled      SharedPCConfigurationSignInOnResume = "Disabled"
	SharedPCConfigurationSignInOnResumeenabled       SharedPCConfigurationSignInOnResume = "Enabled"
	SharedPCConfigurationSignInOnResumenotConfigured SharedPCConfigurationSignInOnResume = "NotConfigured"
)

func PossibleValuesForSharedPCConfigurationSignInOnResume() []string {
	return []string{
		string(SharedPCConfigurationSignInOnResumedisabled),
		string(SharedPCConfigurationSignInOnResumeenabled),
		string(SharedPCConfigurationSignInOnResumenotConfigured),
	}
}

func parseSharedPCConfigurationSignInOnResume(input string) (*SharedPCConfigurationSignInOnResume, error) {
	vals := map[string]SharedPCConfigurationSignInOnResume{
		"disabled":      SharedPCConfigurationSignInOnResumedisabled,
		"enabled":       SharedPCConfigurationSignInOnResumeenabled,
		"notconfigured": SharedPCConfigurationSignInOnResumenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSignInOnResume(input)
	return &out, nil
}
