package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkScepCertificateProfileSubjectNameFormat string

const (
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonName                  AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonName"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidForWorkScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidForWorkScepCertificateProfileSubjectNameFormat_Custom                      AndroidForWorkScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidForWorkScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidForWorkScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidForWorkScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkScepCertificateProfileSubjectNameFormat(input string) (*AndroidForWorkScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkScepCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidForWorkScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
