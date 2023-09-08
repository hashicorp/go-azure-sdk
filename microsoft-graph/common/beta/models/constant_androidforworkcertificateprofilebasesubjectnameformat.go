package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBaseSubjectNameFormat string

const (
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonName                  AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonName"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsEmail           AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    AndroidForWorkCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidForWorkCertificateProfileBaseSubjectNameFormatcustom                      AndroidForWorkCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidForWorkCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonName),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseAndroidForWorkCertificateProfileBaseSubjectNameFormat(input string) (*AndroidForWorkCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidForWorkCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
