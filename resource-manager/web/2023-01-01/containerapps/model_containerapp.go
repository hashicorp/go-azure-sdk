package containerapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerApp struct {
	Id         *string                 `json:"id,omitempty"`
	Kind       *string                 `json:"kind,omitempty"`
	Location   string                  `json:"location"`
	Name       *string                 `json:"name,omitempty"`
	Properties *ContainerAppProperties `json:"properties,omitempty"`
	Tags       *map[string]string      `json:"tags,omitempty"`
	Type       *string                 `json:"type,omitempty"`
}
