package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventQuery struct {
	ODataType *string                      `json:"@odata.type,omitempty"`
	Query     *string                      `json:"query,omitempty"`
	QueryType *SecurityEventQueryQueryType `json:"queryType,omitempty"`
}
