package detectordefinitionresourceoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorDefinitionResourceOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewDetectorDefinitionResourceOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*DetectorDefinitionResourceOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "detectordefinitionresourceoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DetectorDefinitionResourceOperationGroupClient: %+v", err)
	}

	return &DetectorDefinitionResourceOperationGroupClient{
		Client: client,
	}, nil
}
