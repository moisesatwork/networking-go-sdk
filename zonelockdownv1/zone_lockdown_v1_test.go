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

package zonelockdownv1_test

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
	"github.com/IBM/networking-go-sdk/zonelockdownv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe(`ZoneLockdownV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		It(`Instantiate service client`, func() {
			zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(zoneLockdownService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			Expect(zoneLockdownService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
				URL:            "https://zonelockdownv1/api",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(zoneLockdownService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Validation Error`, func() {
			zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{})
			Expect(zoneLockdownService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONE_LOCKDOWN_URL":       "https://zonelockdownv1/api",
				"ZONE_LOCKDOWN_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1UsingExternalConfig(&zonelockdownv1.ZoneLockdownV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(zoneLockdownService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := zoneLockdownService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zoneLockdownService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zoneLockdownService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zoneLockdownService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1UsingExternalConfig(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            "https://testService/api",
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(zoneLockdownService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := zoneLockdownService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zoneLockdownService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zoneLockdownService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zoneLockdownService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1UsingExternalConfig(&zonelockdownv1.ZoneLockdownV1Options{
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				err := zoneLockdownService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := zoneLockdownService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != zoneLockdownService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(zoneLockdownService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(zoneLockdownService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONE_LOCKDOWN_URL":       "https://zonelockdownv1/api",
				"ZONE_LOCKDOWN_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1UsingExternalConfig(&zonelockdownv1.ZoneLockdownV1Options{
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(zoneLockdownService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"ZONE_LOCKDOWN_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1UsingExternalConfig(&zonelockdownv1.ZoneLockdownV1Options{
				URL:            "{BAD_URL_STRING",
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})

			It(`Instantiate service client with error`, func() {
				Expect(zoneLockdownService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = zonelockdownv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListAllZoneLockownRules(listAllZoneLockownRulesOptions *ListAllZoneLockownRulesOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneLockownRulesPath := "/v1/testString/zones/testString/firewall/lockdowns"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllZoneLockownRulesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListAllZoneLockownRules with error: Operation response processing error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneLockownRulesOptions model
				listAllZoneLockownRulesOptionsModel := new(zonelockdownv1.ListAllZoneLockownRulesOptions)
				listAllZoneLockownRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneLockownRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneLockownRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zoneLockdownService.ListAllZoneLockownRules(listAllZoneLockownRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zoneLockdownService.EnableRetries(0, 0)
				result, response, operationErr = zoneLockdownService.ListAllZoneLockownRules(listAllZoneLockownRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`ListAllZoneLockownRules(listAllZoneLockownRulesOptions *ListAllZoneLockownRulesOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		listAllZoneLockownRulesPath := "/v1/testString/zones/testString/firewall/lockdowns"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listAllZoneLockownRulesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["page"]).To(Equal([]string{fmt.Sprint(int64(38))}))

					Expect(req.URL.Query()["per_page"]).To(Equal([]string{fmt.Sprint(int64(5))}))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": [{"id": "372e67954025e0ba6aaa6d586b9e0b59", "priority": 5, "paused": false, "description": "Restrict access to these endpoints to requests from a known IP address", "urls": ["api.mysite.com/some/endpoint*"], "configurations": [{"target": "ip", "value": "198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range"}]}], "result_info": {"page": 1, "per_page": 2, "count": 1, "total_count": 200}}`)
				}))
			})
			It(`Invoke ListAllZoneLockownRules successfully`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())
				zoneLockdownService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zoneLockdownService.ListAllZoneLockownRules(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListAllZoneLockownRulesOptions model
				listAllZoneLockownRulesOptionsModel := new(zonelockdownv1.ListAllZoneLockownRulesOptions)
				listAllZoneLockownRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneLockownRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneLockownRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zoneLockdownService.ListAllZoneLockownRules(listAllZoneLockownRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.ListAllZoneLockownRulesWithContext(ctx, listAllZoneLockownRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zoneLockdownService.DisableRetries()
				result, response, operationErr = zoneLockdownService.ListAllZoneLockownRules(listAllZoneLockownRulesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.ListAllZoneLockownRulesWithContext(ctx, listAllZoneLockownRulesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke ListAllZoneLockownRules with error: Operation request error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the ListAllZoneLockownRulesOptions model
				listAllZoneLockownRulesOptionsModel := new(zonelockdownv1.ListAllZoneLockownRulesOptions)
				listAllZoneLockownRulesOptionsModel.Page = core.Int64Ptr(int64(38))
				listAllZoneLockownRulesOptionsModel.PerPage = core.Int64Ptr(int64(5))
				listAllZoneLockownRulesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zoneLockdownService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zoneLockdownService.ListAllZoneLockownRules(listAllZoneLockownRulesOptionsModel)
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
	Describe(`CreateZoneLockdownRule(createZoneLockdownRuleOptions *CreateZoneLockdownRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneLockdownRulePath := "/v1/testString/zones/testString/firewall/lockdowns"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneLockdownRulePath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateZoneLockdownRule with error: Operation response processing error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")

				// Construct an instance of the CreateZoneLockdownRuleOptions model
				createZoneLockdownRuleOptionsModel := new(zonelockdownv1.CreateZoneLockdownRuleOptions)
				createZoneLockdownRuleOptionsModel.ID = core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")
				createZoneLockdownRuleOptionsModel.Paused = core.BoolPtr(false)
				createZoneLockdownRuleOptionsModel.Description = core.StringPtr("Restrict access to these endpoints to requests from a known IP address")
				createZoneLockdownRuleOptionsModel.Urls = []string{"api.mysite.com/some/endpoint*"}
				createZoneLockdownRuleOptionsModel.Configurations = []zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}
				createZoneLockdownRuleOptionsModel.Priority = core.Int64Ptr(int64(5))
				createZoneLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zoneLockdownService.CreateZoneLockdownRule(createZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zoneLockdownService.EnableRetries(0, 0)
				result, response, operationErr = zoneLockdownService.CreateZoneLockdownRule(createZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`CreateZoneLockdownRule(createZoneLockdownRuleOptions *CreateZoneLockdownRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		createZoneLockdownRulePath := "/v1/testString/zones/testString/firewall/lockdowns"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createZoneLockdownRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "372e67954025e0ba6aaa6d586b9e0b59", "priority": 5, "paused": false, "description": "Restrict access to these endpoints to requests from a known IP address", "urls": ["api.mysite.com/some/endpoint*"], "configurations": [{"target": "ip", "value": "198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range"}]}}`)
				}))
			})
			It(`Invoke CreateZoneLockdownRule successfully`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())
				zoneLockdownService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zoneLockdownService.CreateZoneLockdownRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")

				// Construct an instance of the CreateZoneLockdownRuleOptions model
				createZoneLockdownRuleOptionsModel := new(zonelockdownv1.CreateZoneLockdownRuleOptions)
				createZoneLockdownRuleOptionsModel.ID = core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")
				createZoneLockdownRuleOptionsModel.Paused = core.BoolPtr(false)
				createZoneLockdownRuleOptionsModel.Description = core.StringPtr("Restrict access to these endpoints to requests from a known IP address")
				createZoneLockdownRuleOptionsModel.Urls = []string{"api.mysite.com/some/endpoint*"}
				createZoneLockdownRuleOptionsModel.Configurations = []zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}
				createZoneLockdownRuleOptionsModel.Priority = core.Int64Ptr(int64(5))
				createZoneLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zoneLockdownService.CreateZoneLockdownRule(createZoneLockdownRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.CreateZoneLockdownRuleWithContext(ctx, createZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zoneLockdownService.DisableRetries()
				result, response, operationErr = zoneLockdownService.CreateZoneLockdownRule(createZoneLockdownRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.CreateZoneLockdownRuleWithContext(ctx, createZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke CreateZoneLockdownRule with error: Operation request error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")

				// Construct an instance of the CreateZoneLockdownRuleOptions model
				createZoneLockdownRuleOptionsModel := new(zonelockdownv1.CreateZoneLockdownRuleOptions)
				createZoneLockdownRuleOptionsModel.ID = core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")
				createZoneLockdownRuleOptionsModel.Paused = core.BoolPtr(false)
				createZoneLockdownRuleOptionsModel.Description = core.StringPtr("Restrict access to these endpoints to requests from a known IP address")
				createZoneLockdownRuleOptionsModel.Urls = []string{"api.mysite.com/some/endpoint*"}
				createZoneLockdownRuleOptionsModel.Configurations = []zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}
				createZoneLockdownRuleOptionsModel.Priority = core.Int64Ptr(int64(5))
				createZoneLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zoneLockdownService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zoneLockdownService.CreateZoneLockdownRule(createZoneLockdownRuleOptionsModel)
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
	Describe(`DeleteZoneLockdownRule(deleteZoneLockdownRuleOptions *DeleteZoneLockdownRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneLockdownRulePath := "/v1/testString/zones/testString/firewall/lockdowns/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneLockdownRulePath))
					Expect(req.Method).To(Equal("DELETE"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteZoneLockdownRule with error: Operation response processing error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneLockdownRuleOptions model
				deleteZoneLockdownRuleOptionsModel := new(zonelockdownv1.DeleteZoneLockdownRuleOptions)
				deleteZoneLockdownRuleOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				deleteZoneLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zoneLockdownService.DeleteZoneLockdownRule(deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zoneLockdownService.EnableRetries(0, 0)
				result, response, operationErr = zoneLockdownService.DeleteZoneLockdownRule(deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`DeleteZoneLockdownRule(deleteZoneLockdownRuleOptions *DeleteZoneLockdownRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		deleteZoneLockdownRulePath := "/v1/testString/zones/testString/firewall/lockdowns/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteZoneLockdownRulePath))
					Expect(req.Method).To(Equal("DELETE"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "f1aba936b94213e5b8dca0c0dbf1f9cc"}}`)
				}))
			})
			It(`Invoke DeleteZoneLockdownRule successfully`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())
				zoneLockdownService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zoneLockdownService.DeleteZoneLockdownRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteZoneLockdownRuleOptions model
				deleteZoneLockdownRuleOptionsModel := new(zonelockdownv1.DeleteZoneLockdownRuleOptions)
				deleteZoneLockdownRuleOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				deleteZoneLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zoneLockdownService.DeleteZoneLockdownRule(deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.DeleteZoneLockdownRuleWithContext(ctx, deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zoneLockdownService.DisableRetries()
				result, response, operationErr = zoneLockdownService.DeleteZoneLockdownRule(deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.DeleteZoneLockdownRuleWithContext(ctx, deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke DeleteZoneLockdownRule with error: Operation validation and request error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the DeleteZoneLockdownRuleOptions model
				deleteZoneLockdownRuleOptionsModel := new(zonelockdownv1.DeleteZoneLockdownRuleOptions)
				deleteZoneLockdownRuleOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				deleteZoneLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zoneLockdownService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zoneLockdownService.DeleteZoneLockdownRule(deleteZoneLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteZoneLockdownRuleOptions model with no property values
				deleteZoneLockdownRuleOptionsModelNew := new(zonelockdownv1.DeleteZoneLockdownRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zoneLockdownService.DeleteZoneLockdownRule(deleteZoneLockdownRuleOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetLockdown(getLockdownOptions *GetLockdownOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getLockdownPath := "/v1/testString/zones/testString/firewall/lockdowns/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLockdownPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetLockdown with error: Operation response processing error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the GetLockdownOptions model
				getLockdownOptionsModel := new(zonelockdownv1.GetLockdownOptions)
				getLockdownOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				getLockdownOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zoneLockdownService.GetLockdown(getLockdownOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zoneLockdownService.EnableRetries(0, 0)
				result, response, operationErr = zoneLockdownService.GetLockdown(getLockdownOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`GetLockdown(getLockdownOptions *GetLockdownOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		getLockdownPath := "/v1/testString/zones/testString/firewall/lockdowns/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getLockdownPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(serverSleepTime)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "372e67954025e0ba6aaa6d586b9e0b59", "priority": 5, "paused": false, "description": "Restrict access to these endpoints to requests from a known IP address", "urls": ["api.mysite.com/some/endpoint*"], "configurations": [{"target": "ip", "value": "198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range"}]}}`)
				}))
			})
			It(`Invoke GetLockdown successfully`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())
				zoneLockdownService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zoneLockdownService.GetLockdown(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetLockdownOptions model
				getLockdownOptionsModel := new(zonelockdownv1.GetLockdownOptions)
				getLockdownOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				getLockdownOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zoneLockdownService.GetLockdown(getLockdownOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.GetLockdownWithContext(ctx, getLockdownOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zoneLockdownService.DisableRetries()
				result, response, operationErr = zoneLockdownService.GetLockdown(getLockdownOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.GetLockdownWithContext(ctx, getLockdownOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke GetLockdown with error: Operation validation and request error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the GetLockdownOptions model
				getLockdownOptionsModel := new(zonelockdownv1.GetLockdownOptions)
				getLockdownOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				getLockdownOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zoneLockdownService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zoneLockdownService.GetLockdown(getLockdownOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetLockdownOptions model with no property values
				getLockdownOptionsModelNew := new(zonelockdownv1.GetLockdownOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zoneLockdownService.GetLockdown(getLockdownOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateLockdownRule(updateLockdownRuleOptions *UpdateLockdownRuleOptions) - Operation response error`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateLockdownRulePath := "/v1/testString/zones/testString/firewall/lockdowns/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLockdownRulePath))
					Expect(req.Method).To(Equal("PUT"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateLockdownRule with error: Operation response processing error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")

				// Construct an instance of the UpdateLockdownRuleOptions model
				updateLockdownRuleOptionsModel := new(zonelockdownv1.UpdateLockdownRuleOptions)
				updateLockdownRuleOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				updateLockdownRuleOptionsModel.ID = core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")
				updateLockdownRuleOptionsModel.Paused = core.BoolPtr(false)
				updateLockdownRuleOptionsModel.Description = core.StringPtr("Restrict access to these endpoints to requests from a known IP address")
				updateLockdownRuleOptionsModel.Urls = []string{"api.mysite.com/some/endpoint*"}
				updateLockdownRuleOptionsModel.Configurations = []zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}
				updateLockdownRuleOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := zoneLockdownService.UpdateLockdownRule(updateLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				zoneLockdownService.EnableRetries(0, 0)
				result, response, operationErr = zoneLockdownService.UpdateLockdownRule(updateLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})

	Describe(`UpdateLockdownRule(updateLockdownRuleOptions *UpdateLockdownRuleOptions)`, func() {
		crn := "testString"
		zoneIdentifier := "testString"
		updateLockdownRulePath := "/v1/testString/zones/testString/firewall/lockdowns/testString"
		var serverSleepTime time.Duration
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				serverSleepTime = 0
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateLockdownRulePath))
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
					fmt.Fprintf(res, "%s", `{"success": true, "errors": [["Errors"]], "messages": [["Messages"]], "result": {"id": "372e67954025e0ba6aaa6d586b9e0b59", "priority": 5, "paused": false, "description": "Restrict access to these endpoints to requests from a known IP address", "urls": ["api.mysite.com/some/endpoint*"], "configurations": [{"target": "ip", "value": "198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range"}]}}`)
				}))
			})
			It(`Invoke UpdateLockdownRule successfully`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())
				zoneLockdownService.EnableRetries(0, 0)

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := zoneLockdownService.UpdateLockdownRule(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")

				// Construct an instance of the UpdateLockdownRuleOptions model
				updateLockdownRuleOptionsModel := new(zonelockdownv1.UpdateLockdownRuleOptions)
				updateLockdownRuleOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				updateLockdownRuleOptionsModel.ID = core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")
				updateLockdownRuleOptionsModel.Paused = core.BoolPtr(false)
				updateLockdownRuleOptionsModel.Description = core.StringPtr("Restrict access to these endpoints to requests from a known IP address")
				updateLockdownRuleOptionsModel.Urls = []string{"api.mysite.com/some/endpoint*"}
				updateLockdownRuleOptionsModel.Configurations = []zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}
				updateLockdownRuleOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = zoneLockdownService.UpdateLockdownRule(updateLockdownRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.UpdateLockdownRuleWithContext(ctx, updateLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)

				// Disable retries and test again
				zoneLockdownService.DisableRetries()
				result, response, operationErr = zoneLockdownService.UpdateLockdownRule(updateLockdownRuleOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				serverSleepTime = 100 * time.Millisecond
				_, _, operationErr = zoneLockdownService.UpdateLockdownRuleWithContext(ctx, updateLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
				serverSleepTime = time.Duration(0)
			})
			It(`Invoke UpdateLockdownRule with error: Operation validation and request error`, func() {
				zoneLockdownService, serviceErr := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
					URL:            testServer.URL,
					Authenticator:  &core.NoAuthAuthenticator{},
					Crn:            core.StringPtr(crn),
					ZoneIdentifier: core.StringPtr(zoneIdentifier),
				})
				Expect(serviceErr).To(BeNil())
				Expect(zoneLockdownService).ToNot(BeNil())

				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")

				// Construct an instance of the UpdateLockdownRuleOptions model
				updateLockdownRuleOptionsModel := new(zonelockdownv1.UpdateLockdownRuleOptions)
				updateLockdownRuleOptionsModel.LockdownRuleIdentifier = core.StringPtr("testString")
				updateLockdownRuleOptionsModel.ID = core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")
				updateLockdownRuleOptionsModel.Paused = core.BoolPtr(false)
				updateLockdownRuleOptionsModel.Description = core.StringPtr("Restrict access to these endpoints to requests from a known IP address")
				updateLockdownRuleOptionsModel.Urls = []string{"api.mysite.com/some/endpoint*"}
				updateLockdownRuleOptionsModel.Configurations = []zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}
				updateLockdownRuleOptionsModel.Priority = core.Int64Ptr(int64(5))
				updateLockdownRuleOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := zoneLockdownService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := zoneLockdownService.UpdateLockdownRule(updateLockdownRuleOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateLockdownRuleOptions model with no property values
				updateLockdownRuleOptionsModelNew := new(zonelockdownv1.UpdateLockdownRuleOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = zoneLockdownService.UpdateLockdownRule(updateLockdownRuleOptionsModelNew)
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
			zoneLockdownService, _ := zonelockdownv1.NewZoneLockdownV1(&zonelockdownv1.ZoneLockdownV1Options{
				URL:            "http://zonelockdownv1modelgenerator.com",
				Authenticator:  &core.NoAuthAuthenticator{},
				Crn:            core.StringPtr(crn),
				ZoneIdentifier: core.StringPtr(zoneIdentifier),
			})
			It(`Invoke NewCreateZoneLockdownRuleOptions successfully`, func() {
				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				Expect(lockdownInputConfigurationsItemModel).ToNot(BeNil())
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")
				Expect(lockdownInputConfigurationsItemModel.Target).To(Equal(core.StringPtr("ip")))
				Expect(lockdownInputConfigurationsItemModel.Value).To(Equal(core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")))

				// Construct an instance of the CreateZoneLockdownRuleOptions model
				createZoneLockdownRuleOptionsModel := zoneLockdownService.NewCreateZoneLockdownRuleOptions()
				createZoneLockdownRuleOptionsModel.SetID("372e67954025e0ba6aaa6d586b9e0b59")
				createZoneLockdownRuleOptionsModel.SetPaused(false)
				createZoneLockdownRuleOptionsModel.SetDescription("Restrict access to these endpoints to requests from a known IP address")
				createZoneLockdownRuleOptionsModel.SetUrls([]string{"api.mysite.com/some/endpoint*"})
				createZoneLockdownRuleOptionsModel.SetConfigurations([]zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel})
				createZoneLockdownRuleOptionsModel.SetPriority(int64(5))
				createZoneLockdownRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createZoneLockdownRuleOptionsModel).ToNot(BeNil())
				Expect(createZoneLockdownRuleOptionsModel.ID).To(Equal(core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")))
				Expect(createZoneLockdownRuleOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(createZoneLockdownRuleOptionsModel.Description).To(Equal(core.StringPtr("Restrict access to these endpoints to requests from a known IP address")))
				Expect(createZoneLockdownRuleOptionsModel.Urls).To(Equal([]string{"api.mysite.com/some/endpoint*"}))
				Expect(createZoneLockdownRuleOptionsModel.Configurations).To(Equal([]zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}))
				Expect(createZoneLockdownRuleOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(createZoneLockdownRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteZoneLockdownRuleOptions successfully`, func() {
				// Construct an instance of the DeleteZoneLockdownRuleOptions model
				lockdownRuleIdentifier := "testString"
				deleteZoneLockdownRuleOptionsModel := zoneLockdownService.NewDeleteZoneLockdownRuleOptions(lockdownRuleIdentifier)
				deleteZoneLockdownRuleOptionsModel.SetLockdownRuleIdentifier("testString")
				deleteZoneLockdownRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteZoneLockdownRuleOptionsModel).ToNot(BeNil())
				Expect(deleteZoneLockdownRuleOptionsModel.LockdownRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(deleteZoneLockdownRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetLockdownOptions successfully`, func() {
				// Construct an instance of the GetLockdownOptions model
				lockdownRuleIdentifier := "testString"
				getLockdownOptionsModel := zoneLockdownService.NewGetLockdownOptions(lockdownRuleIdentifier)
				getLockdownOptionsModel.SetLockdownRuleIdentifier("testString")
				getLockdownOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getLockdownOptionsModel).ToNot(BeNil())
				Expect(getLockdownOptionsModel.LockdownRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(getLockdownOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListAllZoneLockownRulesOptions successfully`, func() {
				// Construct an instance of the ListAllZoneLockownRulesOptions model
				listAllZoneLockownRulesOptionsModel := zoneLockdownService.NewListAllZoneLockownRulesOptions()
				listAllZoneLockownRulesOptionsModel.SetPage(int64(38))
				listAllZoneLockownRulesOptionsModel.SetPerPage(int64(5))
				listAllZoneLockownRulesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listAllZoneLockownRulesOptionsModel).ToNot(BeNil())
				Expect(listAllZoneLockownRulesOptionsModel.Page).To(Equal(core.Int64Ptr(int64(38))))
				Expect(listAllZoneLockownRulesOptionsModel.PerPage).To(Equal(core.Int64Ptr(int64(5))))
				Expect(listAllZoneLockownRulesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewLockdownInputConfigurationsItem successfully`, func() {
				target := "ip"
				value := "198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range"
				model, err := zoneLockdownService.NewLockdownInputConfigurationsItem(target, value)
				Expect(model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateLockdownRuleOptions successfully`, func() {
				// Construct an instance of the LockdownInputConfigurationsItem model
				lockdownInputConfigurationsItemModel := new(zonelockdownv1.LockdownInputConfigurationsItem)
				Expect(lockdownInputConfigurationsItemModel).ToNot(BeNil())
				lockdownInputConfigurationsItemModel.Target = core.StringPtr("ip")
				lockdownInputConfigurationsItemModel.Value = core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")
				Expect(lockdownInputConfigurationsItemModel.Target).To(Equal(core.StringPtr("ip")))
				Expect(lockdownInputConfigurationsItemModel.Value).To(Equal(core.StringPtr("198.51.100.4 if target=ip, 2.2.2.0/24 if target=ip_range")))

				// Construct an instance of the UpdateLockdownRuleOptions model
				lockdownRuleIdentifier := "testString"
				updateLockdownRuleOptionsModel := zoneLockdownService.NewUpdateLockdownRuleOptions(lockdownRuleIdentifier)
				updateLockdownRuleOptionsModel.SetLockdownRuleIdentifier("testString")
				updateLockdownRuleOptionsModel.SetID("372e67954025e0ba6aaa6d586b9e0b59")
				updateLockdownRuleOptionsModel.SetPaused(false)
				updateLockdownRuleOptionsModel.SetDescription("Restrict access to these endpoints to requests from a known IP address")
				updateLockdownRuleOptionsModel.SetUrls([]string{"api.mysite.com/some/endpoint*"})
				updateLockdownRuleOptionsModel.SetConfigurations([]zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel})
				updateLockdownRuleOptionsModel.SetPriority(int64(5))
				updateLockdownRuleOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateLockdownRuleOptionsModel).ToNot(BeNil())
				Expect(updateLockdownRuleOptionsModel.LockdownRuleIdentifier).To(Equal(core.StringPtr("testString")))
				Expect(updateLockdownRuleOptionsModel.ID).To(Equal(core.StringPtr("372e67954025e0ba6aaa6d586b9e0b59")))
				Expect(updateLockdownRuleOptionsModel.Paused).To(Equal(core.BoolPtr(false)))
				Expect(updateLockdownRuleOptionsModel.Description).To(Equal(core.StringPtr("Restrict access to these endpoints to requests from a known IP address")))
				Expect(updateLockdownRuleOptionsModel.Urls).To(Equal([]string{"api.mysite.com/some/endpoint*"}))
				Expect(updateLockdownRuleOptionsModel.Configurations).To(Equal([]zonelockdownv1.LockdownInputConfigurationsItem{*lockdownInputConfigurationsItemModel}))
				Expect(updateLockdownRuleOptionsModel.Priority).To(Equal(core.Int64Ptr(int64(5))))
				Expect(updateLockdownRuleOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
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
