package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChecklistItem struct {
	CheckedDateTime *string `json:"checkedDateTime,omitempty"`
	CreatedDateTime *string `json:"createdDateTime,omitempty"`
	DisplayName     *string `json:"displayName,omitempty"`
	Id              *string `json:"id,omitempty"`
	IsChecked       *bool   `json:"isChecked,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
