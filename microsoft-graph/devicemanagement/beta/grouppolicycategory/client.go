package grouppolicycategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyCategoryClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicycategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyCategoryClient: %+v", err)
	}

	return &GroupPolicyCategoryClient{
		Client: client,
	}, nil
}
