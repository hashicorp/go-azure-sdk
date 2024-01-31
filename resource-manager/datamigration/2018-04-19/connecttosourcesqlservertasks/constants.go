package connecttosourcesqlservertasks

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

type DatabaseCompatLevel string

const (
	DatabaseCompatLevelCompatLevelEightZero    DatabaseCompatLevel = "CompatLevel80"
	DatabaseCompatLevelCompatLevelNineZero     DatabaseCompatLevel = "CompatLevel90"
	DatabaseCompatLevelCompatLevelOneFourZero  DatabaseCompatLevel = "CompatLevel140"
	DatabaseCompatLevelCompatLevelOneHundred   DatabaseCompatLevel = "CompatLevel100"
	DatabaseCompatLevelCompatLevelOneOneZero   DatabaseCompatLevel = "CompatLevel110"
	DatabaseCompatLevelCompatLevelOneThreeZero DatabaseCompatLevel = "CompatLevel130"
	DatabaseCompatLevelCompatLevelOneTwoZero   DatabaseCompatLevel = "CompatLevel120"
)

func PossibleValuesForDatabaseCompatLevel() []string {
	return []string{
		string(DatabaseCompatLevelCompatLevelEightZero),
		string(DatabaseCompatLevelCompatLevelNineZero),
		string(DatabaseCompatLevelCompatLevelOneFourZero),
		string(DatabaseCompatLevelCompatLevelOneHundred),
		string(DatabaseCompatLevelCompatLevelOneOneZero),
		string(DatabaseCompatLevelCompatLevelOneThreeZero),
		string(DatabaseCompatLevelCompatLevelOneTwoZero),
	}
}

func (s *DatabaseCompatLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseCompatLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseCompatLevel(input string) (*DatabaseCompatLevel, error) {
	vals := map[string]DatabaseCompatLevel{
		"compatlevel80":  DatabaseCompatLevelCompatLevelEightZero,
		"compatlevel90":  DatabaseCompatLevelCompatLevelNineZero,
		"compatlevel140": DatabaseCompatLevelCompatLevelOneFourZero,
		"compatlevel100": DatabaseCompatLevelCompatLevelOneHundred,
		"compatlevel110": DatabaseCompatLevelCompatLevelOneOneZero,
		"compatlevel130": DatabaseCompatLevelCompatLevelOneThreeZero,
		"compatlevel120": DatabaseCompatLevelCompatLevelOneTwoZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseCompatLevel(input)
	return &out, nil
}

type DatabaseFileType string

const (
	DatabaseFileTypeFilestream   DatabaseFileType = "Filestream"
	DatabaseFileTypeFulltext     DatabaseFileType = "Fulltext"
	DatabaseFileTypeLog          DatabaseFileType = "Log"
	DatabaseFileTypeNotSupported DatabaseFileType = "NotSupported"
	DatabaseFileTypeRows         DatabaseFileType = "Rows"
)

func PossibleValuesForDatabaseFileType() []string {
	return []string{
		string(DatabaseFileTypeFilestream),
		string(DatabaseFileTypeFulltext),
		string(DatabaseFileTypeLog),
		string(DatabaseFileTypeNotSupported),
		string(DatabaseFileTypeRows),
	}
}

func (s *DatabaseFileType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseFileType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseFileType(input string) (*DatabaseFileType, error) {
	vals := map[string]DatabaseFileType{
		"filestream":   DatabaseFileTypeFilestream,
		"fulltext":     DatabaseFileTypeFulltext,
		"log":          DatabaseFileTypeLog,
		"notsupported": DatabaseFileTypeNotSupported,
		"rows":         DatabaseFileTypeRows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseFileType(input)
	return &out, nil
}

type DatabaseState string

const (
	DatabaseStateCopying          DatabaseState = "Copying"
	DatabaseStateEmergency        DatabaseState = "Emergency"
	DatabaseStateOffline          DatabaseState = "Offline"
	DatabaseStateOfflineSecondary DatabaseState = "OfflineSecondary"
	DatabaseStateOnline           DatabaseState = "Online"
	DatabaseStateRecovering       DatabaseState = "Recovering"
	DatabaseStateRecoveryPending  DatabaseState = "RecoveryPending"
	DatabaseStateRestoring        DatabaseState = "Restoring"
	DatabaseStateSuspect          DatabaseState = "Suspect"
)

func PossibleValuesForDatabaseState() []string {
	return []string{
		string(DatabaseStateCopying),
		string(DatabaseStateEmergency),
		string(DatabaseStateOffline),
		string(DatabaseStateOfflineSecondary),
		string(DatabaseStateOnline),
		string(DatabaseStateRecovering),
		string(DatabaseStateRecoveryPending),
		string(DatabaseStateRestoring),
		string(DatabaseStateSuspect),
	}
}

func (s *DatabaseState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDatabaseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDatabaseState(input string) (*DatabaseState, error) {
	vals := map[string]DatabaseState{
		"copying":          DatabaseStateCopying,
		"emergency":        DatabaseStateEmergency,
		"offline":          DatabaseStateOffline,
		"offlinesecondary": DatabaseStateOfflineSecondary,
		"online":           DatabaseStateOnline,
		"recovering":       DatabaseStateRecovering,
		"recoverypending":  DatabaseStateRecoveryPending,
		"restoring":        DatabaseStateRestoring,
		"suspect":          DatabaseStateSuspect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DatabaseState(input)
	return &out, nil
}

type LoginType string

const (
	LoginTypeAsymmetricKey LoginType = "AsymmetricKey"
	LoginTypeCertificate   LoginType = "Certificate"
	LoginTypeExternalGroup LoginType = "ExternalGroup"
	LoginTypeExternalUser  LoginType = "ExternalUser"
	LoginTypeSqlLogin      LoginType = "SqlLogin"
	LoginTypeWindowsGroup  LoginType = "WindowsGroup"
	LoginTypeWindowsUser   LoginType = "WindowsUser"
)

func PossibleValuesForLoginType() []string {
	return []string{
		string(LoginTypeAsymmetricKey),
		string(LoginTypeCertificate),
		string(LoginTypeExternalGroup),
		string(LoginTypeExternalUser),
		string(LoginTypeSqlLogin),
		string(LoginTypeWindowsGroup),
		string(LoginTypeWindowsUser),
	}
}

func (s *LoginType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLoginType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLoginType(input string) (*LoginType, error) {
	vals := map[string]LoginType{
		"asymmetrickey": LoginTypeAsymmetricKey,
		"certificate":   LoginTypeCertificate,
		"externalgroup": LoginTypeExternalGroup,
		"externaluser":  LoginTypeExternalUser,
		"sqllogin":      LoginTypeSqlLogin,
		"windowsgroup":  LoginTypeWindowsGroup,
		"windowsuser":   LoginTypeWindowsUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LoginType(input)
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
