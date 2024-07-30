package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10PkcsCertificateProfileSubjectNameFormat string

const (
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonName                  Windows10PkcsCertificateProfileSubjectNameFormat = "commonName"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail           Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI            Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    Windows10PkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	Windows10PkcsCertificateProfileSubjectNameFormat_Custom                      Windows10PkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForWindows10PkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonName),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(Windows10PkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *Windows10PkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10PkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10PkcsCertificateProfileSubjectNameFormat(input string) (*Windows10PkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]Windows10PkcsCertificateProfileSubjectNameFormat{
		"commonname":                  Windows10PkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows10PkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      Windows10PkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10PkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
