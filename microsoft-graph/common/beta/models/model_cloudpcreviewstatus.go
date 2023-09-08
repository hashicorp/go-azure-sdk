package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReviewStatus struct {
	AzureStorageAccountId     *string                             `json:"azureStorageAccountId,omitempty"`
	AzureStorageAccountName   *string                             `json:"azureStorageAccountName,omitempty"`
	AzureStorageContainerName *string                             `json:"azureStorageContainerName,omitempty"`
	InReview                  *bool                               `json:"inReview,omitempty"`
	ODataType                 *string                             `json:"@odata.type,omitempty"`
	RestorePointDateTime      *string                             `json:"restorePointDateTime,omitempty"`
	ReviewStartDateTime       *string                             `json:"reviewStartDateTime,omitempty"`
	SubscriptionId            *string                             `json:"subscriptionId,omitempty"`
	SubscriptionName          *string                             `json:"subscriptionName,omitempty"`
	UserAccessLevel           *CloudPCReviewStatusUserAccessLevel `json:"userAccessLevel,omitempty"`
}
