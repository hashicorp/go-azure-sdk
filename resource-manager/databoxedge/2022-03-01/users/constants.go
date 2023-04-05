package users

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type ShareAccessType string

const (
	ShareAccessTypeChange ShareAccessType = "Change"
	ShareAccessTypeCustom ShareAccessType = "Custom"
	ShareAccessTypeRead   ShareAccessType = "Read"
)

func PossibleValuesForShareAccessType() []string {
	return []string{
		string(ShareAccessTypeChange),
		string(ShareAccessTypeCustom),
		string(ShareAccessTypeRead),
	}
}

type UserType string

const (
	UserTypeARM             UserType = "ARM"
	UserTypeLocalManagement UserType = "LocalManagement"
	UserTypeShare           UserType = "Share"
)

func PossibleValuesForUserType() []string {
	return []string{
		string(UserTypeARM),
		string(UserTypeLocalManagement),
		string(UserTypeShare),
	}
}
