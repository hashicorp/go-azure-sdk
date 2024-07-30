package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat string

const (
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonName                  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonName"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_Custom                      AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileSubjectNameFormat(input string) (*AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
