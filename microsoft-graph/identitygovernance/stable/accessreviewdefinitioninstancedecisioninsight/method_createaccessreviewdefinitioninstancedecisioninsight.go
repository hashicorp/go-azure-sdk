package accessreviewdefinitioninstancedecisioninsight

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateAccessReviewDefinitionInstanceDecisionInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.GovernanceInsight
}

type CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions() CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions {
	return CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions{}
}

func (o CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateAccessReviewDefinitionInstanceDecisionInsight - Create new navigation property to insights for
// identityGovernance
func (c AccessReviewDefinitionInstanceDecisionInsightClient) CreateAccessReviewDefinitionInstanceDecisionInsight(ctx context.Context, id stable.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId, input stable.GovernanceInsight, options CreateAccessReviewDefinitionInstanceDecisionInsightOperationOptions) (result CreateAccessReviewDefinitionInstanceDecisionInsightOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/insights", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalGovernanceInsightImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
