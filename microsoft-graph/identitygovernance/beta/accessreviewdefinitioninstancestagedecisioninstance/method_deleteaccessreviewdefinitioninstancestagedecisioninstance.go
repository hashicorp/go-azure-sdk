package accessreviewdefinitioninstancestagedecisioninstance

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

type DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions() DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions {
	return DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions{}
}

func (o DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAccessReviewDefinitionInstanceStageDecisionInstance - Delete navigation property instance for
// identityGovernance
func (c AccessReviewDefinitionInstanceStageDecisionInstanceClient) DeleteAccessReviewDefinitionInstanceStageDecisionInstance(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, options DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationOptions) (result DeleteAccessReviewDefinitionInstanceStageDecisionInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance", id.ID()),
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

	return
}
