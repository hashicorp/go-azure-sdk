package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat string

const (
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonName                  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonName"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcustom                      AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonName),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidForWorkImportedPFXCertificateProfileSubjectNameFormat(input string) (*AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidForWorkImportedPFXCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
