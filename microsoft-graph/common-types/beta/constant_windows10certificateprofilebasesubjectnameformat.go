package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CertificateProfileBaseSubjectNameFormat string

const (
	Windows10CertificateProfileBaseSubjectNameFormat_CommonName                  Windows10CertificateProfileBaseSubjectNameFormat = "commonName"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     Windows10CertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId Windows10CertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           Windows10CertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            Windows10CertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  Windows10CertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    Windows10CertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	Windows10CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    Windows10CertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	Windows10CertificateProfileBaseSubjectNameFormat_Custom                      Windows10CertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForWindows10CertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonName),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(Windows10CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(Windows10CertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *Windows10CertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10CertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10CertificateProfileBaseSubjectNameFormat(input string) (*Windows10CertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]Windows10CertificateProfileBaseSubjectNameFormat{
		"commonname":                  Windows10CertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows10CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows10CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      Windows10CertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10CertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
