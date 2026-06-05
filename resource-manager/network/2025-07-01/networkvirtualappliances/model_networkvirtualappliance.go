package networkvirtualappliances

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/identity"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkVirtualAppliance struct {
	Etag       *string                                  `json:"etag,omitempty"`
	Id         *string                                  `json:"id,omitempty"`
	Identity   *identity.SystemAndUserAssignedMap       `json:"identity,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Properties *NetworkVirtualAppliancePropertiesFormat `json:"properties,omitempty"`
	SystemData *systemdata.SystemData                   `json:"systemData,omitempty"`
	Type       *string                                  `json:"type,omitempty"`
}
