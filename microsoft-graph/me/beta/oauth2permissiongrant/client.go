package oauth2permissiongrant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Oauth2PermissionGrantClient struct {
	Client *msgraph.Client
}

func NewOauth2PermissionGrantClientWithBaseURI(sdkApi sdkEnv.Api) (*Oauth2PermissionGrantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "oauth2permissiongrant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating Oauth2PermissionGrantClient: %+v", err)
	}

	return &Oauth2PermissionGrantClient{
		Client: client,
	}, nil
}
