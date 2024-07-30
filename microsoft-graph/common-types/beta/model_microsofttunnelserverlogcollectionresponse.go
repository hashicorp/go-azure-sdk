package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServerLogCollectionResponse struct {
	EndDateTime     *string                                           `json:"endDateTime,omitempty"`
	ExpiryDateTime  *string                                           `json:"expiryDateTime,omitempty"`
	Id              *string                                           `json:"id,omitempty"`
	ODataType       *string                                           `json:"@odata.type,omitempty"`
	RequestDateTime *string                                           `json:"requestDateTime,omitempty"`
	ServerId        *string                                           `json:"serverId,omitempty"`
	SizeInBytes     *int64                                            `json:"sizeInBytes,omitempty"`
	StartDateTime   *string                                           `json:"startDateTime,omitempty"`
	Status          *MicrosoftTunnelServerLogCollectionResponseStatus `json:"status,omitempty"`
}
