package permissionsanalyticazure

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticAzureClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticAzureClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticAzureClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalyticazure", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticAzureClient: %+v", err)
	}

	return &PermissionsAnalyticAzureClient{
		Client: client,
	}, nil
}
