package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionResultActionState string

const (
	CloudPCRemoteActionResultActionStateactive       CloudPCRemoteActionResultActionState = "Active"
	CloudPCRemoteActionResultActionStatecanceled     CloudPCRemoteActionResultActionState = "Canceled"
	CloudPCRemoteActionResultActionStatedone         CloudPCRemoteActionResultActionState = "Done"
	CloudPCRemoteActionResultActionStatefailed       CloudPCRemoteActionResultActionState = "Failed"
	CloudPCRemoteActionResultActionStatenone         CloudPCRemoteActionResultActionState = "None"
	CloudPCRemoteActionResultActionStatenotSupported CloudPCRemoteActionResultActionState = "NotSupported"
	CloudPCRemoteActionResultActionStatepending      CloudPCRemoteActionResultActionState = "Pending"
)

func PossibleValuesForCloudPCRemoteActionResultActionState() []string {
	return []string{
		string(CloudPCRemoteActionResultActionStateactive),
		string(CloudPCRemoteActionResultActionStatecanceled),
		string(CloudPCRemoteActionResultActionStatedone),
		string(CloudPCRemoteActionResultActionStatefailed),
		string(CloudPCRemoteActionResultActionStatenone),
		string(CloudPCRemoteActionResultActionStatenotSupported),
		string(CloudPCRemoteActionResultActionStatepending),
	}
}

func parseCloudPCRemoteActionResultActionState(input string) (*CloudPCRemoteActionResultActionState, error) {
	vals := map[string]CloudPCRemoteActionResultActionState{
		"active":       CloudPCRemoteActionResultActionStateactive,
		"canceled":     CloudPCRemoteActionResultActionStatecanceled,
		"done":         CloudPCRemoteActionResultActionStatedone,
		"failed":       CloudPCRemoteActionResultActionStatefailed,
		"none":         CloudPCRemoteActionResultActionStatenone,
		"notsupported": CloudPCRemoteActionResultActionStatenotSupported,
		"pending":      CloudPCRemoteActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionResultActionState(input)
	return &out, nil
}
