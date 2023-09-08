package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat string

const (
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonName                  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonName"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcustom                      AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonName),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input string) (*AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
