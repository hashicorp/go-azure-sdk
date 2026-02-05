package files

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileProperties struct {
	ContentLength int64   `json:"contentLength"`
	ContentType   *string `json:"contentType,omitempty"`
	CreationTime  *string `json:"creationTime,omitempty"`
	FileMode      *string `json:"fileMode,omitempty"`
	LastModified  string  `json:"lastModified"`
}

func (o *FileProperties) GetCreationTimeAsTime() (*time.Time, error) {
	if o.CreationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *FileProperties) SetCreationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationTime = &formatted
}

func (o *FileProperties) GetLastModifiedAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastModified, "2006-01-02T15:04:05Z07:00")
}

func (o *FileProperties) SetLastModifiedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModified = formatted
}
