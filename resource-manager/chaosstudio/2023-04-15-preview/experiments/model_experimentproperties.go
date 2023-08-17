package experiments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExperimentProperties struct {
	Selectors       []Selector `json:"selectors"`
	StartOnCreation *bool      `json:"startOnCreation,omitempty"`
	Steps           []Step     `json:"steps"`
}
