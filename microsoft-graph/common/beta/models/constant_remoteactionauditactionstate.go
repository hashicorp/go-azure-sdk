package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteActionAuditActionState string

const (
	RemoteActionAuditActionStateactive       RemoteActionAuditActionState = "Active"
	RemoteActionAuditActionStatecanceled     RemoteActionAuditActionState = "Canceled"
	RemoteActionAuditActionStatedone         RemoteActionAuditActionState = "Done"
	RemoteActionAuditActionStatefailed       RemoteActionAuditActionState = "Failed"
	RemoteActionAuditActionStatenone         RemoteActionAuditActionState = "None"
	RemoteActionAuditActionStatenotSupported RemoteActionAuditActionState = "NotSupported"
	RemoteActionAuditActionStatepending      RemoteActionAuditActionState = "Pending"
)

func PossibleValuesForRemoteActionAuditActionState() []string {
	return []string{
		string(RemoteActionAuditActionStateactive),
		string(RemoteActionAuditActionStatecanceled),
		string(RemoteActionAuditActionStatedone),
		string(RemoteActionAuditActionStatefailed),
		string(RemoteActionAuditActionStatenone),
		string(RemoteActionAuditActionStatenotSupported),
		string(RemoteActionAuditActionStatepending),
	}
}

func parseRemoteActionAuditActionState(input string) (*RemoteActionAuditActionState, error) {
	vals := map[string]RemoteActionAuditActionState{
		"active":       RemoteActionAuditActionStateactive,
		"canceled":     RemoteActionAuditActionStatecanceled,
		"done":         RemoteActionAuditActionStatedone,
		"failed":       RemoteActionAuditActionStatefailed,
		"none":         RemoteActionAuditActionStatenone,
		"notsupported": RemoteActionAuditActionStatenotSupported,
		"pending":      RemoteActionAuditActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteActionAuditActionState(input)
	return &out, nil
}
