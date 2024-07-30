package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale string

const (
	AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Days   AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "days"
	AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Months AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "months"
	AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Years  AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale() []string {
	return []string{
		string(AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Days),
		string(AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Months),
		string(AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Years),
	}
}

func (s *AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input string) (*AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale, error) {
	vals := map[string]AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale{
		"days":   AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Days,
		"months": AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Months,
		"years":  AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerScepCertificateProfileCertificateValidityPeriodScale(input)
	return &out, nil
}
