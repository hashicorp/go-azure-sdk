package dscnodeconfiguration

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent ContentSourceType = "embeddedContent"
	ContentSourceTypeUri             ContentSourceType = "uri"
)

func PossibleValuesForContentSourceType() []string {
	return []string{
		string(ContentSourceTypeEmbeddedContent),
		string(ContentSourceTypeUri),
	}
}

func (s *ContentSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForContentSourceType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = ContentSourceType(decoded)
	return nil
}
