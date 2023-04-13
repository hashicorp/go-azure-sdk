package operations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationDetail struct {
	Display      *OperationDisplay    `json:"display,omitempty"`
	IsDataAction *bool                `json:"isDataAction,omitempty"`
	Name         *string              `json:"name,omitempty"`
	Origin       *string              `json:"origin,omitempty"`
	Properties   *OperationProperties `json:"properties,omitempty"`
}
