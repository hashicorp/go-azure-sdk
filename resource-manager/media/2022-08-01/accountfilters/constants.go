package accountfilters

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FilterTrackPropertyCompareOperation string

const (
	FilterTrackPropertyCompareOperationEqual    FilterTrackPropertyCompareOperation = "Equal"
	FilterTrackPropertyCompareOperationNotEqual FilterTrackPropertyCompareOperation = "NotEqual"
)

func PossibleValuesForFilterTrackPropertyCompareOperation() []string {
	return []string{
		string(FilterTrackPropertyCompareOperationEqual),
		string(FilterTrackPropertyCompareOperationNotEqual),
	}
}

func (s *FilterTrackPropertyCompareOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFilterTrackPropertyCompareOperation() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FilterTrackPropertyCompareOperation(decoded)
	return nil
}

type FilterTrackPropertyType string

const (
	FilterTrackPropertyTypeBitrate  FilterTrackPropertyType = "Bitrate"
	FilterTrackPropertyTypeFourCC   FilterTrackPropertyType = "FourCC"
	FilterTrackPropertyTypeLanguage FilterTrackPropertyType = "Language"
	FilterTrackPropertyTypeName     FilterTrackPropertyType = "Name"
	FilterTrackPropertyTypeType     FilterTrackPropertyType = "Type"
	FilterTrackPropertyTypeUnknown  FilterTrackPropertyType = "Unknown"
)

func PossibleValuesForFilterTrackPropertyType() []string {
	return []string{
		string(FilterTrackPropertyTypeBitrate),
		string(FilterTrackPropertyTypeFourCC),
		string(FilterTrackPropertyTypeLanguage),
		string(FilterTrackPropertyTypeName),
		string(FilterTrackPropertyTypeType),
		string(FilterTrackPropertyTypeUnknown),
	}
}

func (s *FilterTrackPropertyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForFilterTrackPropertyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = FilterTrackPropertyType(decoded)
	return nil
}
