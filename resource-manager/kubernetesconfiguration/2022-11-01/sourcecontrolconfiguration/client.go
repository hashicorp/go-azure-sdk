package sourcecontrolconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceControlConfigurationClient struct {
	Client *resourcemanager.Client
}

func NewSourceControlConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*SourceControlConfigurationClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sourcecontrolconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SourceControlConfigurationClient: %+v", err)
	}

	return &SourceControlConfigurationClient{
		Client: client,
	}, nil
}
