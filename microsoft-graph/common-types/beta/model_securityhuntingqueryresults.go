package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHuntingQueryResults struct {
	ODataType *string                         `json:"@odata.type,omitempty"`
	Results   *[]SecurityHuntingRowResult     `json:"results,omitempty"`
	Schema    *[]SecuritySinglePropertySchema `json:"schema,omitempty"`
}
