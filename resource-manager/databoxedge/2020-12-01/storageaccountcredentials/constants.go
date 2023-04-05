package storageaccountcredentials

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountType string

const (
	AccountTypeBlobStorage           AccountType = "BlobStorage"
	AccountTypeGeneralPurposeStorage AccountType = "GeneralPurposeStorage"
)

func PossibleValuesForAccountType() []string {
	return []string{
		string(AccountTypeBlobStorage),
		string(AccountTypeGeneralPurposeStorage),
	}
}

type EncryptionAlgorithm string

const (
	EncryptionAlgorithmAESTwoFiveSix        EncryptionAlgorithm = "AES256"
	EncryptionAlgorithmNone                 EncryptionAlgorithm = "None"
	EncryptionAlgorithmRSAESPKCSOneVOneFive EncryptionAlgorithm = "RSAES_PKCS1_v_1_5"
)

func PossibleValuesForEncryptionAlgorithm() []string {
	return []string{
		string(EncryptionAlgorithmAESTwoFiveSix),
		string(EncryptionAlgorithmNone),
		string(EncryptionAlgorithmRSAESPKCSOneVOneFive),
	}
}

type SSLStatus string

const (
	SSLStatusDisabled SSLStatus = "Disabled"
	SSLStatusEnabled  SSLStatus = "Enabled"
)

func PossibleValuesForSSLStatus() []string {
	return []string{
		string(SSLStatusDisabled),
		string(SSLStatusEnabled),
	}
}
