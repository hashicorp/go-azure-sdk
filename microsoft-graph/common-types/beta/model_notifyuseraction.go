package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotifyUserAction struct {
	ActionLastModifiedDateTime *string   `json:"actionLastModifiedDateTime,omitempty"`
	EmailText                  *string   `json:"emailText,omitempty"`
	ODataType                  *string   `json:"@odata.type,omitempty"`
	PolicyTip                  *string   `json:"policyTip,omitempty"`
	Recipients                 *[]string `json:"recipients,omitempty"`
}
