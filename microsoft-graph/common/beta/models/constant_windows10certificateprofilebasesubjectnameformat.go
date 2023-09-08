package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseSubjectNameFormat string

const (
	Windows10CertificateProfileBaseSubjectNameFormatcommonName                  Windows10CertificateProfileBaseSubjectNameFormat = "CommonName"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsEmail           Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	Windows10CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    Windows10CertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	Windows10CertificateProfileBaseSubjectNameFormatcustom                      Windows10CertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForWindows10CertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonName),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(Windows10CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(Windows10CertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseWindows10CertificateProfileBaseSubjectNameFormat(input string) (*Windows10CertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]Windows10CertificateProfileBaseSubjectNameFormat{
		"commonname":                  Windows10CertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows10CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows10CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      Windows10CertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
