package securescore

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreClient struct {
	Client *resourcemanager.Client
}

func NewSecureScoreClientWithBaseURI(sdkApi sdkEnv.Api) (*SecureScoreClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securescore", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecureScoreClient: %+v", err)
	}

	return &SecureScoreClient{
		Client: client,
	}, nil
}
