package authorizationrulesdisasterrecoveryconfigs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationRulesDisasterRecoveryConfigsClient struct {
	Client *resourcemanager.Client
}

func NewAuthorizationRulesDisasterRecoveryConfigsClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthorizationRulesDisasterRecoveryConfigsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "authorizationrulesdisasterrecoveryconfigs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthorizationRulesDisasterRecoveryConfigsClient: %+v", err)
	}

	return &AuthorizationRulesDisasterRecoveryConfigsClient{
		Client: client,
	}, nil
}
