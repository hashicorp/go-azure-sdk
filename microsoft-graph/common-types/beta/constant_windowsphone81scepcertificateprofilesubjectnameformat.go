package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileSubjectNameFormat string

const (
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonName                  WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonName"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsEmail           WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIMEI            WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	WindowsPhone81SCEPCertificateProfileSubjectNameFormat_Custom                      WindowsPhone81SCEPCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonName),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(WindowsPhone81SCEPCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileSubjectNameFormat(input string) (*WindowsPhone81SCEPCertificateProfileSubjectNameFormat, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileSubjectNameFormat{
		"commonname":                  WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsPhone81SCEPCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      WindowsPhone81SCEPCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
