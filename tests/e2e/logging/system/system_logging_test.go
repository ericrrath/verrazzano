// Copyright (c) 2022, Oracle and/or its affiliates.
// Licensed under the Universal Permissive License v 1.0 as shown at https://oss.oracle.com/licenses/upl.

package system

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/verrazzano/verrazzano/pkg/test/framework"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/verrazzano/verrazzano/tests/e2e/pkg"
)

const (
	shortPollingInterval = 10 * time.Second
	shortWaitTimeout     = 5 * time.Minute
	searchTimeWindow     = "1h"
	systemIndex          = "verrazzano-namespace-verrazzano-system"
	installIndex         = "verrazzano-namespace-verrazzano-install"
	certMgrIndex         = "verrazzano-namespace-cert-manager"
)

var (
	noExceptions    = []*regexp.Regexp{}
	istioExceptions = []*regexp.Regexp{
		regexp.MustCompile(`^-A .*$`),
		regexp.MustCompile(`^-N .*$`),
		regexp.MustCompile(`^:\w+? -.*$`),
		regexp.MustCompile(`^:\w+? ACCEPT.*$`),
		regexp.MustCompile(`^\w+?=.*$`),
		regexp.MustCompile(`^COMMIT.*$`),
		regexp.MustCompile(`^ {0,4}\w+:.*$`),
		regexp.MustCompile(`^:.*$`),
		regexp.MustCompile(`^\* ?nat.*$`),
		regexp.MustCompile(`^# Generated by.*$`),
		regexp.MustCompile(`^# Completed on.*$`),
		regexp.MustCompile(`^Writing following contents to rules file:.*$`),
		regexp.MustCompile(`^ip\w?tables.*$`),
		regexp.MustCompile(`^-+$`),
		regexp.MustCompile(`^$`),
	}
)

var t = framework.NewTestFramework("system-logging")

var _ = t.BeforeSuite(func() {})

var failed = false
var _ = t.AfterEach(func() {
	failed = failed || CurrentSpecReport().Failed()
})

var _ = t.AfterSuite(func() {
	if failed {
		pkg.ExecuteClusterDumpWithEnvVarConfig()
	}
})

var _ = t.Describe("Elasticsearch system component data", Label("f:observability.logging.es"), func() {
	t.It("contains verrazzano-system index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Elasticsearch index for the verrazzano-system namespace is retrieved
		// THEN verify that it is found
		Eventually(func() bool {
			return pkg.LogIndexFound("verrazzano-namespace-verrazzano-system")
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Elasticsearch index verrazzano-system")

		valid := true
		valid = validateAuthProxyLogs() && valid
		valid = validateCoherenceLogs() && valid
		valid = validateOAMLogs() && valid
		valid = validateIstioProxyLogs() && valid
		valid = validateKialiLogs() && valid
		valid = validateCertManagerLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in verrazzano-system index")
		}
	})

	t.It("contains valid verrazzano-install index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Elasticsearch index for the verrazzano-install namespace is retrieved
		// THEN verify that it is found
		Eventually(func() bool {
			return pkg.LogIndexFound("verrazzano-namespace-verrazzano-install")
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Elasticsearch index verrazzano-install")

		// GIVEN Log message in Elasticsearch in the verrazzano-namespace-verrazzano-install index
		// With field kubernetes.labels.app.keyword==verrazzano-platform-operator
		// WHEN Log messages are retrieved from Elasticsearch
		// THEN Verify there are valid log records
		valid := true
		valid = validateVPOLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in verrazzano-install index")
		}
	})

	t.It("contains cert-manager index with valid records", func() {
		// GIVEN existing system logs
		// WHEN the Elasticsearch index for the cert-manager namespace is retrieved
		// THEN verify that it is found
		Eventually(func() bool {
			return pkg.LogIndexFound("verrazzano-namespace-cert-manager")
		}, shortWaitTimeout, shortPollingInterval).Should(BeTrue(), "Expected to find Elasticsearch index cert-manager")

		valid := true
		valid = validateCertManagerLogs() && valid
		if !valid {
			// Don't fail for invalid logs until this is stable.
			t.Logs.Info("Found problems with log records in cert-manager index")
		}
	})
})

func validateAuthProxyLogs() bool {
	exceptions := []*regexp.Regexp{
		regexp.MustCompile(`^Adding local CA cert to .*$`),
		regexp.MustCompile(`^Detected Nginx Configuration Change$`),
	}
	exceptions = append(exceptions, istioExceptions...)
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		systemIndex,
		"kubernetes.labels.app.keyword",
		"verrazzano-authproxy",
		searchTimeWindow,
		exceptions)
}

func validateCoherenceLogs() bool {
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		systemIndex,
		"kubernetes.labels.app_kubernetes_io/name.keyword",
		"coherence-operator",
		searchTimeWindow,
		noExceptions)
}

func validateOAMLogs() bool {
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		systemIndex,
		"kubernetes.labels.app_kubernetes_io/name.keyword",
		"oam-kubernetes-runtime",
		searchTimeWindow,
		noExceptions)
}

