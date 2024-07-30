package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCertificateProfileBaseSubjectNameFormat string

const (
	AndroidCertificateProfileBaseSubjectNameFormat_CommonName                  AndroidCertificateProfileBaseSubjectNameFormat = "commonName"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     AndroidCertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId AndroidCertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           AndroidCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            AndroidCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidCertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    AndroidCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    AndroidCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	AndroidCertificateProfileBaseSubjectNameFormat_Custom                      AndroidCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonName),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *AndroidCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidCertificateProfileBaseSubjectNameFormat(input string) (*AndroidCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
