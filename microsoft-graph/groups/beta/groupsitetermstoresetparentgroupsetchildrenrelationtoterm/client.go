package groupsitetermstoresetparentgroupsetchildrenrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetparentgroupsetchildrenrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreSetParentGroupSetChildrenRelationToTermClient{
		Client: client,
	}, nil
}
