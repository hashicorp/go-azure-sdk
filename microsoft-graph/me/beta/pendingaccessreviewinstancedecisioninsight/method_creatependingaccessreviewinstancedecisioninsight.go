package pendingaccessreviewinstancedecisioninsight

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePendingAccessReviewInstanceDecisionInsightOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.GovernanceInsight
}

type CreatePendingAccessReviewInstanceDecisionInsightOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreatePendingAccessReviewInstanceDecisionInsightOperationOptions() CreatePendingAccessReviewInstanceDecisionInsightOperationOptions {
	return CreatePendingAccessReviewInstanceDecisionInsightOperationOptions{}
}

func (o CreatePendingAccessReviewInstanceDecisionInsightOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInsightOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePendingAccessReviewInstanceDecisionInsightOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePendingAccessReviewInstanceDecisionInsight - Create new navigation property to insights for me
func (c PendingAccessReviewInstanceDecisionInsightClient) CreatePendingAccessReviewInstanceDecisionInsight(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionId, input beta.GovernanceInsight, options CreatePendingAccessReviewInstanceDecisionInsightOperationOptions) (result CreatePendingAccessReviewInstanceDecisionInsightOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/insights", id.ID()),
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
	model, err := beta.UnmarshalGovernanceInsightImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
