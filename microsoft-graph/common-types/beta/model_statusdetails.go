package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusDetails struct {
	AdditionalDetails *string                     `json:"additionalDetails,omitempty"`
	ErrorCategory     *StatusDetailsErrorCategory `json:"errorCategory,omitempty"`
	ErrorCode         *string                     `json:"errorCode,omitempty"`
	ODataType         *string                     `json:"@odata.type,omitempty"`
	Reason            *string                     `json:"reason,omitempty"`
	RecommendedAction *string                     `json:"recommendedAction,omitempty"`
	Status            *StatusDetailsStatus        `json:"status,omitempty"`
}
