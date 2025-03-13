package sourcecontrols

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlsClient struct {
	Client *resourcemanager.Client
}

func NewSourceControlsClientWithBaseURI(sdkApi sdkEnv.Api) (*SourceControlsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sourcecontrols", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SourceControlsClient: %+v", err)
	}

	return &SourceControlsClient{
		Client: client,
	}, nil
}
