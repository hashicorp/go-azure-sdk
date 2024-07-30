package windowsdriverupdateprofile

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

type ListWindowsDriverUpdateProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsDriverUpdateProfile
}

type ListWindowsDriverUpdateProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsDriverUpdateProfile
}

type ListWindowsDriverUpdateProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsDriverUpdateProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsDriverUpdateProfiles ...
func (c WindowsDriverUpdateProfileClient) ListWindowsDriverUpdateProfiles(ctx context.Context) (result ListWindowsDriverUpdateProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsDriverUpdateProfilesCustomPager{},
		Path:       "/deviceManagement/windowsDriverUpdateProfiles",
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
		Values *[]beta.WindowsDriverUpdateProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsDriverUpdateProfilesComplete retrieves all the results into a single object
func (c WindowsDriverUpdateProfileClient) ListWindowsDriverUpdateProfilesComplete(ctx context.Context) (ListWindowsDriverUpdateProfilesCompleteResult, error) {
	return c.ListWindowsDriverUpdateProfilesCompleteMatchingPredicate(ctx, WindowsDriverUpdateProfileOperationPredicate{})
}

// ListWindowsDriverUpdateProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsDriverUpdateProfileClient) ListWindowsDriverUpdateProfilesCompleteMatchingPredicate(ctx context.Context, predicate WindowsDriverUpdateProfileOperationPredicate) (result ListWindowsDriverUpdateProfilesCompleteResult, err error) {
	items := make([]beta.WindowsDriverUpdateProfile, 0)

	resp, err := c.ListWindowsDriverUpdateProfiles(ctx)
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

	result = ListWindowsDriverUpdateProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
