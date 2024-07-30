package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionManagedBy string

const (
	CloudPCOnPremisesConnectionManagedBy_DevBox     CloudPCOnPremisesConnectionManagedBy = "devBox"
	CloudPCOnPremisesConnectionManagedBy_RpaBox     CloudPCOnPremisesConnectionManagedBy = "rpaBox"
	CloudPCOnPremisesConnectionManagedBy_Windows365 CloudPCOnPremisesConnectionManagedBy = "windows365"
)

func PossibleValuesForCloudPCOnPremisesConnectionManagedBy() []string {
	return []string{
		string(CloudPCOnPremisesConnectionManagedBy_DevBox),
		string(CloudPCOnPremisesConnectionManagedBy_RpaBox),
		string(CloudPCOnPremisesConnectionManagedBy_Windows365),
	}
}

func (s *CloudPCOnPremisesConnectionManagedBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOnPremisesConnectionManagedBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOnPremisesConnectionManagedBy(input string) (*CloudPCOnPremisesConnectionManagedBy, error) {
	vals := map[string]CloudPCOnPremisesConnectionManagedBy{
		"devbox":     CloudPCOnPremisesConnectionManagedBy_DevBox,
		"rpabox":     CloudPCOnPremisesConnectionManagedBy_RpaBox,
		"windows365": CloudPCOnPremisesConnectionManagedBy_Windows365,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOnPremisesConnectionManagedBy(input)
	return &out, nil
}
