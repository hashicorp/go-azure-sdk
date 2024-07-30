package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/oauth2permissiongrants/beta/oauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Oauth2PermissionGrant *oauth2permissiongrant.Oauth2PermissionGrantClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	oauth2PermissionGrantClient, err := oauth2permissiongrant.NewOauth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Oauth2PermissionGrant client: %+v", err)
	}
	configureFunc(oauth2PermissionGrantClient.Client)

	return &Client{
		Oauth2PermissionGrant: oauth2PermissionGrantClient,
	}, nil
}
