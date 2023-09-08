package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseSubjectNameFormat string

const (
	WindowsCertificateProfileBaseSubjectNameFormatcommonName                  WindowsCertificateProfileBaseSubjectNameFormat = "CommonName"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsEmail           WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	WindowsCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    WindowsCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	WindowsCertificateProfileBaseSubjectNameFormatcustom                      WindowsCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForWindowsCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonName),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(WindowsCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(WindowsCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseWindowsCertificateProfileBaseSubjectNameFormat(input string) (*WindowsCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]WindowsCertificateProfileBaseSubjectNameFormat{
		"commonname":                  WindowsCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      WindowsCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