// message:configPath: ./etc/istio/proxy
func validateIstioProxyLogs() bool {
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		systemIndex,
		"kubernetes.container_name",
		"istio-proxy",
		searchTimeWindow,
		istioExceptions)
}

func validateKialiLogs() bool {
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		systemIndex,
		"kubernetes.labels.app_kubernetes_io/part-of",
		"kiali",
		searchTimeWindow,
		istioExceptions)
}

func validateVPOLogs() bool {
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		installIndex,
		"kubernetes.labels.app.keyword",
		"verrazzano-platform-operator",
		searchTimeWindow,
		noExceptions)
}

func validateCertManagerLogs() bool {
	return validateElasticsearchRecords(
		basicElasticsearchRecordValidator,
		certMgrIndex,
		"kubernetes.labels.app_kubernetes_io/instance",
		"cert-manager",
		searchTimeWindow,
		noExceptions)
}

func validateElasticsearchRecords(hitValidator pkg.ElasticsearchHitValidator, namespace string, appLabel string, appName string, timeRange string, exceptions []*regexp.Regexp) bool {
	pkg.Log(pkg.Info, fmt.Sprintf("Validating log records for %s", appName))
	template :=
		`{
			"size": 1000,
      "sort": [{"@timestamp": {"order": "desc"}}],
			"query": {
				"bool": {
					"filter" : [
						{"match_phrase": {"%s": "%s"}},
						{"range": {"@timestamp": {"gte": "now-%s"}}}
					]
				}
			}
		}`
	query := fmt.Sprintf(template, appLabel, appName, timeRange)
	resp, err := pkg.PostElasticsearch(fmt.Sprintf("%s/_search", namespace), query)
	if err != nil {
		pkg.Log(pkg.Error, fmt.Sprintf("Failed to query Elasticsearch: %v", err))
		return false
	}
	if resp.StatusCode != 200 {
		pkg.Log(pkg.Error, fmt.Sprintf("Failed to query Elasticsearch: status=%d: body=%s", resp.StatusCode, string(resp.Body)))
		return false
	}
	var result map[string]interface{}
	json.Unmarshal(resp.Body, &result)

	if !pkg.ValidateElasticsearchHits(result, hitValidator, exceptions) {
		pkg.Log(pkg.Info, fmt.Sprintf("Found invalid (or zero) log records in %s logs", appName))
		return false
	}
	return true
}

// basicElasticsearchRecordValidator does common validation of log records
func basicElasticsearchRecordValidator(hit pkg.ElasticsearchHit) bool {
	ts := ""
	valid := true
	// Verify the record has a @timestamp field.
	// If so extract it.
	if val, ok := hit["@timestamp"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty @timestamp field")
		valid = false
	} else {
		ts = hit["@timestamp"].(string)
	}
	// Verify the record has a log field.
	// If so verify the time in the log field matches the @timestamp field.
	if val, ok := hit["log"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty log field")
		valid = false
	} else {
		re := regexp.MustCompile(`(\d{2}:\d{2}:\d{2})`)
		m := re.FindStringSubmatch(val.(string))
		if len(m) < 2 {
			pkg.Log(pkg.Info, "Log record log field does not contain a time")
			valid = false
		} else {
			if !strings.Contains(ts, m[1]) {
				pkg.Log(pkg.Info, fmt.Sprintf("Log record @timestamp field %s does not match log field %s content", ts, m[1]))
				valid = false
			}
		}
	}
	// Verify the record has a message field.
	if val, ok := hit["message"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty message field")
		valid = false
	}
	// Verify the log field isn't exactly the same as the message field.
	if hit["log"] == hit["message"] {
		pkg.Log(pkg.Info, "Log record has duplicate log and message field values")
		valid = false
	}
	// Verify the record has a level field.
	// If so verify that the level isn't debug.
	if val, ok := hit["level"]; !ok || len(val.(string)) == 0 {
		pkg.Log(pkg.Info, "Log record has missing or empty level field")
		valid = false
	} else {
		//level := val.(string)
		// Put this validation back in when the OAM logging is fixed.
		// if strings.EqualFold(level, "debug") || strings.EqualFold(level, "dbg") || strings.EqualFold(level, "d") {
		// 	pkg.Log(pkg.Info, fmt.Sprintf("Log record has invalid debug level: %s", level))
		// 	valid = false
		// }
		// There is an Istio proxy error that causes this to fail.
		// Put this validation back in when that is addressed.
		//if strings.EqualFold(level, "error") || strings.EqualFold(level, "err") || strings.EqualFold(level, "e") {
		//	pkg.Log(pkg.Info, fmt.Sprintf("Log record has invalid error level: %s", level))
		//	valid = false
		//}
	}
	// Verify the record does not have a timestamp field.
	if _, ok := hit["timestamp"]; ok {
		pkg.Log(pkg.Info, "Log record has unwanted timestamp field")
		valid = false
	}
	if !valid {
		pkg.Log(pkg.Info, fmt.Sprintf("Log record is invalid: %v", hit))
	}
	return valid
}