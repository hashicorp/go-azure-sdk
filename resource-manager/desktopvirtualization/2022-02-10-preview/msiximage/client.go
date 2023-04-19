package msiximage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MsixImageClient struct {
	Client *resourcemanager.Client
}

func NewMsixImageClientWithBaseURI(api environments.Api) (*MsixImageClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "msiximage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MsixImageClient: %+v", err)
	}

	return &MsixImageClient{
		Client: client,
	}, nil
}
