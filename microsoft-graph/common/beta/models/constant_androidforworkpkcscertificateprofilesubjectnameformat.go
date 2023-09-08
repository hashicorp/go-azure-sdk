package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkPkcsCertificateProfileSubjectNameFormat string

const (
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonName                  AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonName"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormatcustom                      AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidForWorkPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonName),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidForWorkPkcsCertificateProfileSubjectNameFormat(input string) (*AndroidForWorkPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidForWorkPkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
