package fleetmanagednamespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementV1Toleration struct {
	Effect   *TaintEffect        `json:"effect,omitempty"`
	Key      *string             `json:"key,omitempty"`
	Operator *TolerationOperator `json:"operator,omitempty"`
	Value    *string             `json:"value,omitempty"`
}
