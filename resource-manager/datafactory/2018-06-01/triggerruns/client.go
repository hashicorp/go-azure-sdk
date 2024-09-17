package triggerruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerrunsClient struct {
	Client *resourcemanager.Client
}

func NewTriggerrunsClientWithBaseURI(sdkApi sdkEnv.Api) (*TriggerrunsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "triggerruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggerrunsClient: %+v", err)
	}

	return &TriggerrunsClient{
		Client: client,
	}, nil
}
