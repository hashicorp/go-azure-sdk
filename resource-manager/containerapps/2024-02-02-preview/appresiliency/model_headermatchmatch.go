package appresiliency

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HeaderMatchMatch struct {
	ExactMatch  *string `json:"exactMatch,omitempty"`
	PrefixMatch *string `json:"prefixMatch,omitempty"`
	RegexMatch  *string `json:"regexMatch,omitempty"`
	SuffixMatch *string `json:"suffixMatch,omitempty"`
}
