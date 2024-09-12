package accessreviewdefinitioninstancedecisioninstancedefinition

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessReviewScheduleDefinition
}

type GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions() GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions {
	return GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions{}
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAccessReviewDefinitionInstanceDecisionInstanceDefinition - Get definition from identityGovernance. There's exactly
// one accessReviewScheduleDefinition associated with each instance. It's the parent schedule for the instance, where
// instances are created for each recurrence of a review definition and each group selected to review by the definition.
func (c AccessReviewDefinitionInstanceDecisionInstanceDefinitionClient) GetAccessReviewDefinitionInstanceDecisionInstanceDefinition(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId, options GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationOptions) (result GetAccessReviewDefinitionInstanceDecisionInstanceDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/definition", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	var model beta.AccessReviewScheduleDefinition
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
