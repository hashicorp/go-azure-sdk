package groupsitetermstoresetparentgroupsetchildrenchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetchildrenchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient{
		Client: client,
	}, nil
}
