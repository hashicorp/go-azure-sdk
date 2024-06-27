package transfers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerInitiateTransferProperties struct {
	RecipientEmailId *string `json:"recipientEmailId,omitempty"`
	ResellerId       *string `json:"resellerId,omitempty"`
}
