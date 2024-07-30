package intunebrandingprofile

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

type ListIntuneBrandingProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.IntuneBrandingProfile
}

type ListIntuneBrandingProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.IntuneBrandingProfile
}

type ListIntuneBrandingProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListIntuneBrandingProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListIntuneBrandingProfiles ...
func (c IntuneBrandingProfileClient) ListIntuneBrandingProfiles(ctx context.Context) (result ListIntuneBrandingProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListIntuneBrandingProfilesCustomPager{},
		Path:       "/deviceManagement/intuneBrandingProfiles",
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
		Values *[]beta.IntuneBrandingProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListIntuneBrandingProfilesComplete retrieves all the results into a single object
func (c IntuneBrandingProfileClient) ListIntuneBrandingProfilesComplete(ctx context.Context) (ListIntuneBrandingProfilesCompleteResult, error) {
	return c.ListIntuneBrandingProfilesCompleteMatchingPredicate(ctx, IntuneBrandingProfileOperationPredicate{})
}

// ListIntuneBrandingProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IntuneBrandingProfileClient) ListIntuneBrandingProfilesCompleteMatchingPredicate(ctx context.Context, predicate IntuneBrandingProfileOperationPredicate) (result ListIntuneBrandingProfilesCompleteResult, err error) {
	items := make([]beta.IntuneBrandingProfile, 0)

	resp, err := c.ListIntuneBrandingProfiles(ctx)
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

	result = ListIntuneBrandingProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
