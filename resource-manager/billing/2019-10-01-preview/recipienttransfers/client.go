package recipienttransfers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecipientTransfersClient struct {
	Client *resourcemanager.Client
}

func NewRecipientTransfersClientWithBaseURI(sdkApi sdkEnv.Api) (*RecipientTransfersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "recipienttransfers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RecipientTransfersClient: %+v", err)
	}

	return &RecipientTransfersClient{
		Client: client,
	}, nil
}
