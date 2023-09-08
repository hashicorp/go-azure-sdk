package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileSubjectNameFormat string

const (
	MacOSPkcsCertificateProfileSubjectNameFormatcommonName               MacOSPkcsCertificateProfileSubjectNameFormat = "CommonName"
	MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail        MacOSPkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI         MacOSPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber MacOSPkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	MacOSPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail MacOSPkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	MacOSPkcsCertificateProfileSubjectNameFormatcustom                   MacOSPkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForMacOSPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(MacOSPkcsCertificateProfileSubjectNameFormatcommonName),
		string(MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(MacOSPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(MacOSPkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseMacOSPkcsCertificateProfileSubjectNameFormat(input string) (*MacOSPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]MacOSPkcsCertificateProfileSubjectNameFormat{
		"commonname":               MacOSPkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasemail":        MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": MacOSPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   MacOSPkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
