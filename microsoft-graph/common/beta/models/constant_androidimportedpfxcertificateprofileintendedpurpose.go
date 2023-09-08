package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileIntendedPurpose string

const (
	AndroidImportedPFXCertificateProfileIntendedPurposesmimeEncryption AndroidImportedPFXCertificateProfileIntendedPurpose = "SmimeEncryption"
	AndroidImportedPFXCertificateProfileIntendedPurposesmimeSigning    AndroidImportedPFXCertificateProfileIntendedPurpose = "SmimeSigning"
	AndroidImportedPFXCertificateProfileIntendedPurposeunassigned      AndroidImportedPFXCertificateProfileIntendedPurpose = "Unassigned"
	AndroidImportedPFXCertificateProfileIntendedPurposevpn             AndroidImportedPFXCertificateProfileIntendedPurpose = "Vpn"
	AndroidImportedPFXCertificateProfileIntendedPurposewifi            AndroidImportedPFXCertificateProfileIntendedPurpose = "Wifi"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileIntendedPurposesmimeEncryption),
		string(AndroidImportedPFXCertificateProfileIntendedPurposesmimeSigning),
		string(AndroidImportedPFXCertificateProfileIntendedPurposeunassigned),
		string(AndroidImportedPFXCertificateProfileIntendedPurposevpn),
		string(AndroidImportedPFXCertificateProfileIntendedPurposewifi),
	}
}

func parseAndroidImportedPFXCertificateProfileIntendedPurpose(input string) (*AndroidImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": AndroidImportedPFXCertificateProfileIntendedPurposesmimeEncryption,
		"smimesigning":    AndroidImportedPFXCertificateProfileIntendedPurposesmimeSigning,
		"unassigned":      AndroidImportedPFXCertificateProfileIntendedPurposeunassigned,
		"vpn":             AndroidImportedPFXCertificateProfileIntendedPurposevpn,
		"wifi":            AndroidImportedPFXCertificateProfileIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
