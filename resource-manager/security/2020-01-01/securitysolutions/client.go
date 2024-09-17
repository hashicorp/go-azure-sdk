package securitysolutions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySolutionsClient struct {
	Client *resourcemanager.Client
}

func NewSecuritySolutionsClientWithBaseURI(sdkApi sdkEnv.Api) (*SecuritySolutionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "securitysolutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecuritySolutionsClient: %+v", err)
	}

	return &SecuritySolutionsClient{
		Client: client,
	}, nil
}
