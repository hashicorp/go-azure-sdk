package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale string

const (
	AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Days   AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "days"
	AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Months AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "months"
	AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Years  AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale = "years"
)

func PossibleValuesForAospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale() []string {
	return []string{
		string(AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Days),
		string(AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Months),
		string(AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Years),
	}
}

func (s *AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input string) (*AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale, error) {
	vals := map[string]AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale{
		"days":   AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Days,
		"months": AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Months,
		"years":  AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale_Years,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerCertificateProfileBaseCertificateValidityPeriodScale(input)
	return &out, nil
}
