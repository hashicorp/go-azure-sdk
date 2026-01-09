package informationprotectionsensitivitypolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionSensitivityPolicySettingClient struct {
	Client *msgraph.Client
}

func NewInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionSensitivityPolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "informationprotectionsensitivitypolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionSensitivityPolicySettingClient: %+v", err)
	}

	return &InformationProtectionSensitivityPolicySettingClient{
		Client: client,
	}, nil
}
