package setgroupsitetermstoresetparentgroupsetchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetParentGroupSetChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetParentGroupSetChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetParentGroupSetChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetparentgroupsetchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetParentGroupSetChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetParentGroupSetChildrenRelationClient{
		Client: client,
	}, nil
}
