package siteinformationprotectionpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionPolicyClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionPolicyClient: %+v", err)
	}

	return &SiteInformationProtectionPolicyClient{
		Client: client,
	}, nil
}
