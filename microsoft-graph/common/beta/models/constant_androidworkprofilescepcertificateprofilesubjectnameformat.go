package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileSubjectNameFormat string

const (
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonName                  AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonName"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormatcustom                      AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonName),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidWorkProfileScepCertificateProfileSubjectNameFormat(input string) (*AndroidWorkProfileScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidWorkProfileScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidWorkProfileScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
