package topleveldomains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TldLegalAgreement struct {
	AgreementKey string  `json:"agreementKey"`
	Content      string  `json:"content"`
	Title        string  `json:"title"`
	Url          *string `json:"url,omitempty"`
}
