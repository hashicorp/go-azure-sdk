package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat string

const (
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonName                  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonName"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_Custom                      AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonName),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat(input string) (*AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
