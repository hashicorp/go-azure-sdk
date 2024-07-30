package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosPkcsCertificateProfileSubjectNameFormat string

const (
	IosPkcsCertificateProfileSubjectNameFormat_CommonName               IosPkcsCertificateProfileSubjectNameFormat = "commonName"
	IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail        IosPkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI         IosPkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber IosPkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	IosPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail IosPkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	IosPkcsCertificateProfileSubjectNameFormat_Custom                   IosPkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForIosPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(IosPkcsCertificateProfileSubjectNameFormat_CommonName),
		string(IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(IosPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(IosPkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *IosPkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosPkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosPkcsCertificateProfileSubjectNameFormat(input string) (*IosPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]IosPkcsCertificateProfileSubjectNameFormat{
		"commonname":               IosPkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasemail":        IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": IosPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": IosPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   IosPkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
