package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyProvisioningType string

const (
	CloudPCProvisioningPolicyProvisioningType_Dedicated CloudPCProvisioningPolicyProvisioningType = "dedicated"
	CloudPCProvisioningPolicyProvisioningType_Shared    CloudPCProvisioningPolicyProvisioningType = "shared"
)

func PossibleValuesForCloudPCProvisioningPolicyProvisioningType() []string {
	return []string{
		string(CloudPCProvisioningPolicyProvisioningType_Dedicated),
		string(CloudPCProvisioningPolicyProvisioningType_Shared),
	}
}

func (s *CloudPCProvisioningPolicyProvisioningType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCProvisioningPolicyProvisioningType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCProvisioningPolicyProvisioningType(input string) (*CloudPCProvisioningPolicyProvisioningType, error) {
	vals := map[string]CloudPCProvisioningPolicyProvisioningType{
		"dedicated": CloudPCProvisioningPolicyProvisioningType_Dedicated,
		"shared":    CloudPCProvisioningPolicyProvisioningType_Shared,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningPolicyProvisioningType(input)
	return &out, nil
}
