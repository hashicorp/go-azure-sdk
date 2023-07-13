package querytexts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryTextsClient struct {
	Client *resourcemanager.Client
}

func NewQueryTextsClientWithBaseURI(api environments.Api) (*QueryTextsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "querytexts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QueryTextsClient: %+v", err)
	}

	return &QueryTextsClient{
		Client: client,
	}, nil
}
