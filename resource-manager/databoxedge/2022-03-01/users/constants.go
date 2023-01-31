package users

import "strings"

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

func parseShareAccessType(input string) (*ShareAccessType, error) {
	vals := map[string]ShareAccessType{
		"change": ShareAccessTypeChange,
		"custom": ShareAccessTypeCustom,
		"read":   ShareAccessTypeRead,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ShareAccessType(input)
	return &out, nil
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

func parseUserType(input string) (*UserType, error) {
	vals := map[string]UserType{
		"arm":             UserTypeARM,
		"localmanagement": UserTypeLocalManagement,
		"share":           UserTypeShare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserType(input)
	return &out, nil
}
