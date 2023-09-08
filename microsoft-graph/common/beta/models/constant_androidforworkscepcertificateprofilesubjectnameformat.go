package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileSubjectNameFormat string

const (
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonName                  AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonName"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidForWorkScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidForWorkScepCertificateProfileSubjectNameFormatcustom                      AndroidForWorkScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonName),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidForWorkScepCertificateProfileSubjectNameFormat(input string) (*AndroidForWorkScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidForWorkScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidForWorkScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
