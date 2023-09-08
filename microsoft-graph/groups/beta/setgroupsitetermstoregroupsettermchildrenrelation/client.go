package setgroupsitetermstoregroupsettermchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetTermChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetTermChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetTermChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsettermchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetTermChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetTermChildrenRelationClient{
		Client: client,
	}, nil
}
