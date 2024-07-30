package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LookupColumn struct {
	AllowMultipleValues   *bool   `json:"allowMultipleValues,omitempty"`
	AllowUnlimitedLength  *bool   `json:"allowUnlimitedLength,omitempty"`
	ColumnName            *string `json:"columnName,omitempty"`
	ListId                *string `json:"listId,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	PrimaryLookupColumnId *string `json:"primaryLookupColumnId,omitempty"`
}
