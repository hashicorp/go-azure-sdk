package groupsitetermstoregroupsetchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetChildrenRelationClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetChildrenRelationClient{
		Client: client,
	}, nil
}
