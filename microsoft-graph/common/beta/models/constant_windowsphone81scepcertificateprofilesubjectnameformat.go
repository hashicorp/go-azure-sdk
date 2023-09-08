package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileSubjectNameFormat string

const (
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonName                  WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonName"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsEmail           WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsIMEI            WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameIncludingEmail    WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormatcustom                      WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonName),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormatcustom),
	}
}

func parseWindowsPhone81SCEPCertificateProfileSubjectNameFormat(input string) (*WindowsPhone81SCEPCertificateProfileSubjectNameFormat, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileSubjectNameFormat{
		"commonname":                  WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsPhone81SCEPCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      WindowsPhone81SCEPCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
