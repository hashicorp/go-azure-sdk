package reservationdetails

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationDetailsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReservationDetailsClientWithBaseURI(endpoint string) ReservationDetailsClient {
	return ReservationDetailsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
