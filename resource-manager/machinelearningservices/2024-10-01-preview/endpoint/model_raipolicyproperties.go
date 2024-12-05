package endpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RaiPolicyProperties struct {
	BasePolicyName       *string                   `json:"basePolicyName,omitempty"`
	CompletionBlocklists *[]RaiBlocklistConfig     `json:"completionBlocklists,omitempty"`
	ContentFilters       *[]RaiPolicyContentFilter `json:"contentFilters,omitempty"`
	Mode                 *RaiPolicyMode            `json:"mode,omitempty"`
	PromptBlocklists     *[]RaiBlocklistConfig     `json:"promptBlocklists,omitempty"`
	Type                 *RaiPolicyType            `json:"type,omitempty"`
}
