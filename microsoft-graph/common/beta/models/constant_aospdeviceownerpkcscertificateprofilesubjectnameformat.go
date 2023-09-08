package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat string

const (
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonName                  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonName"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail           AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI            AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcustom                      AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonName),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAospDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input string) (*AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AospDeviceOwnerPkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
