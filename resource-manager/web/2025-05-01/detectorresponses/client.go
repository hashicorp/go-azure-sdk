package detectorresponses

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorResponsesClient struct {
	Client *resourcemanager.Client
}

func NewDetectorResponsesClientWithBaseURI(sdkApi sdkEnv.Api) (*DetectorResponsesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "detectorresponses", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DetectorResponsesClient: %+v", err)
	}

	return &DetectorResponsesClient{
		Client: client,
	}, nil
}
