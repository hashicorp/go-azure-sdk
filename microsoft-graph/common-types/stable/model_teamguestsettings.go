package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamGuestSettings struct {
	AllowCreateUpdateChannels *bool   `json:"allowCreateUpdateChannels,omitempty"`
	AllowDeleteChannels       *bool   `json:"allowDeleteChannels,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
}
