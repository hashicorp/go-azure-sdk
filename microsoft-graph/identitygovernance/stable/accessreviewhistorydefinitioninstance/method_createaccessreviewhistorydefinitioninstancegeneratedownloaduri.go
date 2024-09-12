package accessreviewhistorydefinitioninstance

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

type CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessReviewHistoryInstance
}

// CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUri - Invoke action generateDownloadUri. Generates a URI
// for an accessReviewHistoryInstance object the status for which is done. Each URI can be used to retrieve the
// instance's review history data. Each URI is valid for 24 hours and can be retrieved by fetching the downloadUri
// property from the accessReviewHistoryInstance object.
func (c AccessReviewHistoryDefinitionInstanceClient) CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUri(ctx context.Context, id stable.IdentityGovernanceAccessReviewHistoryDefinitionIdInstanceId) (result CreateAccessReviewHistoryDefinitionInstanceGenerateDownloadUriOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/generateDownloadUri", id.ID()),
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

	var model stable.AccessReviewHistoryInstance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
