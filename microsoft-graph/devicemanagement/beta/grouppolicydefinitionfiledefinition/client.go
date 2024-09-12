package grouppolicydefinitionfiledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionFileDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionFileDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionFileDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionfiledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionFileDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionFileDefinitionClient{
		Client: client,
	}, nil
}
