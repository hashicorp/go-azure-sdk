package experiments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Selector struct {
	Id      string            `json:"id"`
	Targets []TargetReference `json:"targets"`
	Type    SelectorType      `json:"type"`
}
