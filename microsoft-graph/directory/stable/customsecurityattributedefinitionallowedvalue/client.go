package customsecurityattributedefinitionallowedvalue

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeDefinitionAllowedValueClient struct {
	Client *msgraph.Client
}

func NewCustomSecurityAttributeDefinitionAllowedValueClientWithBaseURI(sdkApi sdkEnv.Api) (*CustomSecurityAttributeDefinitionAllowedValueClient, error) {
	client, err := msgraph.NewClient(sdkApi, "customsecurityattributedefinitionallowedvalue", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CustomSecurityAttributeDefinitionAllowedValueClient: %+v", err)
	}

	return &CustomSecurityAttributeDefinitionAllowedValueClient{
		Client: client,
	}, nil
}
