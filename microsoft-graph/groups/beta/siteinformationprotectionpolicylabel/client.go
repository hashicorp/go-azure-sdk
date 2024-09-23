package siteinformationprotectionpolicylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionPolicyLabelClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionPolicyLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionPolicyLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionpolicylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionPolicyLabelClient: %+v", err)
	}

	return &SiteInformationProtectionPolicyLabelClient{
		Client: client,
	}, nil
}
