package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMedia struct {
	CalleeDevice  *CallRecordsDeviceInfo    `json:"calleeDevice,omitempty"`
	CalleeNetwork *CallRecordsNetworkInfo   `json:"calleeNetwork,omitempty"`
	CallerDevice  *CallRecordsDeviceInfo    `json:"callerDevice,omitempty"`
	CallerNetwork *CallRecordsNetworkInfo   `json:"callerNetwork,omitempty"`
	Label         *string                   `json:"label,omitempty"`
	ODataType     *string                   `json:"@odata.type,omitempty"`
	Streams       *[]CallRecordsMediaStream `json:"streams,omitempty"`
}
