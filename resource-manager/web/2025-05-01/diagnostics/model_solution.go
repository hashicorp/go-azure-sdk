package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Solution struct {
	Data        *[][]NameValuePair `json:"data,omitempty"`
	Description *string            `json:"description,omitempty"`
	DisplayName *string            `json:"displayName,omitempty"`
	Id          *float64           `json:"id,omitempty"`
	Metadata    *[][]NameValuePair `json:"metadata,omitempty"`
	Order       *float64           `json:"order,omitempty"`
	Type        *SolutionType      `json:"type,omitempty"`
}
