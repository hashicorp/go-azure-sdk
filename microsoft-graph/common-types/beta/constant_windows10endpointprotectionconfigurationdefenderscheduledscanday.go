package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScheduledScanDay string

const (
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Everyday        Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "everyday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Friday          Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "friday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Monday          Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "monday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_NoScheduledScan Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "noScheduledScan"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Saturday        Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "saturday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Sunday          Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "sunday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Thursday        Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "thursday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Tuesday         Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "tuesday"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_UserDefined     Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "userDefined"
	Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Wednesday       Windows10EndpointProtectionConfigurationDefenderScheduledScanDay = "wednesday"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScheduledScanDay() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Everyday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Friday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Monday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_NoScheduledScan),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Saturday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Sunday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Thursday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Tuesday),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_UserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Wednesday),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScheduledScanDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScheduledScanDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScheduledScanDay(input string) (*Windows10EndpointProtectionConfigurationDefenderScheduledScanDay, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScheduledScanDay{
		"everyday":        Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Everyday,
		"friday":          Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Friday,
		"monday":          Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Monday,
		"noscheduledscan": Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_NoScheduledScan,
		"saturday":        Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Saturday,
		"sunday":          Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Sunday,
		"thursday":        Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Thursday,
		"tuesday":         Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Tuesday,
		"userdefined":     Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_UserDefined,
		"wednesday":       Windows10EndpointProtectionConfigurationDefenderScheduledScanDay_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScheduledScanDay(input)
	return &out, nil
}
