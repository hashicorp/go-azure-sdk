package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseSubjectNameFormat string

const (
	Windows81CertificateProfileBaseSubjectNameFormatcommonName                  Windows81CertificateProfileBaseSubjectNameFormat = "CommonName"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsEmail           Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	Windows81CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    Windows81CertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	Windows81CertificateProfileBaseSubjectNameFormatcustom                      Windows81CertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForWindows81CertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonName),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(Windows81CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(Windows81CertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseWindows81CertificateProfileBaseSubjectNameFormat(input string) (*Windows81CertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]Windows81CertificateProfileBaseSubjectNameFormat{
		"commonname":                  Windows81CertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows81CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows81CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      Windows81CertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
