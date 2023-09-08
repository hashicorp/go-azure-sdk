package groupsitetermstoreset

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoreset", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetClient: %+v", err)
	}

	return &GroupSiteTermStoreSetClient{
		Client: client,
	}, nil
}
