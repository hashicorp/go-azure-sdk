package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCertificateProfileBaseSubjectNameFormat string

const (
	IosCertificateProfileBaseSubjectNameFormatcommonName               IosCertificateProfileBaseSubjectNameFormat = "CommonName"
	IosCertificateProfileBaseSubjectNameFormatcommonNameAsEmail        IosCertificateProfileBaseSubjectNameFormat = "CommonNameAsEmail"
	IosCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI         IosCertificateProfileBaseSubjectNameFormat = "CommonNameAsIMEI"
	IosCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber IosCertificateProfileBaseSubjectNameFormat = "CommonNameAsSerialNumber"
	IosCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail IosCertificateProfileBaseSubjectNameFormat = "CommonNameIncludingEmail"
	IosCertificateProfileBaseSubjectNameFormatcustom                   IosCertificateProfileBaseSubjectNameFormat = "Custom"
)

func PossibleValuesForIosCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(IosCertificateProfileBaseSubjectNameFormatcommonName),
		string(IosCertificateProfileBaseSubjectNameFormatcommonNameAsEmail),
		string(IosCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI),
		string(IosCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber),
		string(IosCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail),
		string(IosCertificateProfileBaseSubjectNameFormatcustom),
	}
}

func parseIosCertificateProfileBaseSubjectNameFormat(input string) (*IosCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]IosCertificateProfileBaseSubjectNameFormat{
		"commonname":               IosCertificateProfileBaseSubjectNameFormatcommonName,
		"commonnameasemail":        IosCertificateProfileBaseSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         IosCertificateProfileBaseSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": IosCertificateProfileBaseSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": IosCertificateProfileBaseSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   IosCertificateProfileBaseSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
