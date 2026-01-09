package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/directoryroletemplates/beta/directoryroletemplate"
	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DirectoryRoleTemplate *directoryroletemplate.DirectoryRoleTemplateClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *msgraph.Client)) (*Client, error) {
	directoryRoleTemplateClient, err := directoryroletemplate.NewDirectoryRoleTemplateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DirectoryRoleTemplate client: %+v", err)
	}
	configureFunc(directoryRoleTemplateClient.Client)

	return &Client{
		DirectoryRoleTemplate: directoryRoleTemplateClient,
	}, nil
}
