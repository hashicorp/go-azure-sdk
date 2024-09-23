package siteinformationprotectionsensitivitypolicysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteInformationProtectionSensitivityPolicySettingClient struct {
	Client *msgraph.Client
}

func NewSiteInformationProtectionSensitivityPolicySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteInformationProtectionSensitivityPolicySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteinformationprotectionsensitivitypolicysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteInformationProtectionSensitivityPolicySettingClient: %+v", err)
	}

	return &SiteInformationProtectionSensitivityPolicySettingClient{
		Client: client,
	}, nil
}
