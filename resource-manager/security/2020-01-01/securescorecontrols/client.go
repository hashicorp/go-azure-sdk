package securescorecontrols

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlsClient struct {
	Client *resourcemanager.Client
}

func NewSecureScoreControlsClientWithBaseURI(sdkApi sdkEnv.Api) (*SecureScoreControlsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securescorecontrols", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecureScoreControlsClient: %+v", err)
	}

	return &SecureScoreControlsClient{
		Client: client,
	}, nil
}
