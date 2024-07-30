package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentApprovalRoute struct {
	AppConsentRequests *[]AppConsentRequest `json:"appConsentRequests,omitempty"`
	Id                 *string              `json:"id,omitempty"`
	ODataType          *string              `json:"@odata.type,omitempty"`
}
