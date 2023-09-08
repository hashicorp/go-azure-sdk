package setgroupsitetermstoregroupsetchildrenrelation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetChildrenRelationClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetChildrenRelationClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetChildrenRelationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsetchildrenrelation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetChildrenRelationClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetChildrenRelationClient{
		Client: client,
	}, nil
}
