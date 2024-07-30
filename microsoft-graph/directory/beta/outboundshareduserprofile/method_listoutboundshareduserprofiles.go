package outboundshareduserprofile

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

type ListOutboundSharedUserProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.OutboundSharedUserProfile
}

type ListOutboundSharedUserProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.OutboundSharedUserProfile
}

type ListOutboundSharedUserProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutboundSharedUserProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutboundSharedUserProfiles ...
func (c OutboundSharedUserProfileClient) ListOutboundSharedUserProfiles(ctx context.Context) (result ListOutboundSharedUserProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOutboundSharedUserProfilesCustomPager{},
		Path:       "/directory/outboundSharedUserProfiles",
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
		Values *[]beta.OutboundSharedUserProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutboundSharedUserProfilesComplete retrieves all the results into a single object
func (c OutboundSharedUserProfileClient) ListOutboundSharedUserProfilesComplete(ctx context.Context) (ListOutboundSharedUserProfilesCompleteResult, error) {
	return c.ListOutboundSharedUserProfilesCompleteMatchingPredicate(ctx, OutboundSharedUserProfileOperationPredicate{})
}

// ListOutboundSharedUserProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutboundSharedUserProfileClient) ListOutboundSharedUserProfilesCompleteMatchingPredicate(ctx context.Context, predicate OutboundSharedUserProfileOperationPredicate) (result ListOutboundSharedUserProfilesCompleteResult, err error) {
	items := make([]beta.OutboundSharedUserProfile, 0)

	resp, err := c.ListOutboundSharedUserProfiles(ctx)
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

	result = ListOutboundSharedUserProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
