package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat string

const (
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonName                  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonName"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_Custom                      AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input string) (*AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
