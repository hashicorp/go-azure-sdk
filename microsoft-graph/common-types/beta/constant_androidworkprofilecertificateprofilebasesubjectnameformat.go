package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCertificateProfileBaseSubjectNameFormat string

const (
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonName                  AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonName"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_Custom                      AndroidWorkProfileCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidWorkProfileCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonName),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *AndroidWorkProfileCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCertificateProfileBaseSubjectNameFormat(input string) (*AndroidWorkProfileCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidWorkProfileCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidWorkProfileCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
