package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesProvisioningError struct {
	Category             *string `json:"category,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	OccurredDateTime     *string `json:"occurredDateTime,omitempty"`
	PropertyCausingError *string `json:"propertyCausingError,omitempty"`
	Value                *string `json:"value,omitempty"`
}
