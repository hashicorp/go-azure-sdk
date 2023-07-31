package reservationdetails

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationDetailsClient struct {
	Client *resourcemanager.Client
}

func NewReservationDetailsClientWithBaseURI(api sdkEnv.Api) (*ReservationDetailsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "reservationdetails", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationDetailsClient: %+v", err)
	}

	return &ReservationDetailsClient{
		Client: client,
	}, nil
}
