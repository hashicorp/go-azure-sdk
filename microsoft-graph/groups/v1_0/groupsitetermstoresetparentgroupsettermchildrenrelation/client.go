package groupsitetermstoresetparentgroupsettermchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetTermChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsettermchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient{
		Client: client,
	}, nil
}
