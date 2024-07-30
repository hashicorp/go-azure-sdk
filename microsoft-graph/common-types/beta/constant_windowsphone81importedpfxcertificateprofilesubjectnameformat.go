package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat string

const (
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonName                  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonName"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail           WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI            WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_Custom                      WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonName),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat(input string) (*WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
