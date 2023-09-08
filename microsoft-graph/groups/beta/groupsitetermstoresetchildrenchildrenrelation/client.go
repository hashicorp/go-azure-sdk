package groupsitetermstoresetchildrenchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildrenchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenChildrenRelationClient{
		Client: client,
	}, nil
}
