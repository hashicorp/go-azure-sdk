package groupsitetermstoresetparentgroupsetrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetRelationClient{
		Client: client,
	}, nil
}
