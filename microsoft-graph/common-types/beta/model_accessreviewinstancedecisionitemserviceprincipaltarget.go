package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemServicePrincipalTarget struct {
	AppId                       *string `json:"appId,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
	ServicePrincipalDisplayName *string `json:"servicePrincipalDisplayName,omitempty"`
	ServicePrincipalId          *string `json:"servicePrincipalId,omitempty"`
}
