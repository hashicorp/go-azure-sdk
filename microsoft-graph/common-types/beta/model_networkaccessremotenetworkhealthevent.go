package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessRemoteNetworkHealthEvent struct {
	BgpRoutesAdvertisedCount *int64                                       `json:"bgpRoutesAdvertisedCount,omitempty"`
	CreatedDateTime          *string                                      `json:"createdDateTime,omitempty"`
	Description              *string                                      `json:"description,omitempty"`
	DestinationIp            *string                                      `json:"destinationIp,omitempty"`
	Id                       *string                                      `json:"id,omitempty"`
	ODataType                *string                                      `json:"@odata.type,omitempty"`
	ReceivedBytes            *int64                                       `json:"receivedBytes,omitempty"`
	RemoteNetworkId          *string                                      `json:"remoteNetworkId,omitempty"`
	SentBytes                *int64                                       `json:"sentBytes,omitempty"`
	SourceIp                 *string                                      `json:"sourceIp,omitempty"`
	Status                   *NetworkaccessRemoteNetworkHealthEventStatus `json:"status,omitempty"`
}
