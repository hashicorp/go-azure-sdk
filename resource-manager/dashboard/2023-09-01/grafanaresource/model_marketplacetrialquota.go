package grafanaresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceTrialQuota struct {
	AvailablePromotion *AvailablePromotion `json:"availablePromotion,omitempty"`
	GrafanaResourceId  *string             `json:"grafanaResourceId,omitempty"`
	TrialEndAt         *string             `json:"trialEndAt,omitempty"`
	TrialStartAt       *string             `json:"trialStartAt,omitempty"`
}
