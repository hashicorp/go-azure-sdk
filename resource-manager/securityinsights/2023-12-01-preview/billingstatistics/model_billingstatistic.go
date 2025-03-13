package billingstatistics

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingStatistic interface {
	BillingStatistic() BaseBillingStatisticImpl
}

var _ BillingStatistic = BaseBillingStatisticImpl{}

type BaseBillingStatisticImpl struct {
	Etag       *string                `json:"etag,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Kind       BillingStatisticKind   `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BaseBillingStatisticImpl) BillingStatistic() BaseBillingStatisticImpl {
	return s
}

var _ BillingStatistic = RawBillingStatisticImpl{}

// RawBillingStatisticImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBillingStatisticImpl struct {
	billingStatistic BaseBillingStatisticImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawBillingStatisticImpl) BillingStatistic() BaseBillingStatisticImpl {
	return s.billingStatistic
}

func UnmarshalBillingStatisticImplementation(input []byte) (BillingStatistic, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BillingStatistic into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "SapSolutionUsage") {
		var out SapSolutionUsageStatistic
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SapSolutionUsageStatistic: %+v", err)
		}
		return out, nil
	}

	var parent BaseBillingStatisticImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBillingStatisticImpl: %+v", err)
	}

	return RawBillingStatisticImpl{
		billingStatistic: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
