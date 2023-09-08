package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType string

const (
	AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypedigiCert      AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "DigiCert"
	AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypemicrosoft     AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "Microsoft"
	AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypenotConfigured AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType = "NotConfigured"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypedigiCert),
		string(AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypemicrosoft),
		string(AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypenotConfigured),
	}
}

func parseAospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input string) (*AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType{
		"digicert":      AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypedigiCert,
		"microsoft":     AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypemicrosoft,
		"notconfigured": AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityTypenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileCertificationAuthorityType(input)
	return &out, nil
}
