package configurationnames

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthType string

const (
	AuthTypeAccessKey                   AuthType = "accessKey"
	AuthTypeEasyAuthMicrosoftEntraID    AuthType = "easyAuthMicrosoftEntraID"
	AuthTypeSecret                      AuthType = "secret"
	AuthTypeServicePrincipalCertificate AuthType = "servicePrincipalCertificate"
	AuthTypeServicePrincipalSecret      AuthType = "servicePrincipalSecret"
	AuthTypeSystemAssignedIdentity      AuthType = "systemAssignedIdentity"
	AuthTypeUserAccount                 AuthType = "userAccount"
	AuthTypeUserAssignedIdentity        AuthType = "userAssignedIdentity"
)

func PossibleValuesForAuthType() []string {
	return []string{
		string(AuthTypeAccessKey),
		string(AuthTypeEasyAuthMicrosoftEntraID),
		string(AuthTypeSecret),
		string(AuthTypeServicePrincipalCertificate),
		string(AuthTypeServicePrincipalSecret),
		string(AuthTypeSystemAssignedIdentity),
		string(AuthTypeUserAccount),
		string(AuthTypeUserAssignedIdentity),
	}
}

func (s *AuthType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthType(input string) (*AuthType, error) {
	vals := map[string]AuthType{
		"accesskey":                   AuthTypeAccessKey,
		"easyauthmicrosoftentraid":    AuthTypeEasyAuthMicrosoftEntraID,
		"secret":                      AuthTypeSecret,
		"serviceprincipalcertificate": AuthTypeServicePrincipalCertificate,
		"serviceprincipalsecret":      AuthTypeServicePrincipalSecret,
		"systemassignedidentity":      AuthTypeSystemAssignedIdentity,
		"useraccount":                 AuthTypeUserAccount,
		"userassignedidentity":        AuthTypeUserAssignedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthType(input)
	return &out, nil
}

type ClientType string

const (
	ClientTypeDapr                    ClientType = "dapr"
	ClientTypeDjango                  ClientType = "django"
	ClientTypeDotnet                  ClientType = "dotnet"
	ClientTypeGo                      ClientType = "go"
	ClientTypeJava                    ClientType = "java"
	ClientTypeJmsNegativespringBoot   ClientType = "jms-springBoot"
	ClientTypeKafkaNegativespringBoot ClientType = "kafka-springBoot"
	ClientTypeNodejs                  ClientType = "nodejs"
	ClientTypeNone                    ClientType = "none"
	ClientTypePhp                     ClientType = "php"
	ClientTypePython                  ClientType = "python"
	ClientTypeRuby                    ClientType = "ruby"
	ClientTypeSpringBoot              ClientType = "springBoot"
)

func PossibleValuesForClientType() []string {
	return []string{
		string(ClientTypeDapr),
		string(ClientTypeDjango),
		string(ClientTypeDotnet),
		string(ClientTypeGo),
		string(ClientTypeJava),
		string(ClientTypeJmsNegativespringBoot),
		string(ClientTypeKafkaNegativespringBoot),
		string(ClientTypeNodejs),
		string(ClientTypeNone),
		string(ClientTypePhp),
		string(ClientTypePython),
		string(ClientTypeRuby),
		string(ClientTypeSpringBoot),
	}
}

func (s *ClientType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseClientType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseClientType(input string) (*ClientType, error) {
	vals := map[string]ClientType{
		"dapr":             ClientTypeDapr,
		"django":           ClientTypeDjango,
		"dotnet":           ClientTypeDotnet,
		"go":               ClientTypeGo,
		"java":             ClientTypeJava,
		"jms-springboot":   ClientTypeJmsNegativespringBoot,
		"kafka-springboot": ClientTypeKafkaNegativespringBoot,
		"nodejs":           ClientTypeNodejs,
		"none":             ClientTypeNone,
		"php":              ClientTypePhp,
		"python":           ClientTypePython,
		"ruby":             ClientTypeRuby,
		"springboot":       ClientTypeSpringBoot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClientType(input)
	return &out, nil
}

type DaprBindingComponentDirection string

const (
	DaprBindingComponentDirectionInput  DaprBindingComponentDirection = "input"
	DaprBindingComponentDirectionOutput DaprBindingComponentDirection = "output"
)

func PossibleValuesForDaprBindingComponentDirection() []string {
	return []string{
		string(DaprBindingComponentDirectionInput),
		string(DaprBindingComponentDirectionOutput),
	}
}

func (s *DaprBindingComponentDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDaprBindingComponentDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDaprBindingComponentDirection(input string) (*DaprBindingComponentDirection, error) {
	vals := map[string]DaprBindingComponentDirection{
		"input":  DaprBindingComponentDirectionInput,
		"output": DaprBindingComponentDirectionOutput,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DaprBindingComponentDirection(input)
	return &out, nil
}

type DaprMetadataRequired string

const (
	DaprMetadataRequiredFalse DaprMetadataRequired = "false"
	DaprMetadataRequiredTrue  DaprMetadataRequired = "true"
)

func PossibleValuesForDaprMetadataRequired() []string {
	return []string{
		string(DaprMetadataRequiredFalse),
		string(DaprMetadataRequiredTrue),
	}
}

func (s *DaprMetadataRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDaprMetadataRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDaprMetadataRequired(input string) (*DaprMetadataRequired, error) {
	vals := map[string]DaprMetadataRequired{
		"false": DaprMetadataRequiredFalse,
		"true":  DaprMetadataRequiredTrue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DaprMetadataRequired(input)
	return &out, nil
}

type SecretSourceType string

const (
	SecretSourceTypeKeyVaultSecret SecretSourceType = "keyVaultSecret"
	SecretSourceTypeRawValue       SecretSourceType = "rawValue"
)

func PossibleValuesForSecretSourceType() []string {
	return []string{
		string(SecretSourceTypeKeyVaultSecret),
		string(SecretSourceTypeRawValue),
	}
}

func (s *SecretSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecretSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecretSourceType(input string) (*SecretSourceType, error) {
	vals := map[string]SecretSourceType{
		"keyvaultsecret": SecretSourceTypeKeyVaultSecret,
		"rawvalue":       SecretSourceTypeRawValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecretSourceType(input)
	return &out, nil
}
