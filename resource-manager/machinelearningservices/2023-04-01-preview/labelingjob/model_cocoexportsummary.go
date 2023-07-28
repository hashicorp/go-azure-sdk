package labelingjob

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ExportSummary = CocoExportSummary{}

type CocoExportSummary struct {
	ContainerName *string `json:"containerName,omitempty"`
	SnapshotPath  *string `json:"snapshotPath,omitempty"`

	// Fields inherited from ExportSummary
	EndDateTime      *string `json:"endDateTime,omitempty"`
	ExportedRowCount *int64  `json:"exportedRowCount,omitempty"`
	LabelingJobId    *string `json:"labelingJobId,omitempty"`
	StartDateTime    *string `json:"startDateTime,omitempty"`
}

func (o *CocoExportSummary) GetEndDateTimeAsTime() (*time.Time, error) {
	if o.EndDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CocoExportSummary) SetEndDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDateTime = &formatted
}

func (o *CocoExportSummary) GetStartDateTimeAsTime() (*time.Time, error) {
	if o.StartDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *CocoExportSummary) SetStartDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDateTime = &formatted
}

var _ json.Marshaler = CocoExportSummary{}

func (s CocoExportSummary) MarshalJSON() ([]byte, error) {
	type wrapper CocoExportSummary
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CocoExportSummary: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CocoExportSummary: %+v", err)
	}
	decoded["format"] = "Coco"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CocoExportSummary: %+v", err)
	}

	return encoded, nil
}
