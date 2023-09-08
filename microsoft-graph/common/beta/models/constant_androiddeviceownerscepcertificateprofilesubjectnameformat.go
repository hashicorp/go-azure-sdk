package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat string

const (
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonName                  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonName"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcustom                      AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonName),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidDeviceOwnerScepCertificateProfileSubjectNameFormat(input string) (*AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
