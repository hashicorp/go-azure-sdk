package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncompleteData struct {
	MissingDataBeforeDateTime *string `json:"missingDataBeforeDateTime,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
	WasThrottled              *bool   `json:"wasThrottled,omitempty"`
}
