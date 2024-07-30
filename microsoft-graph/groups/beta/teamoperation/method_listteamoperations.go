package teamoperation

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

type ListTeamOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.TeamsAsyncOperation
}

type ListTeamOperationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.TeamsAsyncOperation
}

type ListTeamOperationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamOperationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamOperations ...
func (c TeamOperationClient) ListTeamOperations(ctx context.Context, id GroupId) (result ListTeamOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamOperationsCustomPager{},
		Path:       fmt.Sprintf("%s/team/operations", id.ID()),
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
		Values *[]beta.TeamsAsyncOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamOperationsComplete retrieves all the results into a single object
func (c TeamOperationClient) ListTeamOperationsComplete(ctx context.Context, id GroupId) (ListTeamOperationsCompleteResult, error) {
	return c.ListTeamOperationsCompleteMatchingPredicate(ctx, id, TeamsAsyncOperationOperationPredicate{})
}

// ListTeamOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamOperationClient) ListTeamOperationsCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate TeamsAsyncOperationOperationPredicate) (result ListTeamOperationsCompleteResult, err error) {
	items := make([]beta.TeamsAsyncOperation, 0)

	resp, err := c.ListTeamOperations(ctx, id)
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

	result = ListTeamOperationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
