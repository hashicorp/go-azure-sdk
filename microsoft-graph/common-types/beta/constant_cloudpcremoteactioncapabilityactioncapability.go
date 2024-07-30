package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionCapabilityActionCapability string

const (
	CloudPCRemoteActionCapabilityActionCapability_Disabled CloudPCRemoteActionCapabilityActionCapability = "disabled"
	CloudPCRemoteActionCapabilityActionCapability_Enabled  CloudPCRemoteActionCapabilityActionCapability = "enabled"
)

func PossibleValuesForCloudPCRemoteActionCapabilityActionCapability() []string {
	return []string{
		string(CloudPCRemoteActionCapabilityActionCapability_Disabled),
		string(CloudPCRemoteActionCapabilityActionCapability_Enabled),
	}
}

func (s *CloudPCRemoteActionCapabilityActionCapability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRemoteActionCapabilityActionCapability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRemoteActionCapabilityActionCapability(input string) (*CloudPCRemoteActionCapabilityActionCapability, error) {
	vals := map[string]CloudPCRemoteActionCapabilityActionCapability{
		"disabled": CloudPCRemoteActionCapabilityActionCapability_Disabled,
		"enabled":  CloudPCRemoteActionCapabilityActionCapability_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionCapabilityActionCapability(input)
	return &out, nil
}
