package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileIntendedPurpose string

const (
	AndroidForWorkImportedPFXCertificateProfileIntendedPurposesmimeEncryption AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "SmimeEncryption"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurposesmimeSigning    AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "SmimeSigning"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurposeunassigned      AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "Unassigned"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurposevpn             AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "Vpn"
	AndroidForWorkImportedPFXCertificateProfileIntendedPurposewifi            AndroidForWorkImportedPFXCertificateProfileIntendedPurpose = "Wifi"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurposesmimeEncryption),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurposesmimeSigning),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurposeunassigned),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurposevpn),
		string(AndroidForWorkImportedPFXCertificateProfileIntendedPurposewifi),
	}
}

func parseAndroidForWorkImportedPFXCertificateProfileIntendedPurpose(input string) (*AndroidForWorkImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": AndroidForWorkImportedPFXCertificateProfileIntendedPurposesmimeEncryption,
		"smimesigning":    AndroidForWorkImportedPFXCertificateProfileIntendedPurposesmimeSigning,
		"unassigned":      AndroidForWorkImportedPFXCertificateProfileIntendedPurposeunassigned,
		"vpn":             AndroidForWorkImportedPFXCertificateProfileIntendedPurposevpn,
		"wifi":            AndroidForWorkImportedPFXCertificateProfileIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
