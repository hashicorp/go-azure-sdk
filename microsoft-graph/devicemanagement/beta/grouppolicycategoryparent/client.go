package grouppolicycategoryparent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryParentClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyCategoryParentClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyCategoryParentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicycategoryparent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyCategoryParentClient: %+v", err)
	}

	return &GroupPolicyCategoryParentClient{
		Client: client,
	}, nil
}
