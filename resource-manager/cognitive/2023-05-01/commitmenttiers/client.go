package commitmenttiers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentTiersClient struct {
	Client *resourcemanager.Client
}

func NewCommitmentTiersClientWithBaseURI(sdkApi sdkEnv.Api) (*CommitmentTiersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "commitmenttiers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommitmentTiersClient: %+v", err)
	}

	return &CommitmentTiersClient{
		Client: client,
	}, nil
}
