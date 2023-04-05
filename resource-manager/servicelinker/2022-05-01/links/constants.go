package links

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthType string

const (
	AuthTypeSecret                      AuthType = "secret"
	AuthTypeServicePrincipalCertificate AuthType = "servicePrincipalCertificate"
	AuthTypeServicePrincipalSecret      AuthType = "servicePrincipalSecret"
	AuthTypeSystemAssignedIdentity      AuthType = "systemAssignedIdentity"
	AuthTypeUserAssignedIdentity        AuthType = "userAssignedIdentity"
)

func PossibleValuesForAuthType() []string {
	return []string{
		string(AuthTypeSecret),
		string(AuthTypeServicePrincipalCertificate),
		string(AuthTypeServicePrincipalSecret),
		string(AuthTypeSystemAssignedIdentity),
		string(AuthTypeUserAssignedIdentity),
	}
}

type AzureResourceType string

const (
	AzureResourceTypeKeyVault AzureResourceType = "KeyVault"
)

func PossibleValuesForAzureResourceType() []string {
	return []string{
		string(AzureResourceTypeKeyVault),
	}
}

type ClientType string

const (
	ClientTypeDjango                  ClientType = "django"
	ClientTypeDotnet                  ClientType = "dotnet"
	ClientTypeGo                      ClientType = "go"
	ClientTypeJava                    ClientType = "java"
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
		string(ClientTypeDjango),
		string(ClientTypeDotnet),
		string(ClientTypeGo),
		string(ClientTypeJava),
		string(ClientTypeKafkaNegativespringBoot),
		string(ClientTypeNodejs),
		string(ClientTypeNone),
		string(ClientTypePhp),
		string(ClientTypePython),
		string(ClientTypeRuby),
		string(ClientTypeSpringBoot),
	}
}

type SecretType string

const (
	SecretTypeKeyVaultSecretReference SecretType = "keyVaultSecretReference"
	SecretTypeKeyVaultSecretUri       SecretType = "keyVaultSecretUri"
	SecretTypeRawValue                SecretType = "rawValue"
)

func PossibleValuesForSecretType() []string {
	return []string{
		string(SecretTypeKeyVaultSecretReference),
		string(SecretTypeKeyVaultSecretUri),
		string(SecretTypeRawValue),
	}
}

type TargetServiceType string

const (
	TargetServiceTypeAzureResource            TargetServiceType = "AzureResource"
	TargetServiceTypeConfluentBootstrapServer TargetServiceType = "ConfluentBootstrapServer"
	TargetServiceTypeConfluentSchemaRegistry  TargetServiceType = "ConfluentSchemaRegistry"
)

func PossibleValuesForTargetServiceType() []string {
	return []string{
		string(TargetServiceTypeAzureResource),
		string(TargetServiceTypeConfluentBootstrapServer),
		string(TargetServiceTypeConfluentSchemaRegistry),
	}
}

type VNetSolutionType string

const (
	VNetSolutionTypePrivateLink     VNetSolutionType = "privateLink"
	VNetSolutionTypeServiceEndpoint VNetSolutionType = "serviceEndpoint"
)

func PossibleValuesForVNetSolutionType() []string {
	return []string{
		string(VNetSolutionTypePrivateLink),
		string(VNetSolutionTypeServiceEndpoint),
	}
}

type ValidationResultStatus string

const (
	ValidationResultStatusFailure ValidationResultStatus = "failure"
	ValidationResultStatusSuccess ValidationResultStatus = "success"
	ValidationResultStatusWarning ValidationResultStatus = "warning"
)

func PossibleValuesForValidationResultStatus() []string {
	return []string{
		string(ValidationResultStatusFailure),
		string(ValidationResultStatusSuccess),
		string(ValidationResultStatusWarning),
	}
}
