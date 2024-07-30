package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInactiveUsersQueryScope struct {
	InactiveDuration *string `json:"inactiveDuration,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	Query            *string `json:"query,omitempty"`
	QueryRoot        *string `json:"queryRoot,omitempty"`
	QueryType        *string `json:"queryType,omitempty"`
}
