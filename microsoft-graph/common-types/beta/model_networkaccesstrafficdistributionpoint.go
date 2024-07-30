package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTrafficDistributionPoint struct {
	InternetAccessTrafficCount     *int64  `json:"internetAccessTrafficCount,omitempty"`
	Microsoft365AccessTrafficCount *int64  `json:"microsoft365AccessTrafficCount,omitempty"`
	ODataType                      *string `json:"@odata.type,omitempty"`
	PrivateAccessTrafficCount      *int64  `json:"privateAccessTrafficCount,omitempty"`
	TimeStampDateTime              *string `json:"timeStampDateTime,omitempty"`
	TotalTrafficCount              *int64  `json:"totalTrafficCount,omitempty"`
}
