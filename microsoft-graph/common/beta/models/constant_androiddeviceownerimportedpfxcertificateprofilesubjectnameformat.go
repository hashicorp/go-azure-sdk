package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonName                  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonName"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcustom                      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonName),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
