package alertrules

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRule struct {
	Action            RuleAction    `json:"action"`
	Actions           *[]RuleAction `json:"actions,omitempty"`
	Condition         RuleCondition `json:"condition"`
	Description       *string       `json:"description,omitempty"`
	IsEnabled         bool          `json:"isEnabled"`
	LastUpdatedTime   *string       `json:"lastUpdatedTime,omitempty"`
	Name              string        `json:"name"`
	ProvisioningState *string       `json:"provisioningState,omitempty"`
}

func (o *AlertRule) GetLastUpdatedTimeAsTime() (*time.Time, error) {
	if o.LastUpdatedTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AlertRule) SetLastUpdatedTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedTime = &formatted
}

var _ json.Unmarshaler = &AlertRule{}

func (s *AlertRule) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Description       *string `json:"description,omitempty"`
		IsEnabled         bool    `json:"isEnabled"`
		LastUpdatedTime   *string `json:"lastUpdatedTime,omitempty"`
		Name              string  `json:"name"`
		ProvisioningState *string `json:"provisioningState,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Description = decoded.Description
	s.IsEnabled = decoded.IsEnabled
	s.LastUpdatedTime = decoded.LastUpdatedTime
	s.Name = decoded.Name
	s.ProvisioningState = decoded.ProvisioningState

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AlertRule into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["action"]; ok {
		impl, err := UnmarshalRuleActionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Action' for 'AlertRule': %+v", err)
		}
		s.Action = impl
	}

	if v, ok := temp["actions"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Actions into list []json.RawMessage: %+v", err)
		}

		output := make([]RuleAction, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRuleActionImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Actions' for 'AlertRule': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Actions = &output
	}

	if v, ok := temp["condition"]; ok {
		impl, err := UnmarshalRuleConditionImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Condition' for 'AlertRule': %+v", err)
		}
		s.Condition = impl
	}

	return nil
}
