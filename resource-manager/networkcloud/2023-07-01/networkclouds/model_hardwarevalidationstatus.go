package networkclouds

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareValidationStatus struct {
	LastValidationTime *string                                   `json:"lastValidationTime,omitempty"`
	Result             *BareMetalMachineHardwareValidationResult `json:"result,omitempty"`
}

func (o *HardwareValidationStatus) GetLastValidationTimeAsTime() (*time.Time, error) {
	if o.LastValidationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastValidationTime, "2006-01-02T15:04:05Z07:00")
}
