package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Days   MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale = "days"
	MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Months MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale = "months"
	MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Years  MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForMacOSImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Days),
		string(MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Months),
		string(MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSImportedPFXCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Days,
		"months": MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
