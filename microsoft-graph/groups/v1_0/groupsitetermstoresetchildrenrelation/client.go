package groupsitetermstoresetchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreSetChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreSetChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreSetChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoresetchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreSetChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreSetChildrenRelationClient{
		Client: client,
	}, nil
}
