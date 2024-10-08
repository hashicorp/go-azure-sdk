package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzurePostgreSqlLinkedServiceTypeProperties struct {
	ConnectionString    *string                       `json:"connectionString,omitempty"`
	EncryptedCredential *string                       `json:"encryptedCredential,omitempty"`
	Password            *AzureKeyVaultSecretReference `json:"password,omitempty"`
}
