package validateprobe

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateProbeClient struct {
	Client *resourcemanager.Client
}

func NewValidateProbeClientWithBaseURI(sdkApi sdkEnv.Api) (*ValidateProbeClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "validateprobe", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ValidateProbeClient: %+v", err)
	}

	return &ValidateProbeClient{
		Client: client,
	}, nil
}
