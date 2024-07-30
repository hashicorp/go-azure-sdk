package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationEvent struct {
	Count     *int64  `json:"count,omitempty"`
	EventName *string `json:"eventName,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
