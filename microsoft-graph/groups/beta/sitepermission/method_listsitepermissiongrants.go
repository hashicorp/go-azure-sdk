package sitepermission

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

type ListSitePermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Permission
}

type ListSitePermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Permission
}

type ListSitePermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListSitePermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSitePermissionGrants ...
func (c SitePermissionClient) ListSitePermissionGrants(ctx context.Context, id GroupIdSiteIdPermissionId, input ListSitePermissionGrantsRequest) (result ListSitePermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ListSitePermissionGrantsCustomPager{},
		Path:       fmt.Sprintf("%s/grant", id.ID()),
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
		Values *[]beta.Permission `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSitePermissionGrantsComplete retrieves all the results into a single object
func (c SitePermissionClient) ListSitePermissionGrantsComplete(ctx context.Context, id GroupIdSiteIdPermissionId, input ListSitePermissionGrantsRequest) (ListSitePermissionGrantsCompleteResult, error) {
	return c.ListSitePermissionGrantsCompleteMatchingPredicate(ctx, id, input, PermissionOperationPredicate{})
}

// ListSitePermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitePermissionClient) ListSitePermissionGrantsCompleteMatchingPredicate(ctx context.Context, id GroupIdSiteIdPermissionId, input ListSitePermissionGrantsRequest, predicate PermissionOperationPredicate) (result ListSitePermissionGrantsCompleteResult, err error) {
	items := make([]beta.Permission, 0)

	resp, err := c.ListSitePermissionGrants(ctx, id, input)
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

	result = ListSitePermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
