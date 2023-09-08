package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationSigningCertificateType string

const (
	IosEasEmailProfileConfigurationSigningCertificateTypecertificate       IosEasEmailProfileConfigurationSigningCertificateType = "Certificate"
	IosEasEmailProfileConfigurationSigningCertificateTypederivedCredential IosEasEmailProfileConfigurationSigningCertificateType = "DerivedCredential"
	IosEasEmailProfileConfigurationSigningCertificateTypenone              IosEasEmailProfileConfigurationSigningCertificateType = "None"
)

func PossibleValuesForIosEasEmailProfileConfigurationSigningCertificateType() []string {
	return []string{
		string(IosEasEmailProfileConfigurationSigningCertificateTypecertificate),
		string(IosEasEmailProfileConfigurationSigningCertificateTypederivedCredential),
		string(IosEasEmailProfileConfigurationSigningCertificateTypenone),
	}
}

func parseIosEasEmailProfileConfigurationSigningCertificateType(input string) (*IosEasEmailProfileConfigurationSigningCertificateType, error) {
	vals := map[string]IosEasEmailProfileConfigurationSigningCertificateType{
		"certificate":       IosEasEmailProfileConfigurationSigningCertificateTypecertificate,
		"derivedcredential": IosEasEmailProfileConfigurationSigningCertificateTypederivedCredential,
		"none":              IosEasEmailProfileConfigurationSigningCertificateTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationSigningCertificateType(input)
	return &out, nil
}
