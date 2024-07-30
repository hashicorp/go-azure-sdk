package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCertificateProfileBaseSubjectNameFormat string

const (
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonName                  AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonName"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId     AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsAadDeviceId"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsDurableDeviceId"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail           AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsEmail"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI            AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsIMEI"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId  AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber    AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameAsSerialNumber"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail    AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "commonNameIncludingEmail"
	AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_Custom                      AospDeviceOwnerCertificateProfileBaseSubjectNameFormat = "custom"
)

func PossibleValuesForAospDeviceOwnerCertificateProfileBaseSubjectNameFormat() []string {
	return []string{
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonName),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail),
		string(AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_Custom),
	}
}

func (s *AospDeviceOwnerCertificateProfileBaseSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerCertificateProfileBaseSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerCertificateProfileBaseSubjectNameFormat(input string) (*AospDeviceOwnerCertificateProfileBaseSubjectNameFormat, error) {
	vals := map[string]AospDeviceOwnerCertificateProfileBaseSubjectNameFormat{
		"commonname":                  AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AospDeviceOwnerCertificateProfileBaseSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCertificateProfileBaseSubjectNameFormat(input)
	return &out, nil
}
