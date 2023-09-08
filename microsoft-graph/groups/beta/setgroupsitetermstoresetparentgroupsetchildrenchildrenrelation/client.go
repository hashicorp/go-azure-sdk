package setgroupsitetermstoresetparentgroupsetchildrenchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetparentgroupsetchildrenchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient{
		Client: client,
	}, nil
}
