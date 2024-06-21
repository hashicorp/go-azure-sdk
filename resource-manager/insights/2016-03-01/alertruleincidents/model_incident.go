package alertruleincidents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Incident struct {
	ActivatedTime *string `json:"activatedTime,omitempty"`
	IsActive      *bool   `json:"isActive,omitempty"`
	Name          *string `json:"name,omitempty"`
	ResolvedTime  *string `json:"resolvedTime,omitempty"`
	RuleName      *string `json:"ruleName,omitempty"`
}
