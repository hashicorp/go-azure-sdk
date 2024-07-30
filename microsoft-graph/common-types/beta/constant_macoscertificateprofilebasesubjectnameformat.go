package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSCertificateProfileBaseSubjectNameFormat string

const (
	MacOSCertificateProfileBaseSubjectNameFormat_CommonName               MacOSCertificateProfileBaseSubjectNameFormat = "commonName"
	MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail        MacOSCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI         MacOSCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber MacOSCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	MacOSCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail MacOSCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	MacOSCertificateProfileBaseSubjectNameFormat_Custom                   MacOSCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForMacOSCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(MacOSCertificateProfileBaseSubjectNameFormat_CommonName),
		string(MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(MacOSCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(MacOSCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *MacOSCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSCertificateProfileBaseSubjectNameFormat(input string) (*MacOSCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]MacOSCertificateProfileBaseSubjectNameFormat{
		"commonname":               MacOSCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasemail":        MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": MacOSCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   MacOSCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
