package ransomwarereports

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SuspectFile struct {
	FileTimestamp   *string `json:"fileTimestamp,omitempty"`
	SuspectFileName *string `json:"suspectFileName,omitempty"`
}

func (o *SuspectFile) GetFileTimestampAsTime() (*time.Time, error) {
	if o.FileTimestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FileTimestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SuspectFile) SetFileTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FileTimestamp = &formatted
}
