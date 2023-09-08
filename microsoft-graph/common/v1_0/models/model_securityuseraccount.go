package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserAccount struct {
	AccountName       *string `json:"accountName,omitempty"`
	AzureAdUserId     *string `json:"azureAdUserId,omitempty"`
	DisplayName       *string `json:"displayName,omitempty"`
	DomainName        *string `json:"domainName,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
	UserSid           *string `json:"userSid,omitempty"`
}
