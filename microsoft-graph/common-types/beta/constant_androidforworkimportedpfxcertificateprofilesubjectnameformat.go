package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat string

const (
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonName                  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonName"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_Custom                      AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidForWorkImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkImportedPFXCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkImportedPFXCertificateProfileSubjectNameFormat(input string) (*AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
