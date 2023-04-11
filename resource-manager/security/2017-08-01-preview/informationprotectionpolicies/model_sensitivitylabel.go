package informationprotectionpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabel struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	Order       *int64  `json:"order,omitempty"`
	Rank        *Rank   `json:"rank,omitempty"`
}
