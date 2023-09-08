package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileSubjectNameFormat string

const (
	AndroidScepCertificateProfileSubjectNameFormatcommonName                  AndroidScepCertificateProfileSubjectNameFormat = "CommonName"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidScepCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidScepCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidScepCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidScepCertificateProfileSubjectNameFormatcustom                      AndroidScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidScepCertificateProfileSubjectNameFormatcommonName),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidScepCertificateProfileSubjectNameFormat(input string) (*AndroidScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
