package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobStatus struct {
	AcquiredByPrinter          *bool                          `json:"acquiredByPrinter,omitempty"`
	Description                *string                        `json:"description,omitempty"`
	Details                    *PrintJobStatusDetails         `json:"details,omitempty"`
	IsAcquiredByPrinter        *bool                          `json:"isAcquiredByPrinter,omitempty"`
	ODataType                  *string                        `json:"@odata.type,omitempty"`
	ProcessingState            *PrintJobStatusProcessingState `json:"processingState,omitempty"`
	ProcessingStateDescription *string                        `json:"processingStateDescription,omitempty"`
	State                      *PrintJobStatusState           `json:"state,omitempty"`
}
