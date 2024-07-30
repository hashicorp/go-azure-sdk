package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePatchContentCommand struct {
	Action    *OnenotePatchContentCommandAction   `json:"action,omitempty"`
	Content   *string                             `json:"content,omitempty"`
	ODataType *string                             `json:"@odata.type,omitempty"`
	Position  *OnenotePatchContentCommandPosition `json:"position,omitempty"`
	Target    *string                             `json:"target,omitempty"`
}
