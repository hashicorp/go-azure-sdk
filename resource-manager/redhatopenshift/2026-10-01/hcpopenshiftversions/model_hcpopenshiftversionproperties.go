package hcpopenshiftversions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HcpOpenShiftVersionProperties struct {
	ChannelGroup       string `json:"channelGroup"`
	Enabled            bool   `json:"enabled"`
	EndOfLifeTimestamp string `json:"endOfLifeTimestamp"`
}

func (o *HcpOpenShiftVersionProperties) GetEndOfLifeTimestampAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndOfLifeTimestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *HcpOpenShiftVersionProperties) SetEndOfLifeTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndOfLifeTimestamp = formatted
}
