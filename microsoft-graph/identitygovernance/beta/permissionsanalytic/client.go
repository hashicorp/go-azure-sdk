package permissionsanalytic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalyticClient struct {
	Client *msgraph.Client
}

func NewPermissionsAnalyticClientWithBaseURI(sdkApi sdkEnv.Api) (*PermissionsAnalyticClient, error) {
	client, err := msgraph.NewClient(sdkApi, "permissionsanalytic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PermissionsAnalyticClient: %+v", err)
	}

	return &PermissionsAnalyticClient{
		Client: client,
	}, nil
}
