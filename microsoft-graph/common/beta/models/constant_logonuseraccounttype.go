package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUserAccountType string

const (
	LogonUserAccountTypeadministrator LogonUserAccountType = "Administrator"
	LogonUserAccountTypepower         LogonUserAccountType = "Power"
	LogonUserAccountTypestandard      LogonUserAccountType = "Standard"
	LogonUserAccountTypeunknown       LogonUserAccountType = "Unknown"
)

func PossibleValuesForLogonUserAccountType() []string {
	return []string{
		string(LogonUserAccountTypeadministrator),
		string(LogonUserAccountTypepower),
		string(LogonUserAccountTypestandard),
		string(LogonUserAccountTypeunknown),
	}
}

func parseLogonUserAccountType(input string) (*LogonUserAccountType, error) {
	vals := map[string]LogonUserAccountType{
		"administrator": LogonUserAccountTypeadministrator,
		"power":         LogonUserAccountTypepower,
		"standard":      LogonUserAccountTypestandard,
		"unknown":       LogonUserAccountTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogonUserAccountType(input)
	return &out, nil
}
