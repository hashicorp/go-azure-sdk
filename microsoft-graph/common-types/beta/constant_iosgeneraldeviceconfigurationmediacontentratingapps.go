package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationMediaContentRatingApps string

const (
	IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove12 IosGeneralDeviceConfigurationMediaContentRatingApps = "agesAbove12"
	IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove17 IosGeneralDeviceConfigurationMediaContentRatingApps = "agesAbove17"
	IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove4  IosGeneralDeviceConfigurationMediaContentRatingApps = "agesAbove4"
	IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove9  IosGeneralDeviceConfigurationMediaContentRatingApps = "agesAbove9"
	IosGeneralDeviceConfigurationMediaContentRatingApps_AllAllowed  IosGeneralDeviceConfigurationMediaContentRatingApps = "allAllowed"
	IosGeneralDeviceConfigurationMediaContentRatingApps_AllBlocked  IosGeneralDeviceConfigurationMediaContentRatingApps = "allBlocked"
)

func PossibleValuesForIosGeneralDeviceConfigurationMediaContentRatingApps() []string {
	return []string{
		string(IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove12),
		string(IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove17),
		string(IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove4),
		string(IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove9),
		string(IosGeneralDeviceConfigurationMediaContentRatingApps_AllAllowed),
		string(IosGeneralDeviceConfigurationMediaContentRatingApps_AllBlocked),
	}
}

func (s *IosGeneralDeviceConfigurationMediaContentRatingApps) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosGeneralDeviceConfigurationMediaContentRatingApps(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosGeneralDeviceConfigurationMediaContentRatingApps(input string) (*IosGeneralDeviceConfigurationMediaContentRatingApps, error) {
	vals := map[string]IosGeneralDeviceConfigurationMediaContentRatingApps{
		"agesabove12": IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove12,
		"agesabove17": IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove17,
		"agesabove4":  IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove4,
		"agesabove9":  IosGeneralDeviceConfigurationMediaContentRatingApps_AgesAbove9,
		"allallowed":  IosGeneralDeviceConfigurationMediaContentRatingApps_AllAllowed,
		"allblocked":  IosGeneralDeviceConfigurationMediaContentRatingApps_AllBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationMediaContentRatingApps(input)
	return &out, nil
}
