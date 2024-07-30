package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionRegionStatus string

const (
	CloudPCSupportedRegionRegionStatus_Available   CloudPCSupportedRegionRegionStatus = "available"
	CloudPCSupportedRegionRegionStatus_Restricted  CloudPCSupportedRegionRegionStatus = "restricted"
	CloudPCSupportedRegionRegionStatus_Unavailable CloudPCSupportedRegionRegionStatus = "unavailable"
)

func PossibleValuesForCloudPCSupportedRegionRegionStatus() []string {
	return []string{
		string(CloudPCSupportedRegionRegionStatus_Available),
		string(CloudPCSupportedRegionRegionStatus_Restricted),
		string(CloudPCSupportedRegionRegionStatus_Unavailable),
	}
}

func (s *CloudPCSupportedRegionRegionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSupportedRegionRegionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSupportedRegionRegionStatus(input string) (*CloudPCSupportedRegionRegionStatus, error) {
	vals := map[string]CloudPCSupportedRegionRegionStatus{
		"available":   CloudPCSupportedRegionRegionStatus_Available,
		"restricted":  CloudPCSupportedRegionRegionStatus_Restricted,
		"unavailable": CloudPCSupportedRegionRegionStatus_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionRegionStatus(input)
	return &out, nil
}
