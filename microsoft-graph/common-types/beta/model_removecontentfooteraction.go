package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoveContentFooterAction struct {
	ODataType      *string   `json:"@odata.type,omitempty"`
	UiElementNames *[]string `json:"uiElementNames,omitempty"`
}
