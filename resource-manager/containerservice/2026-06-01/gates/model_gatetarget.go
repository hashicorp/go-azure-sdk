package gates

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GateTarget struct {
	Id                  string                         `json:"id"`
	UpdateRunProperties *UpdateRunGateTargetProperties `json:"updateRunProperties,omitempty"`
}
