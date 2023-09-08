package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SelfServiceSignUpAuthenticationFlowConfiguration struct {
	IsEnabled *bool   `json:"isEnabled,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
