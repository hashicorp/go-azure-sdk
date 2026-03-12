package analysisdefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AnalysisDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewAnalysisDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*AnalysisDefinitionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "analysisdefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AnalysisDefinitionsClient: %+v", err)
	}

	return &AnalysisDefinitionsClient{
		Client: client,
	}, nil
}
