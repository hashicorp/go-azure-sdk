package stable

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/oauth2permissiongrants/stable/oauth2permissiongrant"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	OAuth2PermissionGrant *oauth2permissiongrant.OAuth2PermissionGrantClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	oAuth2PermissionGrantClient, err := oauth2permissiongrant.NewOAuth2PermissionGrantClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OAuth2PermissionGrant client: %+v", err)
	}
	configureFunc(oAuth2PermissionGrantClient.Client)

	return &Client{
		OAuth2PermissionGrant: oAuth2PermissionGrantClient,
	}, nil
}
