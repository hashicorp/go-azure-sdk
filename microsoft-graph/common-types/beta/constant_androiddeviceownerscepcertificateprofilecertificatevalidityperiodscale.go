package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale string

const (
	AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Days   AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "days"
	AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Months AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "months"
	AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Years  AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input string) (*AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
