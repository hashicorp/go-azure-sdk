package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCertificateProfileBaseSubjectNameFormat string

const (
	IosCertificateProfileBaseSubjectNameFormat_CommonName               IosCertificateProfileBaseSubjectNameFormat = "commonName"
	IosCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail        IosCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	IosCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI         IosCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	IosCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber IosCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	IosCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail IosCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	IosCertificateProfileBaseSubjectNameFormat_Custom                   IosCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForIosCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(IosCertificateProfileBaseSubjectNameFormat_CommonName),
		string(IosCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(IosCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(IosCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(IosCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(IosCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *IosCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosCertificateProfileBaseSubjectNameFormat(input string) (*IosCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]IosCertificateProfileBaseSubjectNameFormat{
		"commonname":               IosCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasemail":        IosCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         IosCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": IosCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": IosCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   IosCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
