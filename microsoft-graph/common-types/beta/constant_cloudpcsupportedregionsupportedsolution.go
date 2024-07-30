package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionSupportedSolution string

const (
	CloudPCSupportedRegionSupportedSolution_DevBox     CloudPCSupportedRegionSupportedSolution = "devBox"
	CloudPCSupportedRegionSupportedSolution_RpaBox     CloudPCSupportedRegionSupportedSolution = "rpaBox"
	CloudPCSupportedRegionSupportedSolution_Windows365 CloudPCSupportedRegionSupportedSolution = "windows365"
)

func PossibleValuesForCloudPCSupportedRegionSupportedSolution() []string {
	return []string{
		string(CloudPCSupportedRegionSupportedSolution_DevBox),
		string(CloudPCSupportedRegionSupportedSolution_RpaBox),
		string(CloudPCSupportedRegionSupportedSolution_Windows365),
	}
}

func (s *CloudPCSupportedRegionSupportedSolution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSupportedRegionSupportedSolution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSupportedRegionSupportedSolution(input string) (*CloudPCSupportedRegionSupportedSolution, error) {
	vals := map[string]CloudPCSupportedRegionSupportedSolution{
		"devbox":     CloudPCSupportedRegionSupportedSolution_DevBox,
		"rpabox":     CloudPCSupportedRegionSupportedSolution_RpaBox,
		"windows365": CloudPCSupportedRegionSupportedSolution_Windows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionSupportedSolution(input)
	return &out, nil
}
