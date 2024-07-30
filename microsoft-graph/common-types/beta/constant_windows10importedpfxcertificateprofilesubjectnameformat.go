package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10ImportedPFXCertificateProfileSubjectNameFormat string

const (
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonName                  Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonName"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail           Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI            Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    Windows10ImportedPFXCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	Windows10ImportedPFXCertificateProfileSubjectNameFormat_Custom                      Windows10ImportedPFXCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForWindows10ImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonName),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(Windows10ImportedPFXCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *Windows10ImportedPFXCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10ImportedPFXCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10ImportedPFXCertificateProfileSubjectNameFormat(input string) (*Windows10ImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]Windows10ImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows10ImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      Windows10ImportedPFXCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10ImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
