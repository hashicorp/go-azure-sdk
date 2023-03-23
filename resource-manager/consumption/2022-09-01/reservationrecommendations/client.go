package reservationrecommendations

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReservationRecommendationsClientWithBaseURI(endpoint string) ReservationRecommendationsClient {
	return ReservationRecommendationsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
