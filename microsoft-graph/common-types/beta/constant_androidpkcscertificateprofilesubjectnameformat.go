package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidPkcsCertificateProfileSubjectNameFormat string

const (
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonName                  AndroidPkcsCertificateProfileSubjectNameFormat = "commonName"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidPkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidPkcsCertificateProfileSubjectNameFormat_Custom                      AndroidPkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidPkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidPkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidPkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidPkcsCertificateProfileSubjectNameFormat(input string) (*AndroidPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidPkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidPkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
