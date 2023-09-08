package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposesmimeEncryption AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "SmimeEncryption"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposesmimeSigning    AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "SmimeSigning"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposeunassigned      AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "Unassigned"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposevpn             AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "Vpn"
	AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposewifi            AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose = "Wifi"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposesmimeEncryption),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposesmimeSigning),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposeunassigned),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposevpn),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposewifi),
	}
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose{
		"smimeencryption": AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposesmimeEncryption,
		"smimesigning":    AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposesmimeSigning,
		"unassigned":      AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposeunassigned,
		"vpn":             AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposevpn,
		"wifi":            AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurposewifi,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileIntendedPurpose(input)
	return &out, nil
}
