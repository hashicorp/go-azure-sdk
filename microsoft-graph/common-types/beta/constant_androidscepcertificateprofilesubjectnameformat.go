package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidScepCertificateProfileSubjectNameFormat string

const (
	AndroidScepCertificateProfileSubjectNameFormat_CommonName                  AndroidScepCertificateProfileSubjectNameFormat = "commonName"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidScepCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidScepCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidScepCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidScepCertificateProfileSubjectNameFormat_Custom                      AndroidScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidScepCertificateProfileSubjectNameFormat(input string) (*AndroidScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
