package computenodes

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeAgentInformation struct {
	LastUpdateTime string `json:"lastUpdateTime"`
	Version        string `json:"version"`
}

func (o *NodeAgentInformation) GetLastUpdateTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastUpdateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *NodeAgentInformation) SetLastUpdateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdateTime = formatted
}
