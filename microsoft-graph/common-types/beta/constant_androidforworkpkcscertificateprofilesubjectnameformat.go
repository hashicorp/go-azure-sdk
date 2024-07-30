package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkPkcsCertificateProfileSubjectNameFormat string

const (
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonName                  AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonName"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidForWorkPkcsCertificateProfileSubjectNameFormat_Custom                      AndroidForWorkPkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidForWorkPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidForWorkPkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidForWorkPkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkPkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkPkcsCertificateProfileSubjectNameFormat(input string) (*AndroidForWorkPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidForWorkPkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
