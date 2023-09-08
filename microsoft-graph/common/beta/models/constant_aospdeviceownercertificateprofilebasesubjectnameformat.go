package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCertificateProfileBaseSubjectNameFormat string

const (
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonName                  AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonName"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsEmail           AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcustom                      AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForAospDeviceOwnerCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonName),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseAospDeviceOwnerCertificateProfileBaseSubjectNameFormat(input string) (*AospDeviceOwnerCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AospDeviceOwnerCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AospDeviceOwnerCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
