package billingaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddressValidationResponse struct {
	Status             *AddressValidationStatus `json:"status,omitempty"`
	SuggestedAddresses *[]AddressDetails        `json:"suggestedAddresses,omitempty"`
	ValidationMessage  *string                  `json:"validationMessage,omitempty"`
}
