package inferencepool

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferencePoolClient struct {
	Client *resourcemanager.Client
}

func NewInferencePoolClientWithBaseURI(sdkApi sdkEnv.Api) (*InferencePoolClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "inferencepool", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InferencePoolClient: %+v", err)
	}

	return &InferencePoolClient{
		Client: client,
	}, nil
}
