package hcxenterprisesites

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HcxEnterpriseSiteStatus string

const (
	HcxEnterpriseSiteStatusAvailable   HcxEnterpriseSiteStatus = "Available"
	HcxEnterpriseSiteStatusConsumed    HcxEnterpriseSiteStatus = "Consumed"
	HcxEnterpriseSiteStatusDeactivated HcxEnterpriseSiteStatus = "Deactivated"
	HcxEnterpriseSiteStatusDeleted     HcxEnterpriseSiteStatus = "Deleted"
)

func PossibleValuesForHcxEnterpriseSiteStatus() []string {
	return []string{
		string(HcxEnterpriseSiteStatusAvailable),
		string(HcxEnterpriseSiteStatusConsumed),
		string(HcxEnterpriseSiteStatusDeactivated),
		string(HcxEnterpriseSiteStatusDeleted),
	}
}

func (s *HcxEnterpriseSiteStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHcxEnterpriseSiteStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHcxEnterpriseSiteStatus(input string) (*HcxEnterpriseSiteStatus, error) {
	vals := map[string]HcxEnterpriseSiteStatus{
		"available":   HcxEnterpriseSiteStatusAvailable,
		"consumed":    HcxEnterpriseSiteStatusConsumed,
		"deactivated": HcxEnterpriseSiteStatusDeactivated,
		"deleted":     HcxEnterpriseSiteStatusDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HcxEnterpriseSiteStatus(input)
	return &out, nil
}
