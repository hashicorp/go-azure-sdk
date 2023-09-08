package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileSubjectNameFormat string

const (
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonName                  Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonName"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail           Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI            Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail    Windows10ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	Windows10ImportedPFXCertificateProfileSubjectNameFormatcustom                      Windows10ImportedPFXCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonName),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormatcustom),
	}
}

func parseWindows10ImportedPFXCertificateProfileSubjectNameFormat(input string) (*Windows10ImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows10ImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      Windows10ImportedPFXCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
