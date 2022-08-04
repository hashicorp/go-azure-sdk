package post

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotDpsResourceListKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]SharedAccessSignatureAuthorizationRuleAccessRightsDescription

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (IotDpsResourceListKeysOperationResponse, error)
}

type IotDpsResourceListKeysCompleteResult struct {
	Items []SharedAccessSignatureAuthorizationRuleAccessRightsDescription
}

func (r IotDpsResourceListKeysOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r IotDpsResourceListKeysOperationResponse) LoadMore(ctx context.Context) (resp IotDpsResourceListKeysOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// IotDpsResourceListKeys ...
func (c POSTClient) IotDpsResourceListKeys(ctx context.Context, id ProvisioningServiceId) (resp IotDpsResourceListKeysOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceListKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeys", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeys", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForIotDpsResourceListKeys(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeys", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// IotDpsResourceListKeysComplete retrieves all of the results into a single object
func (c POSTClient) IotDpsResourceListKeysComplete(ctx context.Context, id ProvisioningServiceId) (IotDpsResourceListKeysCompleteResult, error) {
	return c.IotDpsResourceListKeysCompleteMatchingPredicate(ctx, id, SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOperationPredicate{})
}

// IotDpsResourceListKeysCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c POSTClient) IotDpsResourceListKeysCompleteMatchingPredicate(ctx context.Context, id ProvisioningServiceId, predicate SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionOperationPredicate) (resp IotDpsResourceListKeysCompleteResult, err error) {
	items := make([]SharedAccessSignatureAuthorizationRuleAccessRightsDescription, 0)

	page, err := c.IotDpsResourceListKeys(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := IotDpsResourceListKeysCompleteResult{
		Items: items,
	}
	return out, nil
}

// preparerForIotDpsResourceListKeys prepares the IotDpsResourceListKeys request.
func (c POSTClient) preparerForIotDpsResourceListKeys(ctx context.Context, id ProvisioningServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listkeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForIotDpsResourceListKeysWithNextLink prepares the IotDpsResourceListKeys request with the given nextLink token.
func (c POSTClient) preparerForIotDpsResourceListKeysWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIotDpsResourceListKeys handles the response to the IotDpsResourceListKeys request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForIotDpsResourceListKeys(resp *http.Response) (result IotDpsResourceListKeysOperationResponse, err error) {
	type page struct {
		Values   []SharedAccessSignatureAuthorizationRuleAccessRightsDescription `json:"value"`
		NextLink *string                                                         `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result IotDpsResourceListKeysOperationResponse, err error) {
			req, err := c.preparerForIotDpsResourceListKeysWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeys", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeys", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForIotDpsResourceListKeys(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "post.POSTClient", "IotDpsResourceListKeys", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}
