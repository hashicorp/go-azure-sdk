package microsoftentraadministrators

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftEntraAdministratorsClient struct {
	Client *resourcemanager.Client
}

func NewMicrosoftEntraAdministratorsClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftEntraAdministratorsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "microsoftentraadministrators", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftEntraAdministratorsClient: %+v", err)
	}

	return &MicrosoftEntraAdministratorsClient{
		Client: client,
	}, nil
}
