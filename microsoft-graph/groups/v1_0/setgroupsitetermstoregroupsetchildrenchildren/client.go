package setgroupsitetermstoregroupsetchildrenchildren

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetGroupSiteTermStoreGroupSetChildrenChildrenClient struct {
	Client *msgraph.Client
}

func NewSetGroupSiteTermStoreGroupSetChildrenChildrenClientWithBaseURI(api sdkEnv.Api) (*SetGroupSiteTermStoreGroupSetChildrenChildrenClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "setgroupsitetermstoregroupsetchildrenchildren", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SetGroupSiteTermStoreGroupSetChildrenChildrenClient: %+v", err)
	}

	return &SetGroupSiteTermStoreGroupSetChildrenChildrenClient{
		Client: client,
	}, nil
}
