package baremetalmachinekeysets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetsClient struct {
	Client *resourcemanager.Client
}

func NewBareMetalMachineKeySetsClientWithBaseURI(sdkApi sdkEnv.Api) (*BareMetalMachineKeySetsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "baremetalmachinekeysets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BareMetalMachineKeySetsClient: %+v", err)
	}

	return &BareMetalMachineKeySetsClient{
		Client: client,
	}, nil
}
