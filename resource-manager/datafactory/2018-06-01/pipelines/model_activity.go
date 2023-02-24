package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Activity struct {
	DependsOn      *[]ActivityDependency `json:"dependsOn,omitempty"`
	Description    *string               `json:"description,omitempty"`
	Name           string                `json:"name"`
	Type           string                `json:"type"`
	UserProperties *[]UserProperty       `json:"userProperties,omitempty"`
}
