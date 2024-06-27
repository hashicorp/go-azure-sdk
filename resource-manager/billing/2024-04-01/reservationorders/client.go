package reservationorders

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationOrdersClient struct {
	Client *resourcemanager.Client
}

func NewReservationOrdersClientWithBaseURI(sdkApi sdkEnv.Api) (*ReservationOrdersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "reservationorders", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationOrdersClient: %+v", err)
	}

	return &ReservationOrdersClient{
		Client: client,
	}, nil
}
