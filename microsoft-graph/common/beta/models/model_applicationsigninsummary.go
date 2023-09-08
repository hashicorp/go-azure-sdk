package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSignInSummary struct {
	AppDisplayName        *string `json:"appDisplayName,omitempty"`
	FailedSignInCount     *int64  `json:"failedSignInCount,omitempty"`
	Id                    *string `json:"id,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	SuccessfulSignInCount *int64  `json:"successfulSignInCount,omitempty"`
}
