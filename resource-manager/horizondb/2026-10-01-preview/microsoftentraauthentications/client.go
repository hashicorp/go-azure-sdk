package microsoftentraauthentications

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftEntraAuthenticationsClient struct {
	Client *resourcemanager.Client
}

func NewMicrosoftEntraAuthenticationsClientWithBaseURI(sdkApi sdkEnv.Api) (*MicrosoftEntraAuthenticationsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "microsoftentraauthentications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MicrosoftEntraAuthenticationsClient: %+v", err)
	}

	return &MicrosoftEntraAuthenticationsClient{
		Client: client,
	}, nil
}
