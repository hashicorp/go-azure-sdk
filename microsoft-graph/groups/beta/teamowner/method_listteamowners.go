package teamowner

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

type ListTeamOwnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.User
}

type ListTeamOwnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.User
}

type ListTeamOwnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTeamOwnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTeamOwners ...
func (c TeamOwnerClient) ListTeamOwners(ctx context.Context, id GroupId) (result ListTeamOwnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTeamOwnersCustomPager{},
		Path:       fmt.Sprintf("%s/team/owners", id.ID()),
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
		Values *[]beta.User `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTeamOwnersComplete retrieves all the results into a single object
func (c TeamOwnerClient) ListTeamOwnersComplete(ctx context.Context, id GroupId) (ListTeamOwnersCompleteResult, error) {
	return c.ListTeamOwnersCompleteMatchingPredicate(ctx, id, UserOperationPredicate{})
}

// ListTeamOwnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TeamOwnerClient) ListTeamOwnersCompleteMatchingPredicate(ctx context.Context, id GroupId, predicate UserOperationPredicate) (result ListTeamOwnersCompleteResult, err error) {
	items := make([]beta.User, 0)

	resp, err := c.ListTeamOwners(ctx, id)
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

	result = ListTeamOwnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
