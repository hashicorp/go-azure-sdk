package resetcifspassword

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetCifsPasswordClient struct {
	Client *resourcemanager.Client
}

func NewResetCifsPasswordClientWithBaseURI(sdkApi sdkEnv.Api) (*ResetCifsPasswordClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "resetcifspassword", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResetCifsPasswordClient: %+v", err)
	}

	return &ResetCifsPasswordClient{
		Client: client,
	}, nil
}
