package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInStatus struct {
	AdditionalDetails *string `json:"additionalDetails,omitempty"`
	ErrorCode         *int64  `json:"errorCode,omitempty"`
	FailureReason     *string `json:"failureReason,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
}
