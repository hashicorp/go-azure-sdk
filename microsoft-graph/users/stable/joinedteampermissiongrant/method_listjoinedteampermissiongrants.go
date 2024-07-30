package joinedteampermissiongrant

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

type ListJoinedTeamPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.ResourceSpecificPermissionGrant
}

type ListJoinedTeamPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.ResourceSpecificPermissionGrant
}

type ListJoinedTeamPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListJoinedTeamPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListJoinedTeamPermissionGrants ...
func (c JoinedTeamPermissionGrantClient) ListJoinedTeamPermissionGrants(ctx context.Context, id UserIdJoinedTeamId) (result ListJoinedTeamPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListJoinedTeamPermissionGrantsCustomPager{},
		Path:       fmt.Sprintf("%s/permissionGrants", id.ID()),
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
		Values *[]stable.ResourceSpecificPermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListJoinedTeamPermissionGrantsComplete retrieves all the results into a single object
func (c JoinedTeamPermissionGrantClient) ListJoinedTeamPermissionGrantsComplete(ctx context.Context, id UserIdJoinedTeamId) (ListJoinedTeamPermissionGrantsCompleteResult, error) {
	return c.ListJoinedTeamPermissionGrantsCompleteMatchingPredicate(ctx, id, ResourceSpecificPermissionGrantOperationPredicate{})
}

// ListJoinedTeamPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c JoinedTeamPermissionGrantClient) ListJoinedTeamPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id UserIdJoinedTeamId, predicate ResourceSpecificPermissionGrantOperationPredicate) (result ListJoinedTeamPermissionGrantsCompleteResult, err error) {
	items := make([]stable.ResourceSpecificPermissionGrant, 0)

	resp, err := c.ListJoinedTeamPermissionGrants(ctx, id)
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

	result = ListJoinedTeamPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
