package labelingjob

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExportSummary = DatasetExportSummary{}

type DatasetExportSummary struct {
	LabeledAssetName *string `json:"labeledAssetName,omitempty"`

	// Fields inherited from ExportSummary
	EndDateTime      *string `json:"endDateTime,omitempty"`
	ExportedRowCount *int64  `json:"exportedRowCount,omitempty"`
	LabelingJobId    *string `json:"labelingJobId,omitempty"`
	StartDateTime    *string `json:"startDateTime,omitempty"`
}

func (o *DatasetExportSummary) GetEndDateTimeAsTime() (*time.Time, error) {
	if o.EndDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DatasetExportSummary) SetEndDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDateTime = &formatted
}

func (o *DatasetExportSummary) GetStartDateTimeAsTime() (*time.Time, error) {
	if o.StartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DatasetExportSummary) SetStartDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDateTime = &formatted
}

var _ json.Marshaler = DatasetExportSummary{}

func (s DatasetExportSummary) MarshalJSON() ([]byte, error) {
	type wrapper DatasetExportSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatasetExportSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatasetExportSummary: %+v", err)
	}
	decoded["format"] = "Dataset"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatasetExportSummary: %+v", err)
	}

	return encoded, nil
}
