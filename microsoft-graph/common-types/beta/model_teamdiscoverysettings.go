package beta

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamDiscoverySettings struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// If set to true, the team is visible via search and suggestions from the Teams client.
	ShowInTeamsSearchAndSuggestions nullable.Type[bool] `json:"showInTeamsSearchAndSuggestions,omitempty"`
}
