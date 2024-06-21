package deletedconfigurationstores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedConfigurationStoreProperties struct {
	ConfigurationStoreId   *string            `json:"configurationStoreId,omitempty"`
	DeletionDate           *string            `json:"deletionDate,omitempty"`
	Location               *string            `json:"location,omitempty"`
	PurgeProtectionEnabled *bool              `json:"purgeProtectionEnabled,omitempty"`
	ScheduledPurgeDate     *string            `json:"scheduledPurgeDate,omitempty"`
	Tags                   *map[string]string `json:"tags,omitempty"`
}
