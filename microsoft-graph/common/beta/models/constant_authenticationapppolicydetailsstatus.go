package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsStatus string

const (
	AuthenticationAppPolicyDetailsStatusappContextNotShown                AuthenticationAppPolicyDetailsStatus = "AppContextNotShown"
	AuthenticationAppPolicyDetailsStatusappContextOutOfDate               AuthenticationAppPolicyDetailsStatus = "AppContextOutOfDate"
	AuthenticationAppPolicyDetailsStatusappContextShown                   AuthenticationAppPolicyDetailsStatus = "AppContextShown"
	AuthenticationAppPolicyDetailsStatusappLockDisabled                   AuthenticationAppPolicyDetailsStatus = "AppLockDisabled"
	AuthenticationAppPolicyDetailsStatusappLockEnabled                    AuthenticationAppPolicyDetailsStatus = "AppLockEnabled"
	AuthenticationAppPolicyDetailsStatusappLockOutOfDate                  AuthenticationAppPolicyDetailsStatus = "AppLockOutOfDate"
	AuthenticationAppPolicyDetailsStatuslocationContextNotShown           AuthenticationAppPolicyDetailsStatus = "LocationContextNotShown"
	AuthenticationAppPolicyDetailsStatuslocationContextOutOfDate          AuthenticationAppPolicyDetailsStatus = "LocationContextOutOfDate"
	AuthenticationAppPolicyDetailsStatuslocationContextShown              AuthenticationAppPolicyDetailsStatus = "LocationContextShown"
	AuthenticationAppPolicyDetailsStatusnumberMatchCorrectNumberEntered   AuthenticationAppPolicyDetailsStatus = "NumberMatchCorrectNumberEntered"
	AuthenticationAppPolicyDetailsStatusnumberMatchDeny                   AuthenticationAppPolicyDetailsStatus = "NumberMatchDeny"
	AuthenticationAppPolicyDetailsStatusnumberMatchIncorrectNumberEntered AuthenticationAppPolicyDetailsStatus = "NumberMatchIncorrectNumberEntered"
	AuthenticationAppPolicyDetailsStatusnumberMatchOutOfDate              AuthenticationAppPolicyDetailsStatus = "NumberMatchOutOfDate"
	AuthenticationAppPolicyDetailsStatustamperResistantHardwareNotUsed    AuthenticationAppPolicyDetailsStatus = "TamperResistantHardwareNotUsed"
	AuthenticationAppPolicyDetailsStatustamperResistantHardwareOutOfDate  AuthenticationAppPolicyDetailsStatus = "TamperResistantHardwareOutOfDate"
	AuthenticationAppPolicyDetailsStatustamperResistantHardwareUsed       AuthenticationAppPolicyDetailsStatus = "TamperResistantHardwareUsed"
	AuthenticationAppPolicyDetailsStatusunknown                           AuthenticationAppPolicyDetailsStatus = "Unknown"
)

func PossibleValuesForAuthenticationAppPolicyDetailsStatus() []string {
	return []string{
		string(AuthenticationAppPolicyDetailsStatusappContextNotShown),
		string(AuthenticationAppPolicyDetailsStatusappContextOutOfDate),
		string(AuthenticationAppPolicyDetailsStatusappContextShown),
		string(AuthenticationAppPolicyDetailsStatusappLockDisabled),
		string(AuthenticationAppPolicyDetailsStatusappLockEnabled),
		string(AuthenticationAppPolicyDetailsStatusappLockOutOfDate),
		string(AuthenticationAppPolicyDetailsStatuslocationContextNotShown),
		string(AuthenticationAppPolicyDetailsStatuslocationContextOutOfDate),
		string(AuthenticationAppPolicyDetailsStatuslocationContextShown),
		string(AuthenticationAppPolicyDetailsStatusnumberMatchCorrectNumberEntered),
		string(AuthenticationAppPolicyDetailsStatusnumberMatchDeny),
		string(AuthenticationAppPolicyDetailsStatusnumberMatchIncorrectNumberEntered),
		string(AuthenticationAppPolicyDetailsStatusnumberMatchOutOfDate),
		string(AuthenticationAppPolicyDetailsStatustamperResistantHardwareNotUsed),
		string(AuthenticationAppPolicyDetailsStatustamperResistantHardwareOutOfDate),
		string(AuthenticationAppPolicyDetailsStatustamperResistantHardwareUsed),
		string(AuthenticationAppPolicyDetailsStatusunknown),
	}
}

func parseAuthenticationAppPolicyDetailsStatus(input string) (*AuthenticationAppPolicyDetailsStatus, error) {
	vals := map[string]AuthenticationAppPolicyDetailsStatus{
		"appcontextnotshown":                AuthenticationAppPolicyDetailsStatusappContextNotShown,
		"appcontextoutofdate":               AuthenticationAppPolicyDetailsStatusappContextOutOfDate,
		"appcontextshown":                   AuthenticationAppPolicyDetailsStatusappContextShown,
		"applockdisabled":                   AuthenticationAppPolicyDetailsStatusappLockDisabled,
		"applockenabled":                    AuthenticationAppPolicyDetailsStatusappLockEnabled,
		"applockoutofdate":                  AuthenticationAppPolicyDetailsStatusappLockOutOfDate,
		"locationcontextnotshown":           AuthenticationAppPolicyDetailsStatuslocationContextNotShown,
		"locationcontextoutofdate":          AuthenticationAppPolicyDetailsStatuslocationContextOutOfDate,
		"locationcontextshown":              AuthenticationAppPolicyDetailsStatuslocationContextShown,
		"numbermatchcorrectnumberentered":   AuthenticationAppPolicyDetailsStatusnumberMatchCorrectNumberEntered,
		"numbermatchdeny":                   AuthenticationAppPolicyDetailsStatusnumberMatchDeny,
		"numbermatchincorrectnumberentered": AuthenticationAppPolicyDetailsStatusnumberMatchIncorrectNumberEntered,
		"numbermatchoutofdate":              AuthenticationAppPolicyDetailsStatusnumberMatchOutOfDate,
		"tamperresistanthardwarenotused":    AuthenticationAppPolicyDetailsStatustamperResistantHardwareNotUsed,
		"tamperresistanthardwareoutofdate":  AuthenticationAppPolicyDetailsStatustamperResistantHardwareOutOfDate,
		"tamperresistanthardwareused":       AuthenticationAppPolicyDetailsStatustamperResistantHardwareUsed,
		"unknown":                           AuthenticationAppPolicyDetailsStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyDetailsStatus(input)
	return &out, nil
}
