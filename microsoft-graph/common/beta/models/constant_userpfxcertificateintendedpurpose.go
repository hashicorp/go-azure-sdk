package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPFXCertificateIntendedPurpose string

const (
	UserPFXCertificateIntendedPurposesmimeEncryption UserPFXCertificateIntendedPurpose = "SmimeEncryption"
	UserPFXCertificateIntendedPurposesmimeSigning    UserPFXCertificateIntendedPurpose = "SmimeSigning"
	UserPFXCertificateIntendedPurposeunassigned      UserPFXCertificateIntendedPurpose = "Unassigned"
	UserPFXCertificateIntendedPurposevpn             UserPFXCertificateIntendedPurpose = "Vpn"
	UserPFXCertificateIntendedPurposewifi            UserPFXCertificateIntendedPurpose = "Wifi"
)

func PossibleValuesForUserPFXCertificateIntendedPurpose() []string {
	return []string{
		string(UserPFXCertificateIntendedPurposesmimeEncryption),
		string(UserPFXCertificateIntendedPurposesmimeSigning),
		string(UserPFXCertificateIntendedPurposeunassigned),
		string(UserPFXCertificateIntendedPurposevpn),
		string(UserPFXCertificateIntendedPurposewifi),
	}
}

func parseUserPFXCertificateIntendedPurpose(input string) (*UserPFXCertificateIntendedPurpose, error) {
	vals := map[string]UserPFXCertificateIntendedPurpose{
		"smimeencryption": UserPFXCertificateIntendedPurposesmimeEncryption,
		"smimesigning":    UserPFXCertificateIntendedPurposesmimeSigning,
		"unassigned":      UserPFXCertificateIntendedPurposeunassigned,
		"vpn":             UserPFXCertificateIntendedPurposevpn,
		"wifi":            UserPFXCertificateIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPFXCertificateIntendedPurpose(input)
	return &out, nil
}
