package groupsitetermstoregroupset

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupset", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetClient{
		Client: client,
	}, nil
}
