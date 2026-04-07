package testraiexternalsafetyprovider

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestRaiExternalSafetyProviderClient struct {
	Client *resourcemanager.Client
}

func NewTestRaiExternalSafetyProviderClientWithBaseURI(sdkApi sdkEnv.Api) (*TestRaiExternalSafetyProviderClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "testraiexternalsafetyprovider", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TestRaiExternalSafetyProviderClient: %+v", err)
	}

	return &TestRaiExternalSafetyProviderClient{
		Client: client,
	}, nil
}
