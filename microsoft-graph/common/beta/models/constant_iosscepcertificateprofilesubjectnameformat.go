package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileSubjectNameFormat string

const (
	IosScepCertificateProfileSubjectNameFormatcommonName               IosScepCertificateProfileSubjectNameFormat = "CommonName"
	IosScepCertificateProfileSubjectNameFormatcommonNameAsEmail        IosScepCertificateProfileSubjectNameFormat = "CommonNameAsEmail"
	IosScepCertificateProfileSubjectNameFormatcommonNameAsIMEI         IosScepCertificateProfileSubjectNameFormat = "CommonNameAsIMEI"
	IosScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber IosScepCertificateProfileSubjectNameFormat = "CommonNameAsSerialNumber"
	IosScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail IosScepCertificateProfileSubjectNameFormat = "CommonNameIncludingEmail"
	IosScepCertificateProfileSubjectNameFormatcustom                   IosScepCertificateProfileSubjectNameFormat = "Custom"
)

func PossibleValuesForIosScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(IosScepCertificateProfileSubjectNameFormatcommonName),
		string(IosScepCertificateProfileSubjectNameFormatcommonNameAsEmail),
		string(IosScepCertificateProfileSubjectNameFormatcommonNameAsIMEI),
		string(IosScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber),
		string(IosScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail),
		string(IosScepCertificateProfileSubjectNameFormatcustom),
	}
}

func parseIosScepCertificateProfileSubjectNameFormat(input string) (*IosScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]IosScepCertificateProfileSubjectNameFormat{
		"commonname":               IosScepCertificateProfileSubjectNameFormatcommonName,
		"commonnameasemail":        IosScepCertificateProfileSubjectNameFormatcommonNameAsEmail,
		"commonnameasimei":         IosScepCertificateProfileSubjectNameFormatcommonNameAsIMEI,
		"commonnameasserialnumber": IosScepCertificateProfileSubjectNameFormatcommonNameAsSerialNumber,
		"commonnameincludingemail": IosScepCertificateProfileSubjectNameFormatcommonNameIncludingEmail,
		"custom":                   IosScepCertificateProfileSubjectNameFormatcustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
