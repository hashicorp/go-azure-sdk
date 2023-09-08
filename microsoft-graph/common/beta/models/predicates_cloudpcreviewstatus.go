package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCReviewStatusOperationPredicate struct {
	AzureStorageAccountId     *string
	AzureStorageAccountName   *string
	AzureStorageContainerName *string
	InReview                  *bool
	ODataType                 *string
	RestorePointDateTime      *string
	ReviewStartDateTime       *string
	SubscriptionId            *string
	SubscriptionName          *string
}

func (p CloudPCReviewStatusOperationPredicate) Matches(input CloudPCReviewStatus) bool {

	if p.AzureStorageAccountId != nil && (input.AzureStorageAccountId == nil || *p.AzureStorageAccountId != *input.AzureStorageAccountId) {
		return false
	}

	if p.AzureStorageAccountName != nil && (input.AzureStorageAccountName == nil || *p.AzureStorageAccountName != *input.AzureStorageAccountName) {
		return false
	}

	if p.AzureStorageContainerName != nil && (input.AzureStorageContainerName == nil || *p.AzureStorageContainerName != *input.AzureStorageContainerName) {
		return false
	}

	if p.InReview != nil && (input.InReview == nil || *p.InReview != *input.InReview) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RestorePointDateTime != nil && (input.RestorePointDateTime == nil || *p.RestorePointDateTime != *input.RestorePointDateTime) {
		return false
	}

	if p.ReviewStartDateTime != nil && (input.ReviewStartDateTime == nil || *p.ReviewStartDateTime != *input.ReviewStartDateTime) {
		return false
	}

	if p.SubscriptionId != nil && (input.SubscriptionId == nil || *p.SubscriptionId != *input.SubscriptionId) {
		return false
	}

	if p.SubscriptionName != nil && (input.SubscriptionName == nil || *p.SubscriptionName != *input.SubscriptionName) {
		return false
	}

	return true
}
