package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatus struct {
	Description         *string                `json:"description,omitempty"`
	Details             *PrintJobStatusDetails `json:"details,omitempty"`
	IsAcquiredByPrinter *bool                  `json:"isAcquiredByPrinter,omitempty"`
	ODataType           *string                `json:"@odata.type,omitempty"`
	State               *PrintJobStatusState   `json:"state,omitempty"`
}
