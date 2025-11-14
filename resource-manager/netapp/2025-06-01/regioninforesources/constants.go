package regioninforesources

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegionStorageToNetworkProximity string

const (
	RegionStorageToNetworkProximityAcrossTTwo               RegionStorageToNetworkProximity = "AcrossT2"
	RegionStorageToNetworkProximityDefault                  RegionStorageToNetworkProximity = "Default"
	RegionStorageToNetworkProximityTOne                     RegionStorageToNetworkProximity = "T1"
	RegionStorageToNetworkProximityTOneAndAcrossTTwo        RegionStorageToNetworkProximity = "T1AndAcrossT2"
	RegionStorageToNetworkProximityTOneAndTTwo              RegionStorageToNetworkProximity = "T1AndT2"
	RegionStorageToNetworkProximityTOneAndTTwoAndAcrossTTwo RegionStorageToNetworkProximity = "T1AndT2AndAcrossT2"
	RegionStorageToNetworkProximityTTwo                     RegionStorageToNetworkProximity = "T2"
	RegionStorageToNetworkProximityTTwoAndAcrossTTwo        RegionStorageToNetworkProximity = "T2AndAcrossT2"
)

func PossibleValuesForRegionStorageToNetworkProximity() []string {
	return []string{
		string(RegionStorageToNetworkProximityAcrossTTwo),
		string(RegionStorageToNetworkProximityDefault),
		string(RegionStorageToNetworkProximityTOne),
		string(RegionStorageToNetworkProximityTOneAndAcrossTTwo),
		string(RegionStorageToNetworkProximityTOneAndTTwo),
		string(RegionStorageToNetworkProximityTOneAndTTwoAndAcrossTTwo),
		string(RegionStorageToNetworkProximityTTwo),
		string(RegionStorageToNetworkProximityTTwoAndAcrossTTwo),
	}
}

func (s *RegionStorageToNetworkProximity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRegionStorageToNetworkProximity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRegionStorageToNetworkProximity(input string) (*RegionStorageToNetworkProximity, error) {
	vals := map[string]RegionStorageToNetworkProximity{
		"acrosst2":           RegionStorageToNetworkProximityAcrossTTwo,
		"default":            RegionStorageToNetworkProximityDefault,
		"t1":                 RegionStorageToNetworkProximityTOne,
		"t1andacrosst2":      RegionStorageToNetworkProximityTOneAndAcrossTTwo,
		"t1andt2":            RegionStorageToNetworkProximityTOneAndTTwo,
		"t1andt2andacrosst2": RegionStorageToNetworkProximityTOneAndTTwoAndAcrossTTwo,
		"t2":                 RegionStorageToNetworkProximityTTwo,
		"t2andacrosst2":      RegionStorageToNetworkProximityTTwoAndAcrossTTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RegionStorageToNetworkProximity(input)
	return &out, nil
}
