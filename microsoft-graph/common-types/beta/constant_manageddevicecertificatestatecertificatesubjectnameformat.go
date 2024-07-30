package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateCertificateSubjectNameFormat string

const (
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonName                  ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonName"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsAadDeviceId     ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameAsAadDeviceId"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsDurableDeviceId ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameAsDurableDeviceId"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsEmail           ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameAsEmail"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsIMEI            ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameAsIMEI"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsIntuneDeviceId  ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameAsIntuneDeviceId"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsSerialNumber    ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameAsSerialNumber"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameIncludingEmail    ManagedDeviceCertificateStateCertificateSubjectNameFormat = "commonNameIncludingEmail"
	ManagedDeviceCertificateStateCertificateSubjectNameFormat_Custom                      ManagedDeviceCertificateStateCertificateSubjectNameFormat = "custom"
)

func PossibleValuesForManagedDeviceCertificateStateCertificateSubjectNameFormat() []string {
	return []string{
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonName),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsAadDeviceId),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsDurableDeviceId),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsEmail),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsIMEI),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsIntuneDeviceId),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsSerialNumber),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameIncludingEmail),
		string(ManagedDeviceCertificateStateCertificateSubjectNameFormat_Custom),
	}
}

func (s *ManagedDeviceCertificateStateCertificateSubjectNameFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateCertificateSubjectNameFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateCertificateSubjectNameFormat(input string) (*ManagedDeviceCertificateStateCertificateSubjectNameFormat, error) {
	vals := map[string]ManagedDeviceCertificateStateCertificateSubjectNameFormat{
		"commonname":                  ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonName,
		"commonnameasaaddeviceid":     ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsAadDeviceId,
		"commonnameasdurabledeviceid": ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsDurableDeviceId,
		"commonnameasemail":           ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsEmail,
		"commonnameasimei":            ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsIMEI,
		"commonnameasintunedeviceid":  ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsIntuneDeviceId,
		"commonnameasserialnumber":    ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameAsSerialNumber,
		"commonnameincludingemail":    ManagedDeviceCertificateStateCertificateSubjectNameFormat_CommonNameIncludingEmail,
		"custom":                      ManagedDeviceCertificateStateCertificateSubjectNameFormat_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateCertificateSubjectNameFormat(input)
	return &out, nil
}
