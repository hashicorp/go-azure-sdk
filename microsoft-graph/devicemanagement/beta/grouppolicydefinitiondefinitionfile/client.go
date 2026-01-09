package grouppolicydefinitiondefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitiondefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyDefinitionDefinitionFileClient{
		Client: client,
	}, nil
}
