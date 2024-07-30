package informationprotectionpolicylabel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPolicyLabelClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionPolicyLabelClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionPolicyLabelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionpolicylabel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionPolicyLabelClient: %+v", err)
	}

	return &InformationProtectionPolicyLabelClient{
		Client: client,
	}, nil
}
