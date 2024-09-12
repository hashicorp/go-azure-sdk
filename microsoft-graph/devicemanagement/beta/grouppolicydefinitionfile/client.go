package grouppolicydefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyDefinitionFileClient{
		Client: client,
	}, nil
}
