package reservationrecommendationdetails

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationDetailsClient struct {
	Client *resourcemanager.Client
}

func NewReservationRecommendationDetailsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReservationRecommendationDetailsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "reservationrecommendationdetails", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationRecommendationDetailsClient: %+v", err)
	}

	return &ReservationRecommendationDetailsClient{
		Client: client,
	}, nil
}
