package deletedbackupinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletionInfo struct {
	BillingEndDate     *string `json:"billingEndDate,omitempty"`
	DeleteActivityID   *string `json:"deleteActivityID,omitempty"`
	DeletionTime       *string `json:"deletionTime,omitempty"`
	ScheduledPurgeTime *string `json:"scheduledPurgeTime,omitempty"`
}
