package setgroupsitetermstoresettermchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreSetTermChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreSetTermChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreSetTermChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoresettermchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreSetTermChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreSetTermChildrenRelationClient{
		Client: client,
	}, nil
}
