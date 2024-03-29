package applications

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Application struct {
	Id         *string               `json:"id,omitempty"`
	Identity   *Identity             `json:"identity,omitempty"`
	Kind       string                `json:"kind"`
	Location   *string               `json:"location,omitempty"`
	ManagedBy  *string               `json:"managedBy,omitempty"`
	Name       *string               `json:"name,omitempty"`
	Plan       *Plan                 `json:"plan,omitempty"`
	Properties ApplicationProperties `json:"properties"`
	Sku        *Sku                  `json:"sku,omitempty"`
	Tags       *map[string]string    `json:"tags,omitempty"`
	Type       *string               `json:"type,omitempty"`
}
