package groupsitetermstoresetparentgroupsettermrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetTermRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetTermRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetTermRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsettermrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetTermRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetTermRelationClient{
		Client: client,
	}, nil
}
