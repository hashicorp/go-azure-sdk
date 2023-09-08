package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileCertificateAccessType string

const (
	AndroidDeviceOwnerScepCertificateProfileCertificateAccessTypespecificApps AndroidDeviceOwnerScepCertificateProfileCertificateAccessType = "SpecificApps"
	AndroidDeviceOwnerScepCertificateProfileCertificateAccessTypeuserApproval AndroidDeviceOwnerScepCertificateProfileCertificateAccessType = "UserApproval"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileCertificateAccessTypespecificApps),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateAccessTypeuserApproval),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileCertificateAccessType(input string) (*AndroidDeviceOwnerScepCertificateProfileCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileCertificateAccessType{
		"specificapps": AndroidDeviceOwnerScepCertificateProfileCertificateAccessTypespecificApps,
		"userapproval": AndroidDeviceOwnerScepCertificateProfileCertificateAccessTypeuserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileCertificateAccessType(input)
	return &out, nil
}
