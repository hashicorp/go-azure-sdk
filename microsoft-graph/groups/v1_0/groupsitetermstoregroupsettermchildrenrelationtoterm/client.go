package groupsitetermstoregroupsettermchildrenrelationtoterm

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteTermStoreGroupSetTermChildrenRelationToTermClient struct {
	Client *msgraph.Client
}

func NewGroupSiteTermStoreGroupSetTermChildrenRelationToTermClientWithBaseURI(api sdkEnv.Api) (*GroupSiteTermStoreGroupSetTermChildrenRelationToTermClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitetermstoregroupsettermchildrenrelationtoterm", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteTermStoreGroupSetTermChildrenRelationToTermClient: %+v", err)
	}

	return &GroupSiteTermStoreGroupSetTermChildrenRelationToTermClient{
		Client: client,
	}, nil
}
