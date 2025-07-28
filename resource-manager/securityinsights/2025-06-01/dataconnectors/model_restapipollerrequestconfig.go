package dataconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestApiPollerRequestConfig struct {
	ApiEndpoint                    string                  `json:"apiEndpoint"`
	EndTimeAttributeName           *string                 `json:"endTimeAttributeName,omitempty"`
	HTTPMethod                     *HTTPMethodVerb         `json:"httpMethod,omitempty"`
	Headers                        *map[string]string      `json:"headers,omitempty"`
	IsPostPayloadJson              *bool                   `json:"isPostPayloadJson,omitempty"`
	QueryParameters                *map[string]interface{} `json:"queryParameters,omitempty"`
	QueryParametersTemplate        *string                 `json:"queryParametersTemplate,omitempty"`
	QueryTimeFormat                *string                 `json:"queryTimeFormat,omitempty"`
	QueryTimeIntervalAttributeName *string                 `json:"queryTimeIntervalAttributeName,omitempty"`
	QueryTimeIntervalDelimiter     *string                 `json:"queryTimeIntervalDelimiter,omitempty"`
	QueryTimeIntervalPrepend       *string                 `json:"queryTimeIntervalPrepend,omitempty"`
	QueryWindowInMin               *int64                  `json:"queryWindowInMin,omitempty"`
	RateLimitQPS                   *int64                  `json:"rateLimitQPS,omitempty"`
	RetryCount                     *int64                  `json:"retryCount,omitempty"`
	StartTimeAttributeName         *string                 `json:"startTimeAttributeName,omitempty"`
	TimeoutInSeconds               *int64                  `json:"timeoutInSeconds,omitempty"`
}
