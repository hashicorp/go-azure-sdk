package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileSubjectNameFormat string

const (
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonName                  AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonName"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsEmail           AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIMEI            AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormatcustom                      AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonName),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAospDeviceOwnerScepCertificateProfileSubjectNameFormat(input string) (*AospDeviceOwnerScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileSubjectNameFormat{
		"commonname":                  AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AospDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AospDeviceOwnerScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
