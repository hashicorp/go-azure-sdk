package extendedueinformation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedUeInformationClient struct {
	Client *resourcemanager.Client
}

func NewExtendedUeInformationClientWithBaseURI(sdkApi sdkEnv.Api) (*ExtendedUeInformationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "extendedueinformation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExtendedUeInformationClient: %+v", err)
	}

	return &ExtendedUeInformationClient{
		Client: client,
	}, nil
}
