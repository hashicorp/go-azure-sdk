package groupsitetermstoresetparentgroupsetchildrenchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetChildrenChildrenClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetChildrenChildrenClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetChildrenChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetchildrenchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetChildrenChildrenClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetChildrenChildrenClient{
		Client: client,
	}, nil
}
