package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81SCEPCertificateProfileSubjectNameFormat string

const (
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonName                  Windows81SCEPCertificateProfileSubjectNameFormat = "commonName"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsEmail           Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIMEI            Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    Windows81SCEPCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	Windows81SCEPCertificateProfileSubjectNameFormat_Custom                      Windows81SCEPCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForWindows81SCEPCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonName),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(Windows81SCEPCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *Windows81SCEPCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81SCEPCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81SCEPCertificateProfileSubjectNameFormat(input string) (*Windows81SCEPCertificateProfileSubjectNameFormat, error) {
	vals := map[string]Windows81SCEPCertificateProfileSubjectNameFormat{
		"commonname":                  Windows81SCEPCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    Windows81SCEPCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      Windows81SCEPCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81SCEPCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
