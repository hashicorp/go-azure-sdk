package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat string

const (
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonName                  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonName"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_Custom                      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonName),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat(input string) (*AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat{
		"commonname":                  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerImportedPFXCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
