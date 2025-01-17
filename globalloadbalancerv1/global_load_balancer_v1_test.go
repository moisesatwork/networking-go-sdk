/**
 * (C) Copyright IBM Corp. 2020.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package globalloadbalancerv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/networking-go-sdk/globalloadbalancerv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`GlobalLoadBalancerV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(globalLoadBalancerService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(globalLoadBalancerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
				URL:            "https://globalloadbalancerv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(globalLoadBalancerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{})
			Expect(globalLoadBalancerService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_URL":       "https://globalloadbalancerv1/api",
				"GLOBAL_LOAD_BALANCER_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1UsingExternalConfig(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(globalLoadBalancerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := globalLoadBalancerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalLoadBalancerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalLoadBalancerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalLoadBalancerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1UsingExternalConfig(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(globalLoadBalancerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalLoadBalancerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalLoadBalancerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalLoadBalancerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalLoadBalancerService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1UsingExternalConfig(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := globalLoadBalancerService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := globalLoadBalancerService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != globalLoadBalancerService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(globalLoadBalancerService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(globalLoadBalancerService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_URL":       "https://globalloadbalancerv1/api",
				"GLOBAL_LOAD_BALANCER_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1UsingExternalConfig(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"GLOBAL_LOAD_BALANCER_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1UsingExternalConfig(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(globalLoadBalancerService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = globalloadbalancerv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAllLoadBalancers(listAllLoadBalancersOptions *ListAllLoadBalancersOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllLoadBalancersPath := "/v1/testString/zones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllLoadBalancers with error: Operation response processing error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the ListAllLoadBalancersOptions model
				listAllLoadBalancersOptionsModel := new(globalloadbalancerv1.ListAllLoadBalancersOptions)
				listAllLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerService.ListAllLoadBalancers(listAllLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerService.ListAllLoadBalancers(listAllLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllLoadBalancers(listAllLoadBalancersOptions *ListAllLoadBalancersOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllLoadBalancersPath := "/v1/testString/zones/testString/load_balancers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllLoadBalancersPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "699d98642c564d2e855e9661899b7252", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Load Balancer for www.example.com", "name": "www.example.com", "ttl": 30, "fallback_pool": "17b5962d775c646f3f9725cbc7a53df4", "default_pools": ["DefaultPools"], "region_pools": {"anyKey": "anyValue"}, "pop_pools": {"anyKey": "anyValue"}, "proxied": true, "enabled": true, "session_affinity": "ip_cookie", "steering_policy": "dynamic_latency"}], "result_info": {"page": 1, "per_page": 20, "count": 1, "total_count": 2000}}`)
				}))
			})
			It(`Invoke ListAllLoadBalancers successfully`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())
				globalLoadBalancerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerService.ListAllLoadBalancers(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllLoadBalancersOptions model
				listAllLoadBalancersOptionsModel := new(globalloadbalancerv1.ListAllLoadBalancersOptions)
				listAllLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerService.ListAllLoadBalancers(listAllLoadBalancersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.ListAllLoadBalancersWithContext(ctx, listAllLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerService.DisableRetries()
				result, response, operationErr = globalLoadBalancerService.ListAllLoadBalancers(listAllLoadBalancersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.ListAllLoadBalancersWithContext(ctx, listAllLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAllLoadBalancers with error: Operation request error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the ListAllLoadBalancersOptions model
				listAllLoadBalancersOptionsModel := new(globalloadbalancerv1.ListAllLoadBalancersOptions)
				listAllLoadBalancersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerService.ListAllLoadBalancers(listAllLoadBalancersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createLoadBalancerPath := "/v1/testString/zones/testString/load_balancers"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLoadBalancerPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateLoadBalancer with error: Operation response processing error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(globalloadbalancerv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.Name = core.StringPtr("www.example.com")
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load Balancer for www.example.com")
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(30))
				createLoadBalancerOptionsModel.RegionPools = map[string]interface{}{"anyKey": "anyValue"}
				createLoadBalancerOptionsModel.PopPools = map[string]interface{}{"anyKey": "anyValue"}
				createLoadBalancerOptionsModel.Proxied = core.BoolPtr(true)
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.SessionAffinity = core.StringPtr("ip_cookie")
				createLoadBalancerOptionsModel.SteeringPolicy = core.StringPtr("dynamic_latency")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateLoadBalancer(createLoadBalancerOptions *CreateLoadBalancerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createLoadBalancerPath := "/v1/testString/zones/testString/load_balancers"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createLoadBalancerPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "699d98642c564d2e855e9661899b7252", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Load Balancer for www.example.com", "name": "www.example.com", "ttl": 30, "fallback_pool": "17b5962d775c646f3f9725cbc7a53df4", "default_pools": ["DefaultPools"], "region_pools": {"anyKey": "anyValue"}, "pop_pools": {"anyKey": "anyValue"}, "proxied": true, "enabled": true, "session_affinity": "ip_cookie", "steering_policy": "dynamic_latency"}}`)
				}))
			})
			It(`Invoke CreateLoadBalancer successfully`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())
				globalLoadBalancerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerService.CreateLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(globalloadbalancerv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.Name = core.StringPtr("www.example.com")
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load Balancer for www.example.com")
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(30))
				createLoadBalancerOptionsModel.RegionPools = map[string]interface{}{"anyKey": "anyValue"}
				createLoadBalancerOptionsModel.PopPools = map[string]interface{}{"anyKey": "anyValue"}
				createLoadBalancerOptionsModel.Proxied = core.BoolPtr(true)
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.SessionAffinity = core.StringPtr("ip_cookie")
				createLoadBalancerOptionsModel.SteeringPolicy = core.StringPtr("dynamic_latency")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.CreateLoadBalancerWithContext(ctx, createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerService.DisableRetries()
				result, response, operationErr = globalLoadBalancerService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.CreateLoadBalancerWithContext(ctx, createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateLoadBalancer with error: Operation request error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := new(globalloadbalancerv1.CreateLoadBalancerOptions)
				createLoadBalancerOptionsModel.Name = core.StringPtr("www.example.com")
				createLoadBalancerOptionsModel.FallbackPool = core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")
				createLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				createLoadBalancerOptionsModel.Description = core.StringPtr("Load Balancer for www.example.com")
				createLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(30))
				createLoadBalancerOptionsModel.RegionPools = map[string]interface{}{"anyKey": "anyValue"}
				createLoadBalancerOptionsModel.PopPools = map[string]interface{}{"anyKey": "anyValue"}
				createLoadBalancerOptionsModel.Proxied = core.BoolPtr(true)
				createLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				createLoadBalancerOptionsModel.SessionAffinity = core.StringPtr("ip_cookie")
				createLoadBalancerOptionsModel.SteeringPolicy = core.StringPtr("dynamic_latency")
				createLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerService.CreateLoadBalancer(createLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`EditLoadBalancer(editLoadBalancerOptions *EditLoadBalancerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		editLoadBalancerPath := "/v1/testString/zones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editLoadBalancerPath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke EditLoadBalancer with error: Operation response processing error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the EditLoadBalancerOptions model
				editLoadBalancerOptionsModel := new(globalloadbalancerv1.EditLoadBalancerOptions)
				editLoadBalancerOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				editLoadBalancerOptionsModel.Name = core.StringPtr("www.example.com")
				editLoadBalancerOptionsModel.FallbackPool = core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")
				editLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				editLoadBalancerOptionsModel.Description = core.StringPtr("Load Balancer for www.example.com")
				editLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(30))
				editLoadBalancerOptionsModel.RegionPools = map[string]interface{}{"anyKey": "anyValue"}
				editLoadBalancerOptionsModel.PopPools = map[string]interface{}{"anyKey": "anyValue"}
				editLoadBalancerOptionsModel.Proxied = core.BoolPtr(true)
				editLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				editLoadBalancerOptionsModel.SessionAffinity = core.StringPtr("ip_cookie")
				editLoadBalancerOptionsModel.SteeringPolicy = core.StringPtr("dynamic_latency")
				editLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerService.EditLoadBalancer(editLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerService.EditLoadBalancer(editLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`EditLoadBalancer(editLoadBalancerOptions *EditLoadBalancerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		editLoadBalancerPath := "/v1/testString/zones/testString/load_balancers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(editLoadBalancerPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "699d98642c564d2e855e9661899b7252", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Load Balancer for www.example.com", "name": "www.example.com", "ttl": 30, "fallback_pool": "17b5962d775c646f3f9725cbc7a53df4", "default_pools": ["DefaultPools"], "region_pools": {"anyKey": "anyValue"}, "pop_pools": {"anyKey": "anyValue"}, "proxied": true, "enabled": true, "session_affinity": "ip_cookie", "steering_policy": "dynamic_latency"}}`)
				}))
			})
			It(`Invoke EditLoadBalancer successfully`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())
				globalLoadBalancerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerService.EditLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the EditLoadBalancerOptions model
				editLoadBalancerOptionsModel := new(globalloadbalancerv1.EditLoadBalancerOptions)
				editLoadBalancerOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				editLoadBalancerOptionsModel.Name = core.StringPtr("www.example.com")
				editLoadBalancerOptionsModel.FallbackPool = core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")
				editLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				editLoadBalancerOptionsModel.Description = core.StringPtr("Load Balancer for www.example.com")
				editLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(30))
				editLoadBalancerOptionsModel.RegionPools = map[string]interface{}{"anyKey": "anyValue"}
				editLoadBalancerOptionsModel.PopPools = map[string]interface{}{"anyKey": "anyValue"}
				editLoadBalancerOptionsModel.Proxied = core.BoolPtr(true)
				editLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				editLoadBalancerOptionsModel.SessionAffinity = core.StringPtr("ip_cookie")
				editLoadBalancerOptionsModel.SteeringPolicy = core.StringPtr("dynamic_latency")
				editLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerService.EditLoadBalancer(editLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.EditLoadBalancerWithContext(ctx, editLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerService.DisableRetries()
				result, response, operationErr = globalLoadBalancerService.EditLoadBalancer(editLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.EditLoadBalancerWithContext(ctx, editLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke EditLoadBalancer with error: Operation validation and request error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the EditLoadBalancerOptions model
				editLoadBalancerOptionsModel := new(globalloadbalancerv1.EditLoadBalancerOptions)
				editLoadBalancerOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				editLoadBalancerOptionsModel.Name = core.StringPtr("www.example.com")
				editLoadBalancerOptionsModel.FallbackPool = core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")
				editLoadBalancerOptionsModel.DefaultPools = []string{"testString"}
				editLoadBalancerOptionsModel.Description = core.StringPtr("Load Balancer for www.example.com")
				editLoadBalancerOptionsModel.TTL = core.Int64Ptr(int64(30))
				editLoadBalancerOptionsModel.RegionPools = map[string]interface{}{"anyKey": "anyValue"}
				editLoadBalancerOptionsModel.PopPools = map[string]interface{}{"anyKey": "anyValue"}
				editLoadBalancerOptionsModel.Proxied = core.BoolPtr(true)
				editLoadBalancerOptionsModel.Enabled = core.BoolPtr(true)
				editLoadBalancerOptionsModel.SessionAffinity = core.StringPtr("ip_cookie")
				editLoadBalancerOptionsModel.SteeringPolicy = core.StringPtr("dynamic_latency")
				editLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerService.EditLoadBalancer(editLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the EditLoadBalancerOptions model with no property values
				editLoadBalancerOptionsModelNew := new(globalloadbalancerv1.EditLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancerService.EditLoadBalancer(editLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteLoadBalancerPath := "/v1/testString/zones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLoadBalancerPath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteLoadBalancer with error: Operation response processing error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(globalloadbalancerv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteLoadBalancer(deleteLoadBalancerOptions *DeleteLoadBalancerOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteLoadBalancerPath := "/v1/testString/zones/testString/load_balancers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteLoadBalancerPath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "699d98642c564d2e855e9661899b7252"}}`)
				}))
			})
			It(`Invoke DeleteLoadBalancer successfully`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())
				globalLoadBalancerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerService.DeleteLoadBalancer(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(globalloadbalancerv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.DeleteLoadBalancerWithContext(ctx, deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerService.DisableRetries()
				result, response, operationErr = globalLoadBalancerService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.DeleteLoadBalancerWithContext(ctx, deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteLoadBalancer with error: Operation validation and request error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the DeleteLoadBalancerOptions model
				deleteLoadBalancerOptionsModel := new(globalloadbalancerv1.DeleteLoadBalancerOptions)
				deleteLoadBalancerOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				deleteLoadBalancerOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerService.DeleteLoadBalancer(deleteLoadBalancerOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteLoadBalancerOptions model with no property values
				deleteLoadBalancerOptionsModelNew := new(globalloadbalancerv1.DeleteLoadBalancerOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancerService.DeleteLoadBalancer(deleteLoadBalancerOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLoadBalancerSettings(getLoadBalancerSettingsOptions *GetLoadBalancerSettingsOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getLoadBalancerSettingsPath := "/v1/testString/zones/testString/load_balancers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoadBalancerSettingsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLoadBalancerSettings with error: Operation response processing error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerSettingsOptions model
				getLoadBalancerSettingsOptionsModel := new(globalloadbalancerv1.GetLoadBalancerSettingsOptions)
				getLoadBalancerSettingsOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				getLoadBalancerSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := globalLoadBalancerService.GetLoadBalancerSettings(getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				globalLoadBalancerService.EnableRetries(0, 0)
				result, response, operationErr = globalLoadBalancerService.GetLoadBalancerSettings(getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLoadBalancerSettings(getLoadBalancerSettingsOptions *GetLoadBalancerSettingsOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getLoadBalancerSettingsPath := "/v1/testString/zones/testString/load_balancers/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLoadBalancerSettingsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "699d98642c564d2e855e9661899b7252", "created_on": "2014-01-01T05:20:00.12345Z", "modified_on": "2014-01-01T05:20:00.12345Z", "description": "Load Balancer for www.example.com", "name": "www.example.com", "ttl": 30, "fallback_pool": "17b5962d775c646f3f9725cbc7a53df4", "default_pools": ["DefaultPools"], "region_pools": {"anyKey": "anyValue"}, "pop_pools": {"anyKey": "anyValue"}, "proxied": true, "enabled": true, "session_affinity": "ip_cookie", "steering_policy": "dynamic_latency"}}`)
				}))
			})
			It(`Invoke GetLoadBalancerSettings successfully`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())
				globalLoadBalancerService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := globalLoadBalancerService.GetLoadBalancerSettings(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLoadBalancerSettingsOptions model
				getLoadBalancerSettingsOptionsModel := new(globalloadbalancerv1.GetLoadBalancerSettingsOptions)
				getLoadBalancerSettingsOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				getLoadBalancerSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = globalLoadBalancerService.GetLoadBalancerSettings(getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.GetLoadBalancerSettingsWithContext(ctx, getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				globalLoadBalancerService.DisableRetries()
				result, response, operationErr = globalLoadBalancerService.GetLoadBalancerSettings(getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = globalLoadBalancerService.GetLoadBalancerSettingsWithContext(ctx, getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLoadBalancerSettings with error: Operation validation and request error`, func() {
				globalLoadBalancerService, serviceErr := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(globalLoadBalancerService).ToNot(BeNil())

				// Construct an instance of the GetLoadBalancerSettingsOptions model
				getLoadBalancerSettingsOptionsModel := new(globalloadbalancerv1.GetLoadBalancerSettingsOptions)
				getLoadBalancerSettingsOptionsModel.LoadBalancerIdentifier = core.StringPtr("testString")
				getLoadBalancerSettingsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := globalLoadBalancerService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := globalLoadBalancerService.GetLoadBalancerSettings(getLoadBalancerSettingsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLoadBalancerSettingsOptions model with no property values
				getLoadBalancerSettingsOptionsModelNew := new(globalloadbalancerv1.GetLoadBalancerSettingsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = globalLoadBalancerService.GetLoadBalancerSettings(getLoadBalancerSettingsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			crn := "testString"
			zoneIdentifier := "testString"
			globalLoadBalancerService, _ := globalloadbalancerv1.NewGlobalLoadBalancerV1(&globalloadbalancerv1.GlobalLoadBalancerV1Options{
				URL:            "http://globalloadbalancerv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateLoadBalancerOptions successfully`, func() {
				// Construct an instance of the CreateLoadBalancerOptions model
				createLoadBalancerOptionsModel := globalLoadBalancerService.NewCreateLoadBalancerOptions()
				createLoadBalancerOptionsModel.SetName("www.example.com")
				createLoadBalancerOptionsModel.SetFallbackPool("17b5962d775c646f3f9725cbc7a53df4")
				createLoadBalancerOptionsModel.SetDefaultPools([]string{"testString"})
				createLoadBalancerOptionsModel.SetDescription("Load Balancer for www.example.com")
				createLoadBalancerOptionsModel.SetTTL(int64(30))
				createLoadBalancerOptionsModel.SetRegionPools(map[string]interface{}{"anyKey": "anyValue"})
				createLoadBalancerOptionsModel.SetPopPools(map[string]interface{}{"anyKey": "anyValue"})
				createLoadBalancerOptionsModel.SetProxied(true)
				createLoadBalancerOptionsModel.SetEnabled(true)
				createLoadBalancerOptionsModel.SetSessionAffinity("ip_cookie")
				createLoadBalancerOptionsModel.SetSteeringPolicy("dynamic_latency")
				createLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(createLoadBalancerOptionsModel.Name).To(Equal(core.StringPtr("www.example.com")))
				Expect(createLoadBalancerOptionsModel.FallbackPool).To(Equal(core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")))
				Expect(createLoadBalancerOptionsModel.DefaultPools).To(Equal([]string{"testString"}))
				Expect(createLoadBalancerOptionsModel.Description).To(Equal(core.StringPtr("Load Balancer for www.example.com")))
				Expect(createLoadBalancerOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(30))))
				Expect(createLoadBalancerOptionsModel.RegionPools).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createLoadBalancerOptionsModel.PopPools).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(createLoadBalancerOptionsModel.Proxied).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(createLoadBalancerOptionsModel.SessionAffinity).To(Equal(core.StringPtr("ip_cookie")))
				Expect(createLoadBalancerOptionsModel.SteeringPolicy).To(Equal(core.StringPtr("dynamic_latency")))
				Expect(createLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteLoadBalancerOptions successfully`, func() {
				// Construct an instance of the DeleteLoadBalancerOptions model
				loadBalancerIdentifier := "testString"
				deleteLoadBalancerOptionsModel := globalLoadBalancerService.NewDeleteLoadBalancerOptions(loadBalancerIdentifier)
				deleteLoadBalancerOptionsModel.SetLoadBalancerIdentifier("testString")
				deleteLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(deleteLoadBalancerOptionsModel.LoadBalancerIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewEditLoadBalancerOptions successfully`, func() {
				// Construct an instance of the EditLoadBalancerOptions model
				loadBalancerIdentifier := "testString"
				editLoadBalancerOptionsModel := globalLoadBalancerService.NewEditLoadBalancerOptions(loadBalancerIdentifier)
				editLoadBalancerOptionsModel.SetLoadBalancerIdentifier("testString")
				editLoadBalancerOptionsModel.SetName("www.example.com")
				editLoadBalancerOptionsModel.SetFallbackPool("17b5962d775c646f3f9725cbc7a53df4")
				editLoadBalancerOptionsModel.SetDefaultPools([]string{"testString"})
				editLoadBalancerOptionsModel.SetDescription("Load Balancer for www.example.com")
				editLoadBalancerOptionsModel.SetTTL(int64(30))
				editLoadBalancerOptionsModel.SetRegionPools(map[string]interface{}{"anyKey": "anyValue"})
				editLoadBalancerOptionsModel.SetPopPools(map[string]interface{}{"anyKey": "anyValue"})
				editLoadBalancerOptionsModel.SetProxied(true)
				editLoadBalancerOptionsModel.SetEnabled(true)
				editLoadBalancerOptionsModel.SetSessionAffinity("ip_cookie")
				editLoadBalancerOptionsModel.SetSteeringPolicy("dynamic_latency")
				editLoadBalancerOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(editLoadBalancerOptionsModel).ToNot(BeNil())
				Expect(editLoadBalancerOptionsModel.LoadBalancerIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(editLoadBalancerOptionsModel.Name).To(Equal(core.StringPtr("www.example.com")))
				Expect(editLoadBalancerOptionsModel.FallbackPool).To(Equal(core.StringPtr("17b5962d775c646f3f9725cbc7a53df4")))
				Expect(editLoadBalancerOptionsModel.DefaultPools).To(Equal([]string{"testString"}))
				Expect(editLoadBalancerOptionsModel.Description).To(Equal(core.StringPtr("Load Balancer for www.example.com")))
				Expect(editLoadBalancerOptionsModel.TTL).To(Equal(core.Int64Ptr(int64(30))))
				Expect(editLoadBalancerOptionsModel.RegionPools).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(editLoadBalancerOptionsModel.PopPools).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(editLoadBalancerOptionsModel.Proxied).To(Equal(core.BoolPtr(true)))
				Expect(editLoadBalancerOptionsModel.Enabled).To(Equal(core.BoolPtr(true)))
				Expect(editLoadBalancerOptionsModel.SessionAffinity).To(Equal(core.StringPtr("ip_cookie")))
				Expect(editLoadBalancerOptionsModel.SteeringPolicy).To(Equal(core.StringPtr("dynamic_latency")))
				Expect(editLoadBalancerOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLoadBalancerSettingsOptions successfully`, func() {
				// Construct an instance of the GetLoadBalancerSettingsOptions model
				loadBalancerIdentifier := "testString"
				getLoadBalancerSettingsOptionsModel := globalLoadBalancerService.NewGetLoadBalancerSettingsOptions(loadBalancerIdentifier)
				getLoadBalancerSettingsOptionsModel.SetLoadBalancerIdentifier("testString")
				getLoadBalancerSettingsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLoadBalancerSettingsOptionsModel).ToNot(BeNil())
				Expect(getLoadBalancerSettingsOptionsModel.LoadBalancerIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getLoadBalancerSettingsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllLoadBalancersOptions successfully`, func() {
				// Construct an instance of the ListAllLoadBalancersOptions model
				listAllLoadBalancersOptionsModel := globalLoadBalancerService.NewListAllLoadBalancersOptions()
				listAllLoadBalancersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllLoadBalancersOptionsModel).ToNot(BeNil())
				Expect(listAllLoadBalancersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate()
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime()
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate() *strfmt.Date {
	d := strfmt.Date(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func CreateMockDateTime() *strfmt.DateTime {
	d := strfmt.DateTime(time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC))
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
