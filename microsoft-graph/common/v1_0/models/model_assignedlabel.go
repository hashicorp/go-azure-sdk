package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedLabel struct {
	DisplayName *string `json:"displayName,omitempty"`
	LabelId     *string `json:"labelId,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
