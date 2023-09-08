package groupsitetermstoresetparentgroupset

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupset", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetClient{
		Client: client,
	}, nil
}
