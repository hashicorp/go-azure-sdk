package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSScepCertificateProfileCertificateValidityPeriodScale string

const (
	MacOSScepCertificateProfileCertificateValidityPeriodScale_Days   MacOSScepCertificateProfileCertificateValidityPeriodScale = "days"
	MacOSScepCertificateProfileCertificateValidityPeriodScale_Months MacOSScepCertificateProfileCertificateValidityPeriodScale = "months"
	MacOSScepCertificateProfileCertificateValidityPeriodScale_Years  MacOSScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForMacOSScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(MacOSScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(MacOSScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(MacOSScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *MacOSScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSScepCertificateProfileCertificateValidityPeriodScale(input string) (*MacOSScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]MacOSScepCertificateProfileCertificateValidityPeriodScale{
		"days":   MacOSScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": MacOSScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  MacOSScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
