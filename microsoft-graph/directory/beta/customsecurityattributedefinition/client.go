package customsecurityattributedefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeDefinitionClient struct {
	Client *msgraph.Client
}

func NewCustomSecurityAttributeDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomSecurityAttributeDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "customsecurityattributedefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomSecurityAttributeDefinitionClient: %+v", err)
	}

	return &CustomSecurityAttributeDefinitionClient{
		Client: client,
	}, nil
}
