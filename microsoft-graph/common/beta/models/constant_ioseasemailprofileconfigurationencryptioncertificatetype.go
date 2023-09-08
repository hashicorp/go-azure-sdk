package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationEncryptionCertificateType string

const (
	IosEasEmailProfileConfigurationEncryptionCertificateTypecertificate       IosEasEmailProfileConfigurationEncryptionCertificateType = "Certificate"
	IosEasEmailProfileConfigurationEncryptionCertificateTypederivedCredential IosEasEmailProfileConfigurationEncryptionCertificateType = "DerivedCredential"
	IosEasEmailProfileConfigurationEncryptionCertificateTypenone              IosEasEmailProfileConfigurationEncryptionCertificateType = "None"
)

func PossibleValuesForIosEasEmailProfileConfigurationEncryptionCertificateType() []string {
	return []string{
		string(IosEasEmailProfileConfigurationEncryptionCertificateTypecertificate),
		string(IosEasEmailProfileConfigurationEncryptionCertificateTypederivedCredential),
		string(IosEasEmailProfileConfigurationEncryptionCertificateTypenone),
	}
}

func parseIosEasEmailProfileConfigurationEncryptionCertificateType(input string) (*IosEasEmailProfileConfigurationEncryptionCertificateType, error) {
	vals := map[string]IosEasEmailProfileConfigurationEncryptionCertificateType{
		"certificate":       IosEasEmailProfileConfigurationEncryptionCertificateTypecertificate,
		"derivedcredential": IosEasEmailProfileConfigurationEncryptionCertificateTypederivedCredential,
		"none":              IosEasEmailProfileConfigurationEncryptionCertificateTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationEncryptionCertificateType(input)
	return &out, nil
}
