package tuningoptions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TuningOptionsClient struct {
	Client *resourcemanager.Client
}

func NewTuningOptionsClientWithBaseURI(sdkApi sdkEnv.Api) (*TuningOptionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "tuningoptions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TuningOptionsClient: %+v", err)
	}

	return &TuningOptionsClient{
		Client: client,
	}, nil
}
