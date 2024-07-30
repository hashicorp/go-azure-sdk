package chatpermissiongrant

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

type ListChatPermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ResourceSpecificPermissionGrant
}

type ListChatPermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ResourceSpecificPermissionGrant
}

type ListChatPermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListChatPermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListChatPermissionGrants ...
func (c ChatPermissionGrantClient) ListChatPermissionGrants(ctx context.Context, id UserIdChatId) (result ListChatPermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListChatPermissionGrantsCustomPager{},
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
		Values *[]beta.ResourceSpecificPermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListChatPermissionGrantsComplete retrieves all the results into a single object
func (c ChatPermissionGrantClient) ListChatPermissionGrantsComplete(ctx context.Context, id UserIdChatId) (ListChatPermissionGrantsCompleteResult, error) {
	return c.ListChatPermissionGrantsCompleteMatchingPredicate(ctx, id, ResourceSpecificPermissionGrantOperationPredicate{})
}

// ListChatPermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ChatPermissionGrantClient) ListChatPermissionGrantsCompleteMatchingPredicate(ctx context.Context, id UserIdChatId, predicate ResourceSpecificPermissionGrantOperationPredicate) (result ListChatPermissionGrantsCompleteResult, err error) {
	items := make([]beta.ResourceSpecificPermissionGrant, 0)

	resp, err := c.ListChatPermissionGrants(ctx, id)
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

	result = ListChatPermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
