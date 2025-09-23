package groupquotasentities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotaRequestBaseProperties struct {
	Comments *string                              `json:"comments,omitempty"`
	Limit    *int64                               `json:"limit,omitempty"`
	Name     *GroupQuotaRequestBasePropertiesName `json:"name,omitempty"`
	Region   *string                              `json:"region,omitempty"`
}
