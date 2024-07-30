package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsItemIdResolver struct {
	ItemId       *string                         `json:"itemId,omitempty"`
	ODataType    *string                         `json:"@odata.type,omitempty"`
	Priority     *int64                          `json:"priority,omitempty"`
	UrlMatchInfo *ExternalConnectorsUrlMatchInfo `json:"urlMatchInfo,omitempty"`
}
