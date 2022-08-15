package reservationtransactions

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationTransactionsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReservationTransactionsClientWithBaseURI(endpoint string) ReservationTransactionsClient {
	return ReservationTransactionsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
