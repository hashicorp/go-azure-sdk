package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Days   AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Months AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Years  AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidImportedPFXCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidImportedPFXCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidImportedPFXCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidImportedPFXCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
