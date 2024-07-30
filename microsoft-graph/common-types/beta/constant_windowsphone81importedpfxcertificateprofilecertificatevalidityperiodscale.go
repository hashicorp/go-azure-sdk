package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Days   WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale = "days"
	WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Months WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale = "months"
	WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Years  WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Days),
		string(WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Months),
		string(WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Days,
		"months": WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81ImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
