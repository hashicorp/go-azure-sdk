package storageaccountcredentials

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountCredentialProperties struct {
	AccountKey       *AsymmetricEncryptedSecret `json:"accountKey,omitempty"`
	AccountType      AccountType                `json:"accountType"`
	Alias            string                     `json:"alias"`
	BlobDomainName   *string                    `json:"blobDomainName,omitempty"`
	ConnectionString *string                    `json:"connectionString,omitempty"`
	SslStatus        SSLStatus                  `json:"sslStatus"`
	StorageAccountId *string                    `json:"storageAccountId,omitempty"`
	UserName         *string                    `json:"userName,omitempty"`
}
