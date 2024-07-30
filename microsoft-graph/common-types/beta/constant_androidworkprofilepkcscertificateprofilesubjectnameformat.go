package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat string

const (
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonName                  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonName"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_Custom                      AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidWorkProfilePkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfilePkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfilePkcsCertificateProfileSubjectNameFormat(input string) (*AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfilePkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
