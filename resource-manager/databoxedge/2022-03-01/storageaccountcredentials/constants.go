package storageaccountcredentials

import "strings"

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

func parseAccountType(input string) (*AccountType, error) {
	vals := map[string]AccountType{
		"blobstorage":           AccountTypeBlobStorage,
		"generalpurposestorage": AccountTypeGeneralPurposeStorage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccountType(input)
	return &out, nil
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

func parseEncryptionAlgorithm(input string) (*EncryptionAlgorithm, error) {
	vals := map[string]EncryptionAlgorithm{
		"aes256":            EncryptionAlgorithmAESTwoFiveSix,
		"none":              EncryptionAlgorithmNone,
		"rsaes_pkcs1_v_1_5": EncryptionAlgorithmRSAESPKCSOneVOneFive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionAlgorithm(input)
	return &out, nil
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

func parseSSLStatus(input string) (*SSLStatus, error) {
	vals := map[string]SSLStatus{
		"disabled": SSLStatusDisabled,
		"enabled":  SSLStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SSLStatus(input)
	return &out, nil
}
