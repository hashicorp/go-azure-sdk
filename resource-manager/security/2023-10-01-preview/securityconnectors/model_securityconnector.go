package securityconnectors

import (
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityConnector struct {
	Etag       *string                      `json:"etag,omitempty"`
	Id         *string                      `json:"id,omitempty"`
	Kind       *string                      `json:"kind,omitempty"`
	Location   *string                      `json:"location,omitempty"`
	Name       *string                      `json:"name,omitempty"`
	Properties *SecurityConnectorProperties `json:"properties,omitempty"`
	SystemData *systemdata.SystemData       `json:"systemData,omitempty"`
	Tags       *map[string]string           `json:"tags,omitempty"`
	Type       *string                      `json:"type,omitempty"`
}
