package gates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRunGateTargetProperties struct {
	Group  *string `json:"group,omitempty"`
	Name   string  `json:"name"`
	Stage  *string `json:"stage,omitempty"`
	Timing Timing  `json:"timing"`
}
