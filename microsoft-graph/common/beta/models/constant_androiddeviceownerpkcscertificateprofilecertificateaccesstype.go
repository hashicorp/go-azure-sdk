package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessTypespecificApps AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType = "SpecificApps"
	AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessTypeuserApproval AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType = "UserApproval"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessTypespecificApps),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessTypeuserApproval),
	}
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType{
		"specificapps": AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessTypespecificApps,
		"userapproval": AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessTypeuserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificateAccessType(input)
	return &out, nil
}
