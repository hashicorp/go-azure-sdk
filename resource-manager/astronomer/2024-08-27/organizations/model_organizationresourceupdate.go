package organizations

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationResourceUpdate struct {
	Identity   *identity.SystemAndUserAssignedMap    `json:"identity,omitempty"`
	Properties *OrganizationResourceUpdateProperties `json:"properties,omitempty"`
	Tags       *map[string]string                    `json:"tags,omitempty"`
}
