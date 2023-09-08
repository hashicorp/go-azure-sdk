package groupsitetermstoresetparentgroupsettermchildrenrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsettermchildrenrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetTermChildrenRelationFromTermClient{
		Client: client,
	}, nil
}
