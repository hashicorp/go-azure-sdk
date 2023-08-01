package topleveldomains

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TopLevelDomainsClient struct {
	Client *resourcemanager.Client
}

func NewTopLevelDomainsClientWithBaseURI(sdkApi sdkEnv.Api) (*TopLevelDomainsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "topleveldomains", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TopLevelDomainsClient: %+v", err)
	}

	return &TopLevelDomainsClient{
		Client: client,
	}, nil
}
