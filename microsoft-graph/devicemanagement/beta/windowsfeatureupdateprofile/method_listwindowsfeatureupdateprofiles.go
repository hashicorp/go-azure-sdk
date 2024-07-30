package windowsfeatureupdateprofile

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

type ListWindowsFeatureUpdateProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsFeatureUpdateProfile
}

type ListWindowsFeatureUpdateProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsFeatureUpdateProfile
}

type ListWindowsFeatureUpdateProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsFeatureUpdateProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsFeatureUpdateProfiles ...
func (c WindowsFeatureUpdateProfileClient) ListWindowsFeatureUpdateProfiles(ctx context.Context) (result ListWindowsFeatureUpdateProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsFeatureUpdateProfilesCustomPager{},
		Path:       "/deviceManagement/windowsFeatureUpdateProfiles",
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
		Values *[]beta.WindowsFeatureUpdateProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsFeatureUpdateProfilesComplete retrieves all the results into a single object
func (c WindowsFeatureUpdateProfileClient) ListWindowsFeatureUpdateProfilesComplete(ctx context.Context) (ListWindowsFeatureUpdateProfilesCompleteResult, error) {
	return c.ListWindowsFeatureUpdateProfilesCompleteMatchingPredicate(ctx, WindowsFeatureUpdateProfileOperationPredicate{})
}

// ListWindowsFeatureUpdateProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsFeatureUpdateProfileClient) ListWindowsFeatureUpdateProfilesCompleteMatchingPredicate(ctx context.Context, predicate WindowsFeatureUpdateProfileOperationPredicate) (result ListWindowsFeatureUpdateProfilesCompleteResult, err error) {
	items := make([]beta.WindowsFeatureUpdateProfile, 0)

	resp, err := c.ListWindowsFeatureUpdateProfiles(ctx)
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

	result = ListWindowsFeatureUpdateProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
