package securescorecontroldefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoreControlDefinitionsClient struct {
	Client *resourcemanager.Client
}

func NewSecureScoreControlDefinitionsClientWithBaseURI(sdkApi sdkEnv.Api) (*SecureScoreControlDefinitionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securescorecontroldefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecureScoreControlDefinitionsClient: %+v", err)
	}

	return &SecureScoreControlDefinitionsClient{
		Client: client,
	}, nil
}
