package accessreviewdefinitioninstancedecisioninstancestage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultStopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions() StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions {
	return StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions{}
}

func (o StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StopAccessReviewDefinitionInstanceDecisionInstanceStage - Invoke action stop. Stop an access review stage that is
// inProgress. After the access review stage stops, the stage status will be Completed and the reviewers can no longer
// give input. If there are subsequent stages that depend on the completed stage, the next stage will be created. The
// accessReviewInstanceDecisionItem objects will always reflect the last decisions recorded across all stages at that
// given time, regardless of the status of the stages.
func (c AccessReviewDefinitionInstanceDecisionInstanceStageClient) StopAccessReviewDefinitionInstanceDecisionInstanceStage(ctx context.Context, id beta.IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId, options StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationOptions) (result StopAccessReviewDefinitionInstanceDecisionInstanceStageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/stop", id.ID()),
		RetryFunc:     options.RetryFunc,
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
