package rules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DeliveryRuleAction = UrlRewriteAction{}

type UrlRewriteAction struct {
	Parameters UrlRewriteActionParameters `json:"parameters"`

	// Fields inherited from DeliveryRuleAction
}

var _ json.Marshaler = UrlRewriteAction{}

func (s UrlRewriteAction) MarshalJSON() ([]byte, error) {
	type wrapper UrlRewriteAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UrlRewriteAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UrlRewriteAction: %+v", err)
	}
	decoded["name"] = "UrlRewrite"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UrlRewriteAction: %+v", err)
	}

	return encoded, nil
}
