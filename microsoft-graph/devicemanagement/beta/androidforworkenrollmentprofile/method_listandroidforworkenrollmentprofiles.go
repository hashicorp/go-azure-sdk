package androidforworkenrollmentprofile

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

type ListAndroidForWorkEnrollmentProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AndroidForWorkEnrollmentProfile
}

type ListAndroidForWorkEnrollmentProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AndroidForWorkEnrollmentProfile
}

type ListAndroidForWorkEnrollmentProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAndroidForWorkEnrollmentProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAndroidForWorkEnrollmentProfiles ...
func (c AndroidForWorkEnrollmentProfileClient) ListAndroidForWorkEnrollmentProfiles(ctx context.Context) (result ListAndroidForWorkEnrollmentProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAndroidForWorkEnrollmentProfilesCustomPager{},
		Path:       "/deviceManagement/androidForWorkEnrollmentProfiles",
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
		Values *[]beta.AndroidForWorkEnrollmentProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAndroidForWorkEnrollmentProfilesComplete retrieves all the results into a single object
func (c AndroidForWorkEnrollmentProfileClient) ListAndroidForWorkEnrollmentProfilesComplete(ctx context.Context) (ListAndroidForWorkEnrollmentProfilesCompleteResult, error) {
	return c.ListAndroidForWorkEnrollmentProfilesCompleteMatchingPredicate(ctx, AndroidForWorkEnrollmentProfileOperationPredicate{})
}

// ListAndroidForWorkEnrollmentProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AndroidForWorkEnrollmentProfileClient) ListAndroidForWorkEnrollmentProfilesCompleteMatchingPredicate(ctx context.Context, predicate AndroidForWorkEnrollmentProfileOperationPredicate) (result ListAndroidForWorkEnrollmentProfilesCompleteResult, err error) {
	items := make([]beta.AndroidForWorkEnrollmentProfile, 0)

	resp, err := c.ListAndroidForWorkEnrollmentProfiles(ctx)
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

	result = ListAndroidForWorkEnrollmentProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
