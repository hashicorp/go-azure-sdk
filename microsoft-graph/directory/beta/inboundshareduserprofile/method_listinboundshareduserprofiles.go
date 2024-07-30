package inboundshareduserprofile

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

type ListInboundSharedUserProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.InboundSharedUserProfile
}

type ListInboundSharedUserProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.InboundSharedUserProfile
}

type ListInboundSharedUserProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListInboundSharedUserProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListInboundSharedUserProfiles ...
func (c InboundSharedUserProfileClient) ListInboundSharedUserProfiles(ctx context.Context) (result ListInboundSharedUserProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListInboundSharedUserProfilesCustomPager{},
		Path:       "/directory/inboundSharedUserProfiles",
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
		Values *[]beta.InboundSharedUserProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListInboundSharedUserProfilesComplete retrieves all the results into a single object
func (c InboundSharedUserProfileClient) ListInboundSharedUserProfilesComplete(ctx context.Context) (ListInboundSharedUserProfilesCompleteResult, error) {
	return c.ListInboundSharedUserProfilesCompleteMatchingPredicate(ctx, InboundSharedUserProfileOperationPredicate{})
}

// ListInboundSharedUserProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c InboundSharedUserProfileClient) ListInboundSharedUserProfilesCompleteMatchingPredicate(ctx context.Context, predicate InboundSharedUserProfileOperationPredicate) (result ListInboundSharedUserProfilesCompleteResult, err error) {
	items := make([]beta.InboundSharedUserProfile, 0)

	resp, err := c.ListInboundSharedUserProfiles(ctx)
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

	result = ListInboundSharedUserProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
