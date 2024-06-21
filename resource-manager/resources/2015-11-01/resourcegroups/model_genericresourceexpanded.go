package resourcegroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenericResourceExpanded struct {
	ChangedTime       *string            `json:"changedTime,omitempty"`
	CreatedTime       *string            `json:"createdTime,omitempty"`
	Id                *string            `json:"id,omitempty"`
	Location          string             `json:"location"`
	Name              *string            `json:"name,omitempty"`
	Plan              *Plan              `json:"plan,omitempty"`
	Properties        *interface{}       `json:"properties,omitempty"`
	ProvisioningState *string            `json:"provisioningState,omitempty"`
	Tags              *map[string]string `json:"tags,omitempty"`
	Type              *string            `json:"type,omitempty"`
}
