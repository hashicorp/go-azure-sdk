package location

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportedSku struct {
	BatchSupportEndOfLife *string          `json:"batchSupportEndOfLife,omitempty"`
	Capabilities          *[]SkuCapability `json:"capabilities,omitempty"`
	FamilyName            *string          `json:"familyName,omitempty"`
	Name                  *string          `json:"name,omitempty"`
}

func (o *SupportedSku) GetBatchSupportEndOfLifeAsTime() (*time.Time, error) {
	if o.BatchSupportEndOfLife == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BatchSupportEndOfLife, "2006-01-02T15:04:05Z07:00")
}

func (o *SupportedSku) SetBatchSupportEndOfLifeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BatchSupportEndOfLife = &formatted
}
