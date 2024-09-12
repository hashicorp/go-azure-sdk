package pendingaccessreviewinstancedecisioninstancestage

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

type StopPendingAccessReviewInstanceDecisionInstanceStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// StopPendingAccessReviewInstanceDecisionInstanceStage - Invoke action stop. Stop an access review stage that is
// inProgress. After the access review stage stops, the stage status will be Completed and the reviewers can no longer
// give input. If there are subsequent stages that depend on the completed stage, the next stage will be created. The
// accessReviewInstanceDecisionItem objects will always reflect the last decisions recorded across all stages at that
// given time, regardless of the status of the stages.
func (c PendingAccessReviewInstanceDecisionInstanceStageClient) StopPendingAccessReviewInstanceDecisionInstanceStage(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionIdInstanceStageId) (result StopPendingAccessReviewInstanceDecisionInstanceStageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/stop", id.ID()),
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
