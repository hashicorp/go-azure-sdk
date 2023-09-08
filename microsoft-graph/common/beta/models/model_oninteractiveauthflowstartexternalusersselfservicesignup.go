package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnInteractiveAuthFlowStartExternalUsersSelfServiceSignUp struct {
	IsSignUpAllowed *bool   `json:"isSignUpAllowed,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
