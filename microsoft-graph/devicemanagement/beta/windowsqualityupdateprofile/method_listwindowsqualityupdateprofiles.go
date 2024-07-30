package windowsqualityupdateprofile

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

type ListWindowsQualityUpdateProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.WindowsQualityUpdateProfile
}

type ListWindowsQualityUpdateProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.WindowsQualityUpdateProfile
}

type ListWindowsQualityUpdateProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListWindowsQualityUpdateProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListWindowsQualityUpdateProfiles ...
func (c WindowsQualityUpdateProfileClient) ListWindowsQualityUpdateProfiles(ctx context.Context) (result ListWindowsQualityUpdateProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListWindowsQualityUpdateProfilesCustomPager{},
		Path:       "/deviceManagement/windowsQualityUpdateProfiles",
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
		Values *[]beta.WindowsQualityUpdateProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListWindowsQualityUpdateProfilesComplete retrieves all the results into a single object
func (c WindowsQualityUpdateProfileClient) ListWindowsQualityUpdateProfilesComplete(ctx context.Context) (ListWindowsQualityUpdateProfilesCompleteResult, error) {
	return c.ListWindowsQualityUpdateProfilesCompleteMatchingPredicate(ctx, WindowsQualityUpdateProfileOperationPredicate{})
}

// ListWindowsQualityUpdateProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WindowsQualityUpdateProfileClient) ListWindowsQualityUpdateProfilesCompleteMatchingPredicate(ctx context.Context, predicate WindowsQualityUpdateProfileOperationPredicate) (result ListWindowsQualityUpdateProfilesCompleteResult, err error) {
	items := make([]beta.WindowsQualityUpdateProfile, 0)

	resp, err := c.ListWindowsQualityUpdateProfiles(ctx)
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

	result = ListWindowsQualityUpdateProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
