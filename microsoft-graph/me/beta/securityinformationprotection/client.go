package securityinformationprotection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtectionClient struct {
	Client *msgraph.Client
}

func NewSecurityInformationProtectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SecurityInformationProtectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "securityinformationprotection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SecurityInformationProtectionClient: %+v", err)
	}

	return &SecurityInformationProtectionClient{
		Client: client,
	}, nil
}
