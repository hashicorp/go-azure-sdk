package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinarStatus string

const (
	VirtualEventWebinarStatuscanceled  VirtualEventWebinarStatus = "Canceled"
	VirtualEventWebinarStatusdraft     VirtualEventWebinarStatus = "Draft"
	VirtualEventWebinarStatuspublished VirtualEventWebinarStatus = "Published"
)

func PossibleValuesForVirtualEventWebinarStatus() []string {
	return []string{
		string(VirtualEventWebinarStatuscanceled),
		string(VirtualEventWebinarStatusdraft),
		string(VirtualEventWebinarStatuspublished),
	}
}

func parseVirtualEventWebinarStatus(input string) (*VirtualEventWebinarStatus, error) {
	vals := map[string]VirtualEventWebinarStatus{
		"canceled":  VirtualEventWebinarStatuscanceled,
		"draft":     VirtualEventWebinarStatusdraft,
		"published": VirtualEventWebinarStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventWebinarStatus(input)
	return &out, nil
}
