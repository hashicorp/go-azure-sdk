package componentcontinuousexportapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationInsightsComponentExportRequest struct {
	DestinationAccountId             *string `json:"DestinationAccountId,omitempty"`
	DestinationAddress               *string `json:"DestinationAddress,omitempty"`
	DestinationStorageLocationId     *string `json:"DestinationStorageLocationId,omitempty"`
	DestinationStorageSubscriptionId *string `json:"DestinationStorageSubscriptionId,omitempty"`
	DestinationType                  *string `json:"DestinationType,omitempty"`
	IsEnabled                        *string `json:"IsEnabled,omitempty"`
	NotificationQueueEnabled         *string `json:"NotificationQueueEnabled,omitempty"`
	NotificationQueueUri             *string `json:"NotificationQueueUri,omitempty"`
	RecordTypes                      *string `json:"RecordTypes,omitempty"`
}
