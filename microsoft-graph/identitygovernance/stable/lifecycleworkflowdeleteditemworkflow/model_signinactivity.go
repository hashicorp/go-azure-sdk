package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInActivity struct {
	LastNonInteractiveSignInDateTime  *string `json:"lastNonInteractiveSignInDateTime,omitempty"`
	LastNonInteractiveSignInRequestId *string `json:"lastNonInteractiveSignInRequestId,omitempty"`
	LastSignInDateTime                *string `json:"lastSignInDateTime,omitempty"`
	LastSignInRequestId               *string `json:"lastSignInRequestId,omitempty"`
	ODataType                         *string `json:"@odata.type,omitempty"`
}
