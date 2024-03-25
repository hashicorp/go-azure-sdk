package cognitiveservicescommitmentplans

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CognitiveServicesCommitmentPlansClient struct {
	Client *resourcemanager.Client
}

func NewCognitiveServicesCommitmentPlansClientWithBaseURI(sdkApi sdkEnv.Api) (*CognitiveServicesCommitmentPlansClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "cognitiveservicescommitmentplans", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CognitiveServicesCommitmentPlansClient: %+v", err)
	}

	return &CognitiveServicesCommitmentPlansClient{
		Client: client,
	}, nil
}
