package adaptivenetworkhardenings

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveNetworkHardeningProperties struct {
	EffectiveNetworkSecurityGroups *[]EffectiveNetworkSecurityGroups `json:"effectiveNetworkSecurityGroups,omitempty"`
	Rules                          *[]Rule                           `json:"rules,omitempty"`
	RulesCalculationTime           *string                           `json:"rulesCalculationTime,omitempty"`
}

func (o *AdaptiveNetworkHardeningProperties) GetRulesCalculationTimeAsTime() (*time.Time, error) {
	if o.RulesCalculationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RulesCalculationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AdaptiveNetworkHardeningProperties) SetRulesCalculationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RulesCalculationTime = &formatted
}
