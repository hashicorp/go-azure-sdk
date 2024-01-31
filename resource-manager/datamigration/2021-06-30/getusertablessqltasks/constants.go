package getusertablessqltasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeActiveDirectoryIntegrated AuthenticationType = "ActiveDirectoryIntegrated"
	AuthenticationTypeActiveDirectoryPassword   AuthenticationType = "ActiveDirectoryPassword"
	AuthenticationTypeNone                      AuthenticationType = "None"
	AuthenticationTypeSqlAuthentication         AuthenticationType = "SqlAuthentication"
	AuthenticationTypeWindowsAuthentication     AuthenticationType = "WindowsAuthentication"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeActiveDirectoryIntegrated),
		string(AuthenticationTypeActiveDirectoryPassword),
		string(AuthenticationTypeNone),
		string(AuthenticationTypeSqlAuthentication),
		string(AuthenticationTypeWindowsAuthentication),
	}
}

func (s *AuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationType(input string) (*AuthenticationType, error) {
	vals := map[string]AuthenticationType{
		"activedirectoryintegrated": AuthenticationTypeActiveDirectoryIntegrated,
		"activedirectorypassword":   AuthenticationTypeActiveDirectoryPassword,
		"none":                      AuthenticationTypeNone,
		"sqlauthentication":         AuthenticationTypeSqlAuthentication,
		"windowsauthentication":     AuthenticationTypeWindowsAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationType(input)
	return &out, nil
}

type SqlSourcePlatform string

const (
	SqlSourcePlatformSqlOnPrem SqlSourcePlatform = "SqlOnPrem"
)

func PossibleValuesForSqlSourcePlatform() []string {
	return []string{
		string(SqlSourcePlatformSqlOnPrem),
	}
}

func (s *SqlSourcePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlSourcePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlSourcePlatform(input string) (*SqlSourcePlatform, error) {
	vals := map[string]SqlSourcePlatform{
		"sqlonprem": SqlSourcePlatformSqlOnPrem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlSourcePlatform(input)
	return &out, nil
}
