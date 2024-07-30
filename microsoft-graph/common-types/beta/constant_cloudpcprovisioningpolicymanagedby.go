package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicyManagedBy string

const (
	CloudPCProvisioningPolicyManagedBy_DevBox     CloudPCProvisioningPolicyManagedBy = "devBox"
	CloudPCProvisioningPolicyManagedBy_RpaBox     CloudPCProvisioningPolicyManagedBy = "rpaBox"
	CloudPCProvisioningPolicyManagedBy_Windows365 CloudPCProvisioningPolicyManagedBy = "windows365"
)

func PossibleValuesForCloudPCProvisioningPolicyManagedBy() []string {
	return []string{
		string(CloudPCProvisioningPolicyManagedBy_DevBox),
		string(CloudPCProvisioningPolicyManagedBy_RpaBox),
		string(CloudPCProvisioningPolicyManagedBy_Windows365),
	}
}

func (s *CloudPCProvisioningPolicyManagedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCProvisioningPolicyManagedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCProvisioningPolicyManagedBy(input string) (*CloudPCProvisioningPolicyManagedBy, error) {
	vals := map[string]CloudPCProvisioningPolicyManagedBy{
		"devbox":     CloudPCProvisioningPolicyManagedBy_DevBox,
		"rpabox":     CloudPCProvisioningPolicyManagedBy_RpaBox,
		"windows365": CloudPCProvisioningPolicyManagedBy_Windows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCProvisioningPolicyManagedBy(input)
	return &out, nil
}
