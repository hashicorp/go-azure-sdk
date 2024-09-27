package standardassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardAssignmentMetadata struct {
	CreatedBy     *string `json:"createdBy,omitempty"`
	CreatedOn     *string `json:"createdOn,omitempty"`
	LastUpdatedBy *string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedOn *string `json:"lastUpdatedOn,omitempty"`
}

func (o *StandardAssignmentMetadata) GetCreatedOnAsTime() (*time.Time, error) {
	if o.CreatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *StandardAssignmentMetadata) SetCreatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedOn = &formatted
}

func (o *StandardAssignmentMetadata) GetLastUpdatedOnAsTime() (*time.Time, error) {
	if o.LastUpdatedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *StandardAssignmentMetadata) SetLastUpdatedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedOn = &formatted
}
