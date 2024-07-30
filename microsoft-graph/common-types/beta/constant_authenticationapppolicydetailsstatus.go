package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsStatus string

const (
	AuthenticationAppPolicyDetailsStatus_AppContextNotShown                AuthenticationAppPolicyDetailsStatus = "appContextNotShown"
	AuthenticationAppPolicyDetailsStatus_AppContextOutOfDate               AuthenticationAppPolicyDetailsStatus = "appContextOutOfDate"
	AuthenticationAppPolicyDetailsStatus_AppContextShown                   AuthenticationAppPolicyDetailsStatus = "appContextShown"
	AuthenticationAppPolicyDetailsStatus_AppLockDisabled                   AuthenticationAppPolicyDetailsStatus = "appLockDisabled"
	AuthenticationAppPolicyDetailsStatus_AppLockEnabled                    AuthenticationAppPolicyDetailsStatus = "appLockEnabled"
	AuthenticationAppPolicyDetailsStatus_AppLockOutOfDate                  AuthenticationAppPolicyDetailsStatus = "appLockOutOfDate"
	AuthenticationAppPolicyDetailsStatus_LocationContextNotShown           AuthenticationAppPolicyDetailsStatus = "locationContextNotShown"
	AuthenticationAppPolicyDetailsStatus_LocationContextOutOfDate          AuthenticationAppPolicyDetailsStatus = "locationContextOutOfDate"
	AuthenticationAppPolicyDetailsStatus_LocationContextShown              AuthenticationAppPolicyDetailsStatus = "locationContextShown"
	AuthenticationAppPolicyDetailsStatus_NumberMatchCorrectNumberEntered   AuthenticationAppPolicyDetailsStatus = "numberMatchCorrectNumberEntered"
	AuthenticationAppPolicyDetailsStatus_NumberMatchDeny                   AuthenticationAppPolicyDetailsStatus = "numberMatchDeny"
	AuthenticationAppPolicyDetailsStatus_NumberMatchIncorrectNumberEntered AuthenticationAppPolicyDetailsStatus = "numberMatchIncorrectNumberEntered"
	AuthenticationAppPolicyDetailsStatus_NumberMatchOutOfDate              AuthenticationAppPolicyDetailsStatus = "numberMatchOutOfDate"
	AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareNotUsed    AuthenticationAppPolicyDetailsStatus = "tamperResistantHardwareNotUsed"
	AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareOutOfDate  AuthenticationAppPolicyDetailsStatus = "tamperResistantHardwareOutOfDate"
	AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareUsed       AuthenticationAppPolicyDetailsStatus = "tamperResistantHardwareUsed"
	AuthenticationAppPolicyDetailsStatus_Unknown                           AuthenticationAppPolicyDetailsStatus = "unknown"
)

func PossibleValuesForAuthenticationAppPolicyDetailsStatus() []string {
	return []string{
		string(AuthenticationAppPolicyDetailsStatus_AppContextNotShown),
		string(AuthenticationAppPolicyDetailsStatus_AppContextOutOfDate),
		string(AuthenticationAppPolicyDetailsStatus_AppContextShown),
		string(AuthenticationAppPolicyDetailsStatus_AppLockDisabled),
		string(AuthenticationAppPolicyDetailsStatus_AppLockEnabled),
		string(AuthenticationAppPolicyDetailsStatus_AppLockOutOfDate),
		string(AuthenticationAppPolicyDetailsStatus_LocationContextNotShown),
		string(AuthenticationAppPolicyDetailsStatus_LocationContextOutOfDate),
		string(AuthenticationAppPolicyDetailsStatus_LocationContextShown),
		string(AuthenticationAppPolicyDetailsStatus_NumberMatchCorrectNumberEntered),
		string(AuthenticationAppPolicyDetailsStatus_NumberMatchDeny),
		string(AuthenticationAppPolicyDetailsStatus_NumberMatchIncorrectNumberEntered),
		string(AuthenticationAppPolicyDetailsStatus_NumberMatchOutOfDate),
		string(AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareNotUsed),
		string(AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareOutOfDate),
		string(AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareUsed),
		string(AuthenticationAppPolicyDetailsStatus_Unknown),
	}
}

func (s *AuthenticationAppPolicyDetailsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAppPolicyDetailsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAppPolicyDetailsStatus(input string) (*AuthenticationAppPolicyDetailsStatus, error) {
	vals := map[string]AuthenticationAppPolicyDetailsStatus{
		"appcontextnotshown":                AuthenticationAppPolicyDetailsStatus_AppContextNotShown,
		"appcontextoutofdate":               AuthenticationAppPolicyDetailsStatus_AppContextOutOfDate,
		"appcontextshown":                   AuthenticationAppPolicyDetailsStatus_AppContextShown,
		"applockdisabled":                   AuthenticationAppPolicyDetailsStatus_AppLockDisabled,
		"applockenabled":                    AuthenticationAppPolicyDetailsStatus_AppLockEnabled,
		"applockoutofdate":                  AuthenticationAppPolicyDetailsStatus_AppLockOutOfDate,
		"locationcontextnotshown":           AuthenticationAppPolicyDetailsStatus_LocationContextNotShown,
		"locationcontextoutofdate":          AuthenticationAppPolicyDetailsStatus_LocationContextOutOfDate,
		"locationcontextshown":              AuthenticationAppPolicyDetailsStatus_LocationContextShown,
		"numbermatchcorrectnumberentered":   AuthenticationAppPolicyDetailsStatus_NumberMatchCorrectNumberEntered,
		"numbermatchdeny":                   AuthenticationAppPolicyDetailsStatus_NumberMatchDeny,
		"numbermatchincorrectnumberentered": AuthenticationAppPolicyDetailsStatus_NumberMatchIncorrectNumberEntered,
		"numbermatchoutofdate":              AuthenticationAppPolicyDetailsStatus_NumberMatchOutOfDate,
		"tamperresistanthardwarenotused":    AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareNotUsed,
		"tamperresistanthardwareoutofdate":  AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareOutOfDate,
		"tamperresistanthardwareused":       AuthenticationAppPolicyDetailsStatus_TamperResistantHardwareUsed,
		"unknown":                           AuthenticationAppPolicyDetailsStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyDetailsStatus(input)
	return &out, nil
}
