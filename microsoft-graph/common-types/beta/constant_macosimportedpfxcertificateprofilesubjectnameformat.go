package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileSubjectNameFormat string

const (
	MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonName               MacOSImportedPFXCertificateProfileSubjectNameFormat = "commonName"
	MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail        MacOSImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI         MacOSImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber MacOSImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail MacOSImportedPFXCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	MacOSImportedPFXCertificateProfileSubjectNameFormat_Custom                   MacOSImportedPFXCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonName),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(MacOSImportedPFXCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *MacOSImportedPFXCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSImportedPFXCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSImportedPFXCertificateProfileSubjectNameFormat(input string) (*MacOSImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":               MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasemail":        MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":         MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasserialnumber": MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail": MacOSImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                   MacOSImportedPFXCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
