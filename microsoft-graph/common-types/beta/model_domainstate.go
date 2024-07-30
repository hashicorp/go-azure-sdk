package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainState struct {
	LastActionDateTime *string `json:"lastActionDateTime,omitempty"`
	ODataType          *string `json:"@odata.type,omitempty"`
	Operation          *string `json:"operation,omitempty"`
	Status             *string `json:"status,omitempty"`
}
