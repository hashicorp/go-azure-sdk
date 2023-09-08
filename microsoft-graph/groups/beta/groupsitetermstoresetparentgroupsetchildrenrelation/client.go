package groupsitetermstoresetparentgroupsetchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetChildrenRelationClient{
		Client: client,
	}, nil
}
