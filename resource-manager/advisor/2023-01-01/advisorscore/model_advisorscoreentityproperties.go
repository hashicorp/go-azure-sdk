package advisorscore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorScoreEntityProperties struct {
	LastRefreshedScore *ScoreEntity                         `json:"lastRefreshedScore,omitempty"`
	TimeSeries         *[]TimeSeriesEntityTimeSeriesInlined `json:"timeSeries,omitempty"`
}
