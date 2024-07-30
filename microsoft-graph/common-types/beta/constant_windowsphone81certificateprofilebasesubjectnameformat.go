package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81CertificateProfileBaseSubjectNameFormat string

const (
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonName                  WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonName"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    WindowsPhone81CertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	WindowsPhone81CertificateProfileBaseSubjectNameFormat_Custom                      WindowsPhone81CertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForWindowsPhone81CertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonName),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(WindowsPhone81CertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *WindowsPhone81CertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81CertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81CertificateProfileBaseSubjectNameFormat(input string) (*WindowsPhone81CertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]WindowsPhone81CertificateProfileBaseSubjectNameFormat{
		"commonname":                  WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsPhone81CertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      WindowsPhone81CertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81CertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
