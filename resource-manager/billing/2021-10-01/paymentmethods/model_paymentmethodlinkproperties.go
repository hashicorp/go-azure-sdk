package paymentmethods

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PaymentMethodLinkProperties struct {
	PaymentMethod *PaymentMethodProjectionProperties `json:"paymentMethod,omitempty"`
}
