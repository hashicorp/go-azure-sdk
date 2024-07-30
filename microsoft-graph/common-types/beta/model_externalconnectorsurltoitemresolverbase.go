package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsUrlToItemResolverBase struct {
	ODataType *string `json:"@odata.type,omitempty"`
	Priority  *int64  `json:"priority,omitempty"`
}
