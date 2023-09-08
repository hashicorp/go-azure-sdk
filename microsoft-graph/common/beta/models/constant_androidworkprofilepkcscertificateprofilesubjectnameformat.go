package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat string

const (
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonName                  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonName"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId     AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameAsAadDeviceId"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameAsDurableDeviceId"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsEmail           AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI            AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameAsIntuneDeviceId"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcustom                      AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonName),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseAndroidWorkProfilePkcsCertificateProfileSubjectNameFormat(input string) (*AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasaaddeviceid":     AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":            AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                      AndroidWorkProfilePkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
