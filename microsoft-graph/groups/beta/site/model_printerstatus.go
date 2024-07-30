package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterStatus struct {
	Description                *string                              `json:"description,omitempty"`
	Details                    *PrinterStatusDetails                `json:"details,omitempty"`
	ODataType                  *string                              `json:"@odata.type,omitempty"`
	ProcessingState            *PrinterStatusProcessingState        `json:"processingState,omitempty"`
	ProcessingStateDescription *string                              `json:"processingStateDescription,omitempty"`
	ProcessingStateReasons     *PrinterStatusProcessingStateReasons `json:"processingStateReasons,omitempty"`
	State                      *PrinterStatusState                  `json:"state,omitempty"`
}
