package accounts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageInformation struct {
	BatchSupportEndOfLife *string          `json:"batchSupportEndOfLife,omitempty"`
	Capabilities          *[]string        `json:"capabilities,omitempty"`
	ImageReference        ImageReference   `json:"imageReference"`
	NodeAgentSKUId        string           `json:"nodeAgentSKUId"`
	OsType                OSType           `json:"osType"`
	VerificationType      VerificationType `json:"verificationType"`
}

func (o *ImageInformation) GetBatchSupportEndOfLifeAsTime() (*time.Time, error) {
	if o.BatchSupportEndOfLife == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BatchSupportEndOfLife, "2006-01-02T15:04:05Z07:00")
}

func (o *ImageInformation) SetBatchSupportEndOfLifeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BatchSupportEndOfLife = &formatted
}
