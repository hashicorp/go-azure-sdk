package rules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterType string

const (
	FilterTypeCorrelationFilter FilterType = "CorrelationFilter"
	FilterTypeSqlFilter         FilterType = "SqlFilter"
)

func PossibleValuesForFilterType() []string {
	return []string{
		string(FilterTypeCorrelationFilter),
		string(FilterTypeSqlFilter),
	}
}

func (s *FilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFilterType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FilterType(decoded)
	return nil
}
