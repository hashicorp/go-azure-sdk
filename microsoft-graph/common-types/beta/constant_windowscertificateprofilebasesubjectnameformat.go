package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsCertificateProfileBaseSubjectNameFormat string

const (
	WindowsCertificateProfileBaseSubjectNameFormat_CommonName                  WindowsCertificateProfileBaseSubjectNameFormat = "commonName"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     WindowsCertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId WindowsCertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           WindowsCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            WindowsCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  WindowsCertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    WindowsCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	WindowsCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    WindowsCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	WindowsCertificateProfileBaseSubjectNameFormat_Custom                      WindowsCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForWindowsCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonName),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(WindowsCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(WindowsCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *WindowsCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsCertificateProfileBaseSubjectNameFormat(input string) (*WindowsCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]WindowsCertificateProfileBaseSubjectNameFormat{
		"commonname":                  WindowsCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    WindowsCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    WindowsCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      WindowsCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
