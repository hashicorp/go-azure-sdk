package flux

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectStatusConditionDefinition struct {
	LastTransitionTime *string `json:"lastTransitionTime,omitempty"`
	Message            *string `json:"message,omitempty"`
	Reason             *string `json:"reason,omitempty"`
	Status             *string `json:"status,omitempty"`
	Type               *string `json:"type,omitempty"`
}
