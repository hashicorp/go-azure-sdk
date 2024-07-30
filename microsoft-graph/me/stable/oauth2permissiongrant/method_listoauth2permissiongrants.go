package oauth2permissiongrant

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

type ListOauth2PermissionGrantsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.OAuth2PermissionGrant
}

type ListOauth2PermissionGrantsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.OAuth2PermissionGrant
}

type ListOauth2PermissionGrantsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOauth2PermissionGrantsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOauth2PermissionGrants ...
func (c Oauth2PermissionGrantClient) ListOauth2PermissionGrants(ctx context.Context) (result ListOauth2PermissionGrantsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOauth2PermissionGrantsCustomPager{},
		Path:       "/me/oauth2PermissionGrants",
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
		Values *[]stable.OAuth2PermissionGrant `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOauth2PermissionGrantsComplete retrieves all the results into a single object
func (c Oauth2PermissionGrantClient) ListOauth2PermissionGrantsComplete(ctx context.Context) (ListOauth2PermissionGrantsCompleteResult, error) {
	return c.ListOauth2PermissionGrantsCompleteMatchingPredicate(ctx, OAuth2PermissionGrantOperationPredicate{})
}

// ListOauth2PermissionGrantsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c Oauth2PermissionGrantClient) ListOauth2PermissionGrantsCompleteMatchingPredicate(ctx context.Context, predicate OAuth2PermissionGrantOperationPredicate) (result ListOauth2PermissionGrantsCompleteResult, err error) {
	items := make([]stable.OAuth2PermissionGrant, 0)

	resp, err := c.ListOauth2PermissionGrants(ctx)
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

	result = ListOauth2PermissionGrantsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
