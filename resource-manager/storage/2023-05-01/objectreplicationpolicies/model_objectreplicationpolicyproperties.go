package objectreplicationpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectReplicationPolicyProperties struct {
	DestinationAccount string                         `json:"destinationAccount"`
	EnabledTime        *string                        `json:"enabledTime,omitempty"`
	PolicyId           *string                        `json:"policyId,omitempty"`
	Rules              *[]ObjectReplicationPolicyRule `json:"rules,omitempty"`
	SourceAccount      string                         `json:"sourceAccount"`
}
