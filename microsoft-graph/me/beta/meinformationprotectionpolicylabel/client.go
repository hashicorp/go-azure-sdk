package meinformationprotectionpolicylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeInformationProtectionPolicyLabelClient struct {
	Client *msgraph.Client
}

func NewMeInformationProtectionPolicyLabelClientWithBaseURI(api sdkEnv.Api) (*MeInformationProtectionPolicyLabelClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meinformationprotectionpolicylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeInformationProtectionPolicyLabelClient: %+v", err)
	}

	return &MeInformationProtectionPolicyLabelClient{
		Client: client,
	}, nil
}
