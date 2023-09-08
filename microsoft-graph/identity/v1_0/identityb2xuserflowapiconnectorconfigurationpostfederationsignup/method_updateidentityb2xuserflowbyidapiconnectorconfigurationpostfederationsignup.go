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

type UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *models.IdentityApiConnector
}

// UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup ...
func (c IdentityB2xUserFlowApiConnectorConfigurationPostFederationSignupClient) UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignup(ctx context.Context, id IdentityB2xUserFlowId, input models.IdentityApiConnector) (result UpdateIdentityB2xUserFlowByIdApiConnectorConfigurationPostFederationSignupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPatch,
		Path:       fmt.Sprintf("%s/apiConnectorConfiguration/postFederationSignup", id.ID()),
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
