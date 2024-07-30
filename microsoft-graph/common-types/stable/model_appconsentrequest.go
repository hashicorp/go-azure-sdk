package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequest struct {
	AppDisplayName      *string                   `json:"appDisplayName,omitempty"`
	AppId               *string                   `json:"appId,omitempty"`
	Id                  *string                   `json:"id,omitempty"`
	ODataType           *string                   `json:"@odata.type,omitempty"`
	PendingScopes       *[]AppConsentRequestScope `json:"pendingScopes,omitempty"`
	UserConsentRequests *[]UserConsentRequest     `json:"userConsentRequests,omitempty"`
}
