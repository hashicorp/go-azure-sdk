package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextColumn struct {
	AllowMultipleLines          *bool   `json:"allowMultipleLines,omitempty"`
	AppendChangesToExistingText *bool   `json:"appendChangesToExistingText,omitempty"`
	LinesForEditing             *int64  `json:"linesForEditing,omitempty"`
	MaxLength                   *int64  `json:"maxLength,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
	TextType                    *string `json:"textType,omitempty"`
}
