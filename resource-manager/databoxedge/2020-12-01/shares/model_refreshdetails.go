package shares

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RefreshDetails struct {
	ErrorManifestFile                *string `json:"errorManifestFile,omitempty"`
	InProgressRefreshJobId           *string `json:"inProgressRefreshJobId,omitempty"`
	LastCompletedRefreshJobTimeInUTC *string `json:"lastCompletedRefreshJobTimeInUTC,omitempty"`
	LastJob                          *string `json:"lastJob,omitempty"`
}

func (o *RefreshDetails) GetLastCompletedRefreshJobTimeInUTCAsTime() (*time.Time, error) {
	if o.LastCompletedRefreshJobTimeInUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastCompletedRefreshJobTimeInUTC, "2006-01-02T15:04:05Z07:00")
}

func (o *RefreshDetails) SetLastCompletedRefreshJobTimeInUTCAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastCompletedRefreshJobTimeInUTC = &formatted
}
