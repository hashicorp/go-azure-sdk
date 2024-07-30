package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationProgress struct {
	CompletedUnits              *int64  `json:"completedUnits,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
	ProgressObservationDateTime *string `json:"progressObservationDateTime,omitempty"`
	TotalUnits                  *int64  `json:"totalUnits,omitempty"`
	Units                       *string `json:"units,omitempty"`
}
