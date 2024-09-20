package externalsecuritysolutions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalSecuritySolutionsClient struct {
	Client *resourcemanager.Client
}

func NewExternalSecuritySolutionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ExternalSecuritySolutionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "externalsecuritysolutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ExternalSecuritySolutionsClient: %+v", err)
	}

	return &ExternalSecuritySolutionsClient{
		Client: client,
	}, nil
}
