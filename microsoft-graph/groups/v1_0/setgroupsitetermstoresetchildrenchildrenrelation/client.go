package setgroupsitetermstoresetchildrenchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetChildrenChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetChildrenChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetChildrenChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresetchildrenchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetChildrenChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetChildrenChildrenRelationClient{
		Client: client,
	}, nil
}
