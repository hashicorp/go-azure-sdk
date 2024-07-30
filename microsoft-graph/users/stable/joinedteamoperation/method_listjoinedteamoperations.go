package joinedteamoperation

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

type ListJoinedTeamOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.TeamsAsyncOperation
}

type ListJoinedTeamOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.TeamsAsyncOperation
}

type ListJoinedTeamOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamOperations ...
func (c JoinedTeamOperationClient) ListJoinedTeamOperations(ctx context.Context, id UserIdJoinedTeamId) (result ListJoinedTeamOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamOperationsCustomPager{},
		Path:       fmt.Sprintf("%s/operations", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.TeamsAsyncOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamOperationsComplete retrieves all the results into a single object
func (c JoinedTeamOperationClient) ListJoinedTeamOperationsComplete(ctx context.Context, id UserIdJoinedTeamId) (ListJoinedTeamOperationsCompleteResult, error) {
	return c.ListJoinedTeamOperationsCompleteMatchingPredicate(ctx, id, TeamsAsyncOperationOperationPredicate{})
}

// ListJoinedTeamOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamOperationClient) ListJoinedTeamOperationsCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamId, predicate TeamsAsyncOperationOperationPredicate) (result ListJoinedTeamOperationsCompleteResult, err error) {
	items := make([]stable.TeamsAsyncOperation, 0)

	resp, err := c.ListJoinedTeamOperations(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListJoinedTeamOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
