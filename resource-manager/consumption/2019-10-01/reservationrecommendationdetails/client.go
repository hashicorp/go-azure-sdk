package reservationrecommendationdetails

import "github.com/Azure/go-autorest/autorest"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsClient struct {
	Client  autorest.Client
	baseUri string
}

func NewReservationRecommendationDetailsClientWithBaseURI(endpoint string) ReservationRecommendationDetailsClient {
	return ReservationRecommendationDetailsClient{
		Client:  autorest.NewClientWithUserAgent(userAgent()),
		baseUri: endpoint,
	}
}
