package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseSubjectNameFormat string

const (
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonName                  WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonName"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsEmail           WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    WindowsPhone81CertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	WindowsPhone81CertificateProfileBaseSubjectNameFormatcustom                      WindowsPhone81CertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonName),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseWindowsPhone81CertificateProfileBaseSubjectNameFormat(input string) (*WindowsPhone81CertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseSubjectNameFormat{
		"commonname":                  WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsPhone81CertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      WindowsPhone81CertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
