package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessTypespecificApps AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType = "SpecificApps"
	AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessTypeuserApproval AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType = "UserApproval"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessTypespecificApps),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessTypeuserApproval),
	}
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType{
		"specificapps": AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessTypespecificApps,
		"userapproval": AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessTypeuserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileCertificateAccessType(input)
	return &out, nil
}
