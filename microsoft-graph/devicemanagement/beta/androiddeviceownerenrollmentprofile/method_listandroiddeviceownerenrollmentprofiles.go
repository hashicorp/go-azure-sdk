package androiddeviceownerenrollmentprofile

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

type ListAndroidDeviceOwnerEnrollmentProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AndroidDeviceOwnerEnrollmentProfile
}

type ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AndroidDeviceOwnerEnrollmentProfile
}

type ListAndroidDeviceOwnerEnrollmentProfilesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListAndroidDeviceOwnerEnrollmentProfilesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAndroidDeviceOwnerEnrollmentProfiles ...
func (c AndroidDeviceOwnerEnrollmentProfileClient) ListAndroidDeviceOwnerEnrollmentProfiles(ctx context.Context) (result ListAndroidDeviceOwnerEnrollmentProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListAndroidDeviceOwnerEnrollmentProfilesCustomPager{},
		Path:       "/deviceManagement/androidDeviceOwnerEnrollmentProfiles",
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
		Values *[]beta.AndroidDeviceOwnerEnrollmentProfile `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAndroidDeviceOwnerEnrollmentProfilesComplete retrieves all the results into a single object
func (c AndroidDeviceOwnerEnrollmentProfileClient) ListAndroidDeviceOwnerEnrollmentProfilesComplete(ctx context.Context) (ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult, error) {
	return c.ListAndroidDeviceOwnerEnrollmentProfilesCompleteMatchingPredicate(ctx, AndroidDeviceOwnerEnrollmentProfileOperationPredicate{})
}

// ListAndroidDeviceOwnerEnrollmentProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AndroidDeviceOwnerEnrollmentProfileClient) ListAndroidDeviceOwnerEnrollmentProfilesCompleteMatchingPredicate(ctx context.Context, predicate AndroidDeviceOwnerEnrollmentProfileOperationPredicate) (result ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult, err error) {
	items := make([]beta.AndroidDeviceOwnerEnrollmentProfile, 0)

	resp, err := c.ListAndroidDeviceOwnerEnrollmentProfiles(ctx)
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

	result = ListAndroidDeviceOwnerEnrollmentProfilesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
