package alertruleincidents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Incident struct {
	ActivatedTime *string `json:"activatedTime,omitempty"`
	IsActive      *bool   `json:"isActive,omitempty"`
	Name          *string `json:"name,omitempty"`
	ResolvedTime  *string `json:"resolvedTime,omitempty"`
	RuleName      *string `json:"ruleName,omitempty"`
}

func (o *Incident) GetActivatedTimeAsTime() (*time.Time, error) {
	if o.ActivatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ActivatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Incident) GetResolvedTimeAsTime() (*time.Time, error) {
	if o.ResolvedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ResolvedTime, "2006-01-02T15:04:05Z07:00")
}
