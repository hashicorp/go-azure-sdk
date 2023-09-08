package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType string

const (
	AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypedigiCert      AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "DigiCert"
	AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypemicrosoft     AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "Microsoft"
	AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypenotConfigured AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "NotConfigured"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypedigiCert),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypemicrosoft),
		string(AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypenotConfigured),
	}
}

func parseAndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input string) (*AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType{
		"digicert":      AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypedigiCert,
		"microsoft":     AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypemicrosoft,
		"notconfigured": AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input)
	return &out, nil
}
