package groupsitetermstoregroupsetchildrenrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetChildrenRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetChildrenRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetChildrenRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsetchildrenrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetChildrenRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetChildrenRelationToTermClient{
		Client: client,
	}, nil
}
