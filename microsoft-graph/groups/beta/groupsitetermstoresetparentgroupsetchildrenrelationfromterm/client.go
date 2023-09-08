package groupsitetermstoresetparentgroupsetchildrenrelationfromterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetchildrenrelationfromterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetChildrenRelationFromTermClient{
		Client: client,
	}, nil
}
