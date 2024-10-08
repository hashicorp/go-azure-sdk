package dataflowdebugsession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FlowletTypeProperties struct {
	Script          *string           `json:"script,omitempty"`
	ScriptLines     *[]string         `json:"scriptLines,omitempty"`
	Sinks           *[]DataFlowSink   `json:"sinks,omitempty"`
	Sources         *[]DataFlowSource `json:"sources,omitempty"`
	Transformations *[]Transformation `json:"transformations,omitempty"`
}
