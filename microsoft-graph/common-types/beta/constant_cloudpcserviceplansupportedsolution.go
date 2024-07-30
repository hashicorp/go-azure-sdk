package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanSupportedSolution string

const (
	CloudPCServicePlanSupportedSolution_DevBox     CloudPCServicePlanSupportedSolution = "devBox"
	CloudPCServicePlanSupportedSolution_RpaBox     CloudPCServicePlanSupportedSolution = "rpaBox"
	CloudPCServicePlanSupportedSolution_Windows365 CloudPCServicePlanSupportedSolution = "windows365"
)

func PossibleValuesForCloudPCServicePlanSupportedSolution() []string {
	return []string{
		string(CloudPCServicePlanSupportedSolution_DevBox),
		string(CloudPCServicePlanSupportedSolution_RpaBox),
		string(CloudPCServicePlanSupportedSolution_Windows365),
	}
}

func (s *CloudPCServicePlanSupportedSolution) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCServicePlanSupportedSolution(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCServicePlanSupportedSolution(input string) (*CloudPCServicePlanSupportedSolution, error) {
	vals := map[string]CloudPCServicePlanSupportedSolution{
		"devbox":     CloudPCServicePlanSupportedSolution_DevBox,
		"rpabox":     CloudPCServicePlanSupportedSolution_RpaBox,
		"windows365": CloudPCServicePlanSupportedSolution_Windows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanSupportedSolution(input)
	return &out, nil
}
