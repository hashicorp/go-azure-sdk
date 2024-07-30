package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCServicePlanProvisioningType string

const (
	CloudPCServicePlanProvisioningType_Dedicated CloudPCServicePlanProvisioningType = "dedicated"
	CloudPCServicePlanProvisioningType_Shared    CloudPCServicePlanProvisioningType = "shared"
)

func PossibleValuesForCloudPCServicePlanProvisioningType() []string {
	return []string{
		string(CloudPCServicePlanProvisioningType_Dedicated),
		string(CloudPCServicePlanProvisioningType_Shared),
	}
}

func (s *CloudPCServicePlanProvisioningType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCServicePlanProvisioningType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCServicePlanProvisioningType(input string) (*CloudPCServicePlanProvisioningType, error) {
	vals := map[string]CloudPCServicePlanProvisioningType{
		"dedicated": CloudPCServicePlanProvisioningType_Dedicated,
		"shared":    CloudPCServicePlanProvisioningType_Shared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCServicePlanProvisioningType(input)
	return &out, nil
}
