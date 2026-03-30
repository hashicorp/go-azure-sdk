package resourcevalidator

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceValidatorClient struct {
	Client *resourcemanager.Client
}

func NewResourceValidatorClientWithBaseURI(sdkApi sdkEnv.Api) (*ResourceValidatorClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "resourcevalidator", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ResourceValidatorClient: %+v", err)
	}

	return &ResourceValidatorClient{
		Client: client,
	}, nil
}
