package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileSubjectNameFormat string

const (
	MacOSScepCertificateProfileSubjectNameFormat_CommonName               MacOSScepCertificateProfileSubjectNameFormat = "commonName"
	MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsEmail        MacOSScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI         MacOSScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber MacOSScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	MacOSScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail MacOSScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	MacOSScepCertificateProfileSubjectNameFormat_Custom                   MacOSScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForMacOSScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(MacOSScepCertificateProfileSubjectNameFormat_CommonName),
		string(MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(MacOSScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(MacOSScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *MacOSScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileSubjectNameFormat(input string) (*MacOSScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]MacOSScepCertificateProfileSubjectNameFormat{
		"commonname":               MacOSScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasemail":        MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": MacOSScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   MacOSScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
