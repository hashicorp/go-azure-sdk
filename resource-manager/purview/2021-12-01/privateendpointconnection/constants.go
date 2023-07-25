package privateendpointconnection

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Status string

const (
	StatusApproved     Status = "Approved"
	StatusDisconnected Status = "Disconnected"
	StatusPending      Status = "Pending"
	StatusRejected     Status = "Rejected"
	StatusUnknown      Status = "Unknown"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusApproved),
		string(StatusDisconnected),
		string(StatusPending),
		string(StatusRejected),
		string(StatusUnknown),
	}
}

func (s *Status) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"approved":     StatusApproved,
		"disconnected": StatusDisconnected,
		"pending":      StatusPending,
		"rejected":     StatusRejected,
		"unknown":      StatusUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}
