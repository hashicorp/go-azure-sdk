package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPFXCertificatePaddingScheme string

const (
	UserPFXCertificatePaddingSchemenone       UserPFXCertificatePaddingScheme = "None"
	UserPFXCertificatePaddingSchemeoaepSha1   UserPFXCertificatePaddingScheme = "OaepSha1"
	UserPFXCertificatePaddingSchemeoaepSha256 UserPFXCertificatePaddingScheme = "OaepSha256"
	UserPFXCertificatePaddingSchemeoaepSha384 UserPFXCertificatePaddingScheme = "OaepSha384"
	UserPFXCertificatePaddingSchemeoaepSha512 UserPFXCertificatePaddingScheme = "OaepSha512"
	UserPFXCertificatePaddingSchemepkcs1      UserPFXCertificatePaddingScheme = "Pkcs1"
)

func PossibleValuesForUserPFXCertificatePaddingScheme() []string {
	return []string{
		string(UserPFXCertificatePaddingSchemenone),
		string(UserPFXCertificatePaddingSchemeoaepSha1),
		string(UserPFXCertificatePaddingSchemeoaepSha256),
		string(UserPFXCertificatePaddingSchemeoaepSha384),
		string(UserPFXCertificatePaddingSchemeoaepSha512),
		string(UserPFXCertificatePaddingSchemepkcs1),
	}
}

func parseUserPFXCertificatePaddingScheme(input string) (*UserPFXCertificatePaddingScheme, error) {
	vals := map[string]UserPFXCertificatePaddingScheme{
		"none":       UserPFXCertificatePaddingSchemenone,
		"oaepsha1":   UserPFXCertificatePaddingSchemeoaepSha1,
		"oaepsha256": UserPFXCertificatePaddingSchemeoaepSha256,
		"oaepsha384": UserPFXCertificatePaddingSchemeoaepSha384,
		"oaepsha512": UserPFXCertificatePaddingSchemeoaepSha512,
		"pkcs1":      UserPFXCertificatePaddingSchemepkcs1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPFXCertificatePaddingScheme(input)
	return &out, nil
}
