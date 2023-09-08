package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPkcsCertificateProfileSubjectNameFormat string

const (
	AndroidPkcsCertificateProfileSubjectNameFormatcommonName                  AndroidPkcsCertificateProfileSubjectNameFormat = "CommonName"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidPkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidPkcsCertificateProfileSubjectNameFormatcustom                      AndroidPkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonName),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidPkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidPkcsCertificateProfileSubjectNameFormat(input string) (*AndroidPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidPkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidPkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
