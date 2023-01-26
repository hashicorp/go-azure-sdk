package application

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationResource struct {
	Etag       *string                            `json:"etag,omitempty"`
	Id         *string                            `json:"id,omitempty"`
	Identity   *identity.SystemAndUserAssignedMap `json:"identity,omitempty"`
	Location   *string                            `json:"location,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Properties *ApplicationResourceProperties     `json:"properties,omitempty"`
	SystemData *SystemData                        `json:"systemData,omitempty"`
	Tags       *map[string]string                 `json:"tags,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}
