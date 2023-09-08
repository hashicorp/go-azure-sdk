package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat string

const (
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonName                  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonName"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail           WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI            WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcustom                      WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonName),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcustom),
	}
}

func parseWindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat(input string) (*WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
