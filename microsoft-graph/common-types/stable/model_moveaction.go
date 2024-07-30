package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MoveAction struct {
	From      *string `json:"from,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	To        *string `json:"to,omitempty"`
}
