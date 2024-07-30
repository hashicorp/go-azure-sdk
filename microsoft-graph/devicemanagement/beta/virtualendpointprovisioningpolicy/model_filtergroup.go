package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterGroup struct {
	Clauses   *[]FilterClause `json:"clauses,omitempty"`
	Name      *string         `json:"name,omitempty"`
	ODataType *string         `json:"@odata.type,omitempty"`
}
