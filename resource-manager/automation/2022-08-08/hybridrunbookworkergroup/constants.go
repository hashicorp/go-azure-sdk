package hybridrunbookworkergroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTypeEnum string

const (
	GroupTypeEnumSystem GroupTypeEnum = "System"
	GroupTypeEnumUser   GroupTypeEnum = "User"
)

func PossibleValuesForGroupTypeEnum() []string {
	return []string{
		string(GroupTypeEnumSystem),
		string(GroupTypeEnumUser),
	}
}

func (s *GroupTypeEnum) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForGroupTypeEnum() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = GroupTypeEnum(decoded)
	return nil
}
