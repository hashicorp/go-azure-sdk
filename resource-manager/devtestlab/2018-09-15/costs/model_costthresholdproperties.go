package costs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CostThresholdProperties struct {
	DisplayOnChart               *CostThresholdStatus               `json:"displayOnChart,omitempty"`
	NotificationSent             *string                            `json:"notificationSent,omitempty"`
	PercentageThreshold          *PercentageCostThresholdProperties `json:"percentageThreshold,omitempty"`
	SendNotificationWhenExceeded *CostThresholdStatus               `json:"sendNotificationWhenExceeded,omitempty"`
	ThresholdId                  *string                            `json:"thresholdId,omitempty"`
}
