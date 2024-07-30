package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationQuarantine struct {
	CurrentBegan *string                          `json:"currentBegan,omitempty"`
	Error        *SynchronizationError            `json:"error,omitempty"`
	NextAttempt  *string                          `json:"nextAttempt,omitempty"`
	ODataType    *string                          `json:"@odata.type,omitempty"`
	Reason       *SynchronizationQuarantineReason `json:"reason,omitempty"`
	SeriesBegan  *string                          `json:"seriesBegan,omitempty"`
	SeriesCount  *int64                           `json:"seriesCount,omitempty"`
}
