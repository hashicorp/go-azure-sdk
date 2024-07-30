package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosScepCertificateProfileSubjectNameFormat string

const (
	IosScepCertificateProfileSubjectNameFormat_CommonName               IosScepCertificateProfileSubjectNameFormat = "commonName"
	IosScepCertificateProfileSubjectNameFormat_CommonNameAsEmail        IosScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	IosScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI         IosScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	IosScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber IosScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	IosScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail IosScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	IosScepCertificateProfileSubjectNameFormat_Custom                   IosScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForIosScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(IosScepCertificateProfileSubjectNameFormat_CommonName),
		string(IosScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(IosScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(IosScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(IosScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(IosScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *IosScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosScepCertificateProfileSubjectNameFormat(input string) (*IosScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]IosScepCertificateProfileSubjectNameFormat{
		"commonname":               IosScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasemail":        IosScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         IosScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": IosScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": IosScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   IosScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
