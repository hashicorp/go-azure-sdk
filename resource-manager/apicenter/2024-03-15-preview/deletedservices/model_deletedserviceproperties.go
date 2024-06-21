package deletedservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServiceProperties struct {
	ScheduledPurgeDate *string `json:"scheduledPurgeDate,omitempty"`
	SoftDeletionDate   *string `json:"softDeletionDate,omitempty"`
}
