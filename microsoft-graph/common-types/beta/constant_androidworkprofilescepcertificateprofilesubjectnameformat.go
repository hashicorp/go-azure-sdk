package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileScepCertificateProfileSubjectNameFormat string

const (
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonName                  AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonName"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidWorkProfileScepCertificateProfileSubjectNameFormat_Custom                      AndroidWorkProfileScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidWorkProfileScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidWorkProfileScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidWorkProfileScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileScepCertificateProfileSubjectNameFormat(input string) (*AndroidWorkProfileScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidWorkProfileScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidWorkProfileScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidWorkProfileScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
