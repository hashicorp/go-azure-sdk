package recoverypointscrr

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureWorkloadSQLRecoveryPointExtendedInfo struct {
	DataDirectoryPaths     *[]SQLDataDirectory `json:"dataDirectoryPaths,omitempty"`
	DataDirectoryTimeInUTC *string             `json:"dataDirectoryTimeInUTC,omitempty"`
}

func (o *AzureWorkloadSQLRecoveryPointExtendedInfo) GetDataDirectoryTimeInUTCAsTime() (*time.Time, error) {
	if o.DataDirectoryTimeInUTC == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DataDirectoryTimeInUTC, "2006-01-02T15:04:05Z07:00")
}
