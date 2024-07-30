package grouppolicycategorychildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryChildrenClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyCategoryChildrenClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyCategoryChildrenClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicycategorychildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyCategoryChildrenClient: %+v", err)
	}

	return &GroupPolicyCategoryChildrenClient{
		Client: client,
	}, nil
}
