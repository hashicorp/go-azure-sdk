package groupsitetermstoresetparentgroupsettermrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetTermRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetTermRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetTermRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsettermrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetTermRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetTermRelationToTermClient{
		Client: client,
	}, nil
}
