package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationSetAccountManager string

const (
	SharedPCConfigurationSetAccountManagerdisabled      SharedPCConfigurationSetAccountManager = "Disabled"
	SharedPCConfigurationSetAccountManagerenabled       SharedPCConfigurationSetAccountManager = "Enabled"
	SharedPCConfigurationSetAccountManagernotConfigured SharedPCConfigurationSetAccountManager = "NotConfigured"
)

func PossibleValuesForSharedPCConfigurationSetAccountManager() []string {
	return []string{
		string(SharedPCConfigurationSetAccountManagerdisabled),
		string(SharedPCConfigurationSetAccountManagerenabled),
		string(SharedPCConfigurationSetAccountManagernotConfigured),
	}
}

func parseSharedPCConfigurationSetAccountManager(input string) (*SharedPCConfigurationSetAccountManager, error) {
	vals := map[string]SharedPCConfigurationSetAccountManager{
		"disabled":      SharedPCConfigurationSetAccountManagerdisabled,
		"enabled":       SharedPCConfigurationSetAccountManagerenabled,
		"notconfigured": SharedPCConfigurationSetAccountManagernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationSetAccountManager(input)
	return &out, nil
}
