package summaryrules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SummaryLogsRetryBinProperties struct {
	RetryBinStartTime string `json:"retryBinStartTime"`
}

func (o *SummaryLogsRetryBinProperties) GetRetryBinStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.RetryBinStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SummaryLogsRetryBinProperties) SetRetryBinStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RetryBinStartTime = formatted
}
