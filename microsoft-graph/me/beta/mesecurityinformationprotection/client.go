package mesecurityinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeSecurityInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewMeSecurityInformationProtectionClientWithBaseURI(api sdkEnv.Api) (*MeSecurityInformationProtectionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mesecurityinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeSecurityInformationProtectionClient: %+v", err)
	}

	return &MeSecurityInformationProtectionClient{
		Client: client,
	}, nil
}
