package compliancesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceSettingClient struct {
	Client *msgraph.Client
}

func NewComplianceSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ComplianceSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComplianceSettingClient: %+v", err)
	}

	return &ComplianceSettingClient{
		Client: client,
	}, nil
}
