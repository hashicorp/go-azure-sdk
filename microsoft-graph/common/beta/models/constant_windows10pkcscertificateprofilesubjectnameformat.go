package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileSubjectNameFormat string

const (
	Windows10PkcsCertificateProfileSubjectNameFormatcommonName                  Windows10PkcsCertificateProfileSubjectNameFormat = "CommonName"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsEmail           Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI            Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	Windows10PkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail    Windows10PkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	Windows10PkcsCertificateProfileSubjectNameFormatcustom                      Windows10PkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForWindows10PkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonName),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(Windows10PkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseWindows10PkcsCertificateProfileSubjectNameFormat(input string) (*Windows10PkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]Windows10PkcsCertificateProfileSubjectNameFormat{
		"commonname":                  Windows10PkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows10PkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows10PkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      Windows10PkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
