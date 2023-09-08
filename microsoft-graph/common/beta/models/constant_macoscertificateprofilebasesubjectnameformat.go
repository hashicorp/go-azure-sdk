package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCertificateProfileBaseSubjectNameFormat string

const (
	MacOSCertificateProfileBaseSubjectNameFormatcommonName               MacOSCertificateProfileBaseSubjectNameFormat = "CommonName"
	MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsEmail        MacOSCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI         MacOSCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber MacOSCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	MacOSCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail MacOSCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	MacOSCertificateProfileBaseSubjectNameFormatcustom                   MacOSCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForMacOSCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(MacOSCertificateProfileBaseSubjectNameFormatcommonName),
		string(MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(MacOSCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(MacOSCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseMacOSCertificateProfileBaseSubjectNameFormat(input string) (*MacOSCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]MacOSCertificateProfileBaseSubjectNameFormat{
		"commonname":               MacOSCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasemail":        MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": MacOSCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   MacOSCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
