package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCertificateProfileBaseSubjectNameFormat string

const (
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonName                  AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonName"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    AndroidForWorkCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	AndroidForWorkCertificateProfileBaseSubjectNameFormat_Custom                      AndroidForWorkCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidForWorkCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonName),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidForWorkCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *AndroidForWorkCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCertificateProfileBaseSubjectNameFormat(input string) (*AndroidForWorkCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidForWorkCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
