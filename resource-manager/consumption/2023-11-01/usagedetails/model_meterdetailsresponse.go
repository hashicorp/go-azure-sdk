package usagedetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeterDetailsResponse struct {
	MeterCategory    *string `json:"meterCategory,omitempty"`
	MeterName        *string `json:"meterName,omitempty"`
	MeterSubCategory *string `json:"meterSubCategory,omitempty"`
	ServiceFamily    *string `json:"serviceFamily,omitempty"`
	UnitOfMeasure    *string `json:"unitOfMeasure,omitempty"`
}
