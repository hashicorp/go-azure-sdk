package directoryinboundshareduserprofile

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryInboundSharedUserProfilesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.InboundSharedUserProfileCollectionResponse
}

type ListDirectoryInboundSharedUserProfilesCompleteResult struct {
	Items []models.InboundSharedUserProfileCollectionResponse
}

// ListDirectoryInboundSharedUserProfiles ...
func (c DirectoryInboundSharedUserProfileClient) ListDirectoryInboundSharedUserProfiles(ctx context.Context) (result ListDirectoryInboundSharedUserProfilesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.InboundSharedUserProfileCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryInboundSharedUserProfilesComplete retrieves all the results into a single object
func (c DirectoryInboundSharedUserProfileClient) ListDirectoryInboundSharedUserProfilesComplete(ctx context.Context) (ListDirectoryInboundSharedUserProfilesCompleteResult, error) {
	return c.ListDirectoryInboundSharedUserProfilesCompleteMatchingPredicate(ctx, models.InboundSharedUserProfileCollectionResponseOperationPredicate{})
}

// ListDirectoryInboundSharedUserProfilesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryInboundSharedUserProfileClient) ListDirectoryInboundSharedUserProfilesCompleteMatchingPredicate(ctx context.Context, predicate models.InboundSharedUserProfileCollectionResponseOperationPredicate) (result ListDirectoryInboundSharedUserProfilesCompleteResult, err error) {
	items := make([]models.InboundSharedUserProfileCollectionResponse, 0)

	resp, err := c.ListDirectoryInboundSharedUserProfiles(ctx)
	if err != nil {
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

	result = ListDirectoryInboundSharedUserProfilesCompleteResult{
		Items: items,
	}
	return
}
