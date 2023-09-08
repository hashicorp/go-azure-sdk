package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPCConfigurationAllowedAccounts string

const (
	SharedPCConfigurationAllowedAccountsdomain        SharedPCConfigurationAllowedAccounts = "Domain"
	SharedPCConfigurationAllowedAccountsguest         SharedPCConfigurationAllowedAccounts = "Guest"
	SharedPCConfigurationAllowedAccountsnotConfigured SharedPCConfigurationAllowedAccounts = "NotConfigured"
)

func PossibleValuesForSharedPCConfigurationAllowedAccounts() []string {
	return []string{
		string(SharedPCConfigurationAllowedAccountsdomain),
		string(SharedPCConfigurationAllowedAccountsguest),
		string(SharedPCConfigurationAllowedAccountsnotConfigured),
	}
}

func parseSharedPCConfigurationAllowedAccounts(input string) (*SharedPCConfigurationAllowedAccounts, error) {
	vals := map[string]SharedPCConfigurationAllowedAccounts{
		"domain":        SharedPCConfigurationAllowedAccountsdomain,
		"guest":         SharedPCConfigurationAllowedAccountsguest,
		"notconfigured": SharedPCConfigurationAllowedAccountsnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharedPCConfigurationAllowedAccounts(input)
	return &out, nil
}
