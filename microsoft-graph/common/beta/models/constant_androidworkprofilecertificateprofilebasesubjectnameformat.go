package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCertificateProfileBaseSubjectNameFormat string

const (
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonName                  AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonName"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId     AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsEmail           AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI            AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber    AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail    AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcustom                      AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidWorkProfileCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonName),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseAndroidWorkProfileCertificateProfileBaseSubjectNameFormat(input string) (*AndroidWorkProfileCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidWorkProfileCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidWorkProfileCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
