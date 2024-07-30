package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale string

const (
	WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Days   WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale = "days"
	WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Months WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale = "months"
	WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Years  WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForWindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Days),
		string(WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Months),
		string(WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale(input string) (*WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale{
		"days":   WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Days,
		"months": WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81SCEPCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
