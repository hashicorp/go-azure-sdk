package incidentcomments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentCommentProperties struct {
	Author              *ClientInfo `json:"author"`
	CreatedTimeUtc      *string     `json:"createdTimeUtc,omitempty"`
	LastModifiedTimeUtc *string     `json:"lastModifiedTimeUtc,omitempty"`
	Message             string      `json:"message"`
}

func (o *IncidentCommentProperties) GetCreatedTimeUtcAsTime() (*time.Time, error) {
	if o.CreatedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentCommentProperties) SetCreatedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedTimeUtc = &formatted
}

func (o *IncidentCommentProperties) GetLastModifiedTimeUtcAsTime() (*time.Time, error) {
	if o.LastModifiedTimeUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModifiedTimeUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IncidentCommentProperties) SetLastModifiedTimeUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModifiedTimeUtc = &formatted
}
