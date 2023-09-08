package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileSubjectNameFormat string

const (
	MacOSImportedPFXCertificateProfileSubjectNameFormatcommonName               MacOSImportedPFXCertificateProfileSubjectNameFormat = "CommonName"
	MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail        MacOSImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI         MacOSImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber MacOSImportedPFXCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail MacOSImportedPFXCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	MacOSImportedPFXCertificateProfileSubjectNameFormatcustom                   MacOSImportedPFXCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileSubjectNameFormatcommonName),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormatcustom),
	}
}

func parseMacOSImportedPFXCertificateProfileSubjectNameFormat(input string) (*MacOSImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":               MacOSImportedPFXCertificateProfileSubjectNameFormatcommonName,
		"commonnameasemail":        MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSImportedPFXCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   MacOSImportedPFXCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
