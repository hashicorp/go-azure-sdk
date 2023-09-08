package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileSubjectNameFormat string

const (
	MacOSScepCertificateProfileSubjectNameFormatcommonName               MacOSScepCertificateProfileSubjectNameFormat = "CommonName"
	MacOSScepCertificateProfileSubjectNameFormatcommonNameAsEmail        MacOSScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	MacOSScepCertificateProfileSubjectNameFormatcommonNameAsIMEI         MacOSScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	MacOSScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber MacOSScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	MacOSScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail MacOSScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	MacOSScepCertificateProfileSubjectNameFormatcustom                   MacOSScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForMacOSScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(MacOSScepCertificateProfileSubjectNameFormatcommonName),
		string(MacOSScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(MacOSScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(MacOSScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(MacOSScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(MacOSScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseMacOSScepCertificateProfileSubjectNameFormat(input string) (*MacOSScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]MacOSScepCertificateProfileSubjectNameFormat{
		"commonname":               MacOSScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasemail":        MacOSScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         MacOSScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": MacOSScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   MacOSScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
