package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileSubjectNameFormat string

const (
	IosPkcsCertificateProfileSubjectNameFormatcommonName               IosPkcsCertificateProfileSubjectNameFormat = "CommonName"
	IosPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail        IosPkcsCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	IosPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI         IosPkcsCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	IosPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber IosPkcsCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	IosPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail IosPkcsCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	IosPkcsCertificateProfileSubjectNameFormatcustom                   IosPkcsCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForIosPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(IosPkcsCertificateProfileSubjectNameFormatcommonName),
		string(IosPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(IosPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(IosPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(IosPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(IosPkcsCertificateProfileSubjectNameFormatcustom),
	}
}

func parseIosPkcsCertificateProfileSubjectNameFormat(input string) (*IosPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]IosPkcsCertificateProfileSubjectNameFormat{
		"commonname":               IosPkcsCertificateProfileSubjectNameFormatcommonName,
		"commonnameasemail":        IosPkcsCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         IosPkcsCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": IosPkcsCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": IosPkcsCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   IosPkcsCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
