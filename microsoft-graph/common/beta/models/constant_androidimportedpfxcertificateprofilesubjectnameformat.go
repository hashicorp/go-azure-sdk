package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileSubjectNameFormat string

const (
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonName                  AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonName"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidImportedPFXCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidImportedPFXCertificateProfileSubjectNameFormatcustom                      AndroidImportedPFXCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonName),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidImportedPFXCertificateProfileSubjectNameFormat(input string) (*AndroidImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidImportedPFXCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidImportedPFXCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
