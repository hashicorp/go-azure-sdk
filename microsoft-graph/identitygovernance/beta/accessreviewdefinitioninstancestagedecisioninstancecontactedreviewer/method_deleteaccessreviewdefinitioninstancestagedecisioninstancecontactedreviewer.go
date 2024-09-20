package accessreviewdefinitioninstancestagedecisioninstancecontactedreviewer

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

type DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultDeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions() DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions {
	return DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions{}
}

func (o DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewer - Delete navigation property
// contactedReviewers for identityGovernance
func (c AccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerClient) DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewer(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId, options DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationOptions) (result DeleteAccessReviewDefinitionInstanceStageDecisionInstanceContactedReviewerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
