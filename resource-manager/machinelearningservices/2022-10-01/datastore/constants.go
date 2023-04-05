package datastore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialsType string

const (
	CredentialsTypeAccountKey       CredentialsType = "AccountKey"
	CredentialsTypeCertificate      CredentialsType = "Certificate"
	CredentialsTypeNone             CredentialsType = "None"
	CredentialsTypeSas              CredentialsType = "Sas"
	CredentialsTypeServicePrincipal CredentialsType = "ServicePrincipal"
)

func PossibleValuesForCredentialsType() []string {
	return []string{
		string(CredentialsTypeAccountKey),
		string(CredentialsTypeCertificate),
		string(CredentialsTypeNone),
		string(CredentialsTypeSas),
		string(CredentialsTypeServicePrincipal),
	}
}

type DatastoreType string

const (
	DatastoreTypeAzureBlob           DatastoreType = "AzureBlob"
	DatastoreTypeAzureDataLakeGenOne DatastoreType = "AzureDataLakeGen1"
	DatastoreTypeAzureDataLakeGenTwo DatastoreType = "AzureDataLakeGen2"
	DatastoreTypeAzureFile           DatastoreType = "AzureFile"
)

func PossibleValuesForDatastoreType() []string {
	return []string{
		string(DatastoreTypeAzureBlob),
		string(DatastoreTypeAzureDataLakeGenOne),
		string(DatastoreTypeAzureDataLakeGenTwo),
		string(DatastoreTypeAzureFile),
	}
}

type SecretsType string

const (
	SecretsTypeAccountKey       SecretsType = "AccountKey"
	SecretsTypeCertificate      SecretsType = "Certificate"
	SecretsTypeSas              SecretsType = "Sas"
	SecretsTypeServicePrincipal SecretsType = "ServicePrincipal"
)

func PossibleValuesForSecretsType() []string {
	return []string{
		string(SecretsTypeAccountKey),
		string(SecretsTypeCertificate),
		string(SecretsTypeSas),
		string(SecretsTypeServicePrincipal),
	}
}

type ServiceDataAccessAuthIdentity string

const (
	ServiceDataAccessAuthIdentityNone                            ServiceDataAccessAuthIdentity = "None"
	ServiceDataAccessAuthIdentityWorkspaceSystemAssignedIdentity ServiceDataAccessAuthIdentity = "WorkspaceSystemAssignedIdentity"
	ServiceDataAccessAuthIdentityWorkspaceUserAssignedIdentity   ServiceDataAccessAuthIdentity = "WorkspaceUserAssignedIdentity"
)

func PossibleValuesForServiceDataAccessAuthIdentity() []string {
	return []string{
		string(ServiceDataAccessAuthIdentityNone),
		string(ServiceDataAccessAuthIdentityWorkspaceSystemAssignedIdentity),
		string(ServiceDataAccessAuthIdentityWorkspaceUserAssignedIdentity),
	}
}
