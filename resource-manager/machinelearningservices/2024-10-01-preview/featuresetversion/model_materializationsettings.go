package featuresetversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaterializationSettings struct {
	Notification       *NotificationSetting            `json:"notification,omitempty"`
	Resource           *MaterializationComputeResource `json:"resource,omitempty"`
	Schedule           *RecurrenceTrigger              `json:"schedule,omitempty"`
	SparkConfiguration *map[string]string              `json:"sparkConfiguration,omitempty"`
	StoreType          *MaterializationStoreType       `json:"storeType,omitempty"`
}
