package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat string

const (
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonName                  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonName"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId     AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsAadDeviceId"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsDurableDeviceId"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail           AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsEmail"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI            AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsIMEI"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsIntuneDeviceId"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameAsSerialNumber"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "commonNameIncludingEmail"
	AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_Custom                      AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat = "custom"
)

func PossibleValuesForAospDeviceOwnerPkcsCertificateProfileSubjectNameFormat() []string {
	return []string{
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonName),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail),
		string(AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_Custom),
	}
}

func (s *AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerPkcsCertificateProfileSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input string) (*AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat, error) {
	vals := map[string]AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat{
		"commonname":                  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerPkcsCertificateProfileSubjectNameFormat(input)
	return &out, nil
}
