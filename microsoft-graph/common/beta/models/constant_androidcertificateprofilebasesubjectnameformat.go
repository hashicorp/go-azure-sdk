package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCertificateProfileBaseSubjectNameFormat string

const (
	AndroidCertificateProfileBaseSubjectNameFormatcommonName                  AndroidCertificateProfileBaseSubjectNameFormat = "CommonName"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsEmail           AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    AndroidCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidCertificateProfileBaseSubjectNameFormatcustom                      AndroidCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonName),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseAndroidCertificateProfileBaseSubjectNameFormat(input string) (*AndroidCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
