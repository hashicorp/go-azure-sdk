package grouppolicydefinitionpresentation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPresentationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPresentationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPresentationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpresentation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPresentationClient: %+v", err)
	}

	return &GroupPolicyDefinitionPresentationClient{
		Client: client,
	}, nil
}
