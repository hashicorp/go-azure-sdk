package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningErrorInfo struct {
	AdditionalDetails *string                             `json:"additionalDetails,omitempty"`
	ErrorCategory     *ProvisioningErrorInfoErrorCategory `json:"errorCategory,omitempty"`
	ErrorCode         *string                             `json:"errorCode,omitempty"`
	ODataType         *string                             `json:"@odata.type,omitempty"`
	Reason            *string                             `json:"reason,omitempty"`
	RecommendedAction *string                             `json:"recommendedAction,omitempty"`
}
