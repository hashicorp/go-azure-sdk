package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UriClickSecurityState struct {
	ClickAction   *string `json:"clickAction,omitempty"`
	ClickDateTime *string `json:"clickDateTime,omitempty"`
	Id            *string `json:"id,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	SourceId      *string `json:"sourceId,omitempty"`
	UriDomain     *string `json:"uriDomain,omitempty"`
	Verdict       *string `json:"verdict,omitempty"`
}
