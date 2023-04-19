package invitation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationStatus string

const (
	InvitationStatusAccepted  InvitationStatus = "Accepted"
	InvitationStatusPending   InvitationStatus = "Pending"
	InvitationStatusRejected  InvitationStatus = "Rejected"
	InvitationStatusWithdrawn InvitationStatus = "Withdrawn"
)

func PossibleValuesForInvitationStatus() []string {
	return []string{
		string(InvitationStatusAccepted),
		string(InvitationStatusPending),
		string(InvitationStatusRejected),
		string(InvitationStatusWithdrawn),
	}
}

func (s *InvitationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInvitationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInvitationStatus(input string) (*InvitationStatus, error) {
	vals := map[string]InvitationStatus{
		"accepted":  InvitationStatusAccepted,
		"pending":   InvitationStatusPending,
		"rejected":  InvitationStatusRejected,
		"withdrawn": InvitationStatusWithdrawn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InvitationStatus(input)
	return &out, nil
}
