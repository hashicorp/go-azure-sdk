package grouppolicycategorychild

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryChildClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyCategoryChildClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyCategoryChildClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicycategorychild", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyCategoryChildClient: %+v", err)
	}

	return &GroupPolicyCategoryChildClient{
		Client: client,
	}, nil
}
