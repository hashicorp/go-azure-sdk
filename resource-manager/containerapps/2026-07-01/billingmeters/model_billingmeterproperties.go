package billingmeters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingMeterProperties struct {
	Category    *string `json:"category,omitempty"`
	DisplayName *string `json:"displayName,omitempty"`
	MeterType   *string `json:"meterType,omitempty"`
}
