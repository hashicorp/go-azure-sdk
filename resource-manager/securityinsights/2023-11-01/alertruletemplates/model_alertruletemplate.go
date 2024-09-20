package alertruletemplates

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRuleTemplate interface {
	AlertRuleTemplate() BaseAlertRuleTemplateImpl
}

var _ AlertRuleTemplate = BaseAlertRuleTemplateImpl{}

type BaseAlertRuleTemplateImpl struct {
	Id         *string                `json:"id,omitempty"`
	Kind       AlertRuleKind          `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseAlertRuleTemplateImpl) AlertRuleTemplate() BaseAlertRuleTemplateImpl {
	return s
}

var _ AlertRuleTemplate = RawAlertRuleTemplateImpl{}

// RawAlertRuleTemplateImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAlertRuleTemplateImpl struct {
	alertRuleTemplate BaseAlertRuleTemplateImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawAlertRuleTemplateImpl) AlertRuleTemplate() BaseAlertRuleTemplateImpl {
	return s.alertRuleTemplate
}

func UnmarshalAlertRuleTemplateImplementation(input []byte) (AlertRuleTemplate, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AlertRuleTemplate into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Fusion") {
		var out FusionAlertRuleTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FusionAlertRuleTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MicrosoftSecurityIncidentCreation") {
		var out MicrosoftSecurityIncidentCreationAlertRuleTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MicrosoftSecurityIncidentCreationAlertRuleTemplate: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Scheduled") {
		var out ScheduledAlertRuleTemplate
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScheduledAlertRuleTemplate: %+v", err)
		}
		return out, nil
	}

	var parent BaseAlertRuleTemplateImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAlertRuleTemplateImpl: %+v", err)
	}

	return RawAlertRuleTemplateImpl{
		alertRuleTemplate: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
