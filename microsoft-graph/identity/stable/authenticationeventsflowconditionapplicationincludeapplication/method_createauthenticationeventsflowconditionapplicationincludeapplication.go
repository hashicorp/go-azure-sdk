package authenticationeventsflowconditionapplicationincludeapplication

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAuthenticationEventsFlowConditionApplicationIncludeApplicationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AuthenticationConditionApplication
}

// CreateAuthenticationEventsFlowConditionApplicationIncludeApplication - Add includeApplication (to a user flow). Add
// or link an application to a user flow, or authenticationEventsFlow. This enables the authentication experience
// defined by the user flow to be enabled for the application. An application can only be linked to one user flow.
func (c AuthenticationEventsFlowConditionApplicationIncludeApplicationClient) CreateAuthenticationEventsFlowConditionApplicationIncludeApplication(ctx context.Context, id stable.IdentityAuthenticationEventsFlowId, input stable.AuthenticationConditionApplication) (result CreateAuthenticationEventsFlowConditionApplicationIncludeApplicationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/conditions/applications/includeApplications", id.ID()),
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

	var model stable.AuthenticationConditionApplication
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
