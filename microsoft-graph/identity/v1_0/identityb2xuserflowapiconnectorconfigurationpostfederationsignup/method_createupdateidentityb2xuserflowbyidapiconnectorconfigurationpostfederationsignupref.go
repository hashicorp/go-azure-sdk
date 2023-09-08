package identityb2xuserflowapiconnectorconfigurationpostfederationsignup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef ...
func (c IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient) CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRef(ctx context.Context, id IdentityB2xUserFlowId, input models.ReferenceUpdate) (result CreateUpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup/$ref", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
