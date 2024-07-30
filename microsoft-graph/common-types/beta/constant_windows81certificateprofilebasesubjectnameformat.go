package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81CertificateProfileBaseSubjectNameFormat string

const (
	Windows81CertificateProfileBaseSubjectNameFormat_CommonName                  Windows81CertificateProfileBaseSubjectNameFormat = "commonName"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     Windows81CertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId Windows81CertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           Windows81CertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            Windows81CertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  Windows81CertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    Windows81CertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	Windows81CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    Windows81CertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	Windows81CertificateProfileBaseSubjectNameFormat_Custom                      Windows81CertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForWindows81CertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonName),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(Windows81CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(Windows81CertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *Windows81CertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81CertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81CertificateProfileBaseSubjectNameFormat(input string) (*Windows81CertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]Windows81CertificateProfileBaseSubjectNameFormat{
		"commonname":                  Windows81CertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows81CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows81CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      Windows81CertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81CertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
