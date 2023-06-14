package managedinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstance struct {
	Id         *string                    `json:"id,omitempty"`
	Identity   *ResourceIdentity          `json:"identity,omitempty"`
	Location   string                     `json:"location"`
	Name       *string                    `json:"name,omitempty"`
	Properties *ManagedInstanceProperties `json:"properties,omitempty"`
	Sku        *Sku                       `json:"sku,omitempty"`
	Tags       *map[string]string         `json:"tags,omitempty"`
	Type       *string                    `json:"type,omitempty"`
}
