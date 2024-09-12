package grouppolicycategorydefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyCategoryDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyCategoryDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicycategorydefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyCategoryDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyCategoryDefinitionFileClient{
		Client: client,
	}, nil
}
