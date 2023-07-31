package graphservicesprods

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GraphservicesprodsClient struct {
	Client *resourcemanager.Client
}

func NewGraphservicesprodsClientWithBaseURI(api sdkEnv.Api) (*GraphservicesprodsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "graphservicesprods", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GraphservicesprodsClient: %+v", err)
	}

	return &GraphservicesprodsClient{
		Client: client,
	}, nil
}
