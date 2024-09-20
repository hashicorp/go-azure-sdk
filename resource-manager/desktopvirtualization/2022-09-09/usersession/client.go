package usersession

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSessionClient struct {
	Client *resourcemanager.Client
}

func NewUserSessionClientWithBaseURI(sdkApi sdkEnv.Api) (*UserSessionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "usersession", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserSessionClient: %+v", err)
	}

	return &UserSessionClient{
		Client: client,
	}, nil
}
