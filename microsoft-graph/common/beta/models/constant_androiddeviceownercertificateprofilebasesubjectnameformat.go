package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat string

const (
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonName                  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonName"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsEmail           AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcustom                      AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonName),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseAndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat(input string) (*AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
