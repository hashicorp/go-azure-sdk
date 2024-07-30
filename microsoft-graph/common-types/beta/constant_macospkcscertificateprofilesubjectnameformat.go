package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPkcsCertificateProfileSubjectNameFormat string

const (
	MacOSPkcsCertificateProfileSubjectNameFormat_CommonName               MacOSPkcsCertificateProfileSubjectNameFormat = "commonName"
	MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail        MacOSPkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI         MacOSPkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber MacOSPkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail MacOSPkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	MacOSPkcsCertificateProfileSubjectNameFormat_Custom                   MacOSPkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForMacOSPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(MacOSPkcsCertificateProfileSubjectNameFormat_CommonName),
		string(MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(MacOSPkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *MacOSPkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPkcsCertificateProfileSubjectNameFormat(input string) (*MacOSPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]MacOSPkcsCertificateProfileSubjectNameFormat{
		"commonname":               MacOSPkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasemail":        MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   MacOSPkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
