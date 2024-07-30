package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileSubjectNameFormat string

const (
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonName                  AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonName"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidImportedPFXCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidImportedPFXCertificateProfileSubjectNameFormat_Custom                      AndroidImportedPFXCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidImportedPFXCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidImportedPFXCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidImportedPFXCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidImportedPFXCertificateProfileSubjectNameFormat(input string) (*AndroidImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidImportedPFXCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
