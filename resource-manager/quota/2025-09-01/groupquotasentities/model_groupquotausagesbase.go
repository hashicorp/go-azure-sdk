package groupquotasentities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaUsagesBase struct {
	Limit  *int64                    `json:"limit,omitempty"`
	Name   *GroupQuotaUsagesBaseName `json:"name,omitempty"`
	Unit   *string                   `json:"unit,omitempty"`
	Usages *int64                    `json:"usages,omitempty"`
}
