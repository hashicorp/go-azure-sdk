package endpoint

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerlessEndpointCapacityReservation struct {
	CapacityReservationGroupId string `json:"capacityReservationGroupId"`
	EndpointReservedCapacity   *int64 `json:"endpointReservedCapacity,omitempty"`
}
