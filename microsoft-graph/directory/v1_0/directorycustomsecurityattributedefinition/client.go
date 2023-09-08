package directorycustomsecurityattributedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryCustomSecurityAttributeDefinitionClient struct {
	Client *msgraph.Client
}

func NewDirectoryCustomSecurityAttributeDefinitionClientWithBaseURI(api sdkEnv.Api) (*DirectoryCustomSecurityAttributeDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directorycustomsecurityattributedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryCustomSecurityAttributeDefinitionClient: %+v", err)
	}

	return &DirectoryCustomSecurityAttributeDefinitionClient{
		Client: client,
	}, nil
}
