package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlterationResponse struct {
	ODataType           *string                                `json:"@odata.type,omitempty"`
	OriginalQueryString *string                                `json:"originalQueryString,omitempty"`
	QueryAlteration     *SearchAlteration                      `json:"queryAlteration,omitempty"`
	QueryAlterationType *AlterationResponseQueryAlterationType `json:"queryAlterationType,omitempty"`
}
