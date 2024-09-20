package reservationorder

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationOrderClient struct {
	Client *resourcemanager.Client
}

func NewReservationOrderClientWithBaseURI(sdkApi sdkEnv.Api) (*ReservationOrderClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "reservationorder", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationOrderClient: %+v", err)
	}

	return &ReservationOrderClient{
		Client: client,
	}, nil
}
