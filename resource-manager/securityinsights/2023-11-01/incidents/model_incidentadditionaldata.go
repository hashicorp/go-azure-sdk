package incidents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentAdditionalData struct {
	AlertProductNames   *[]string       `json:"alertProductNames,omitempty"`
	AlertsCount         *int64          `json:"alertsCount,omitempty"`
	BookmarksCount      *int64          `json:"bookmarksCount,omitempty"`
	CommentsCount       *int64          `json:"commentsCount,omitempty"`
	ProviderIncidentUrl *string         `json:"providerIncidentUrl,omitempty"`
	Tactics             *[]AttackTactic `json:"tactics,omitempty"`
}
