package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinarStatus string

const (
	VirtualEventWebinarStatus_Canceled  VirtualEventWebinarStatus = "canceled"
	VirtualEventWebinarStatus_Draft     VirtualEventWebinarStatus = "draft"
	VirtualEventWebinarStatus_Published VirtualEventWebinarStatus = "published"
)

func PossibleValuesForVirtualEventWebinarStatus() []string {
	return []string{
		string(VirtualEventWebinarStatus_Canceled),
		string(VirtualEventWebinarStatus_Draft),
		string(VirtualEventWebinarStatus_Published),
	}
}

func (s *VirtualEventWebinarStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventWebinarStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventWebinarStatus(input string) (*VirtualEventWebinarStatus, error) {
	vals := map[string]VirtualEventWebinarStatus{
		"canceled":  VirtualEventWebinarStatus_Canceled,
		"draft":     VirtualEventWebinarStatus_Draft,
		"published": VirtualEventWebinarStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventWebinarStatus(input)
	return &out, nil
}
