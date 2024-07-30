package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileSubjectNameFormat string

const (
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonName                  AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonName"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsEmail           AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AospDeviceOwnerScepCertificateProfileSubjectNameFormat_Custom                      AospDeviceOwnerScepCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonName),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AospDeviceOwnerScepCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileSubjectNameFormat(input string) (*AospDeviceOwnerScepCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileSubjectNameFormat{
		"commonname":                  AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AospDeviceOwnerScepCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AospDeviceOwnerScepCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
