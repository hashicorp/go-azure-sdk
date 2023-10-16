package componentcontinuousexportapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationInsightsComponentExportConfiguration struct {
	ApplicationName                  *string `json:"ApplicationName,omitempty"`
	ContainerName                    *string `json:"ContainerName,omitempty"`
	DestinationAccountId             *string `json:"DestinationAccountId,omitempty"`
	DestinationStorageLocationId     *string `json:"DestinationStorageLocationId,omitempty"`
	DestinationStorageSubscriptionId *string `json:"DestinationStorageSubscriptionId,omitempty"`
	DestinationType                  *string `json:"DestinationType,omitempty"`
	ExportId                         *string `json:"ExportId,omitempty"`
	ExportStatus                     *string `json:"ExportStatus,omitempty"`
	InstrumentationKey               *string `json:"InstrumentationKey,omitempty"`
	IsUserEnabled                    *string `json:"IsUserEnabled,omitempty"`
	LastGapTime                      *string `json:"LastGapTime,omitempty"`
	LastSuccessTime                  *string `json:"LastSuccessTime,omitempty"`
	LastUserUpdate                   *string `json:"LastUserUpdate,omitempty"`
	NotificationQueueEnabled         *string `json:"NotificationQueueEnabled,omitempty"`
	PermanentErrorReason             *string `json:"PermanentErrorReason,omitempty"`
	RecordTypes                      *string `json:"RecordTypes,omitempty"`
	ResourceGroup                    *string `json:"ResourceGroup,omitempty"`
	StorageName                      *string `json:"StorageName,omitempty"`
	SubscriptionId                   *string `json:"SubscriptionId,omitempty"`
}
