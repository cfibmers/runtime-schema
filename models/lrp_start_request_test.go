package models_test

import (
	"encoding/json"

	. "github.com/cloudfoundry-incubator/runtime-schema/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LRPStartRequest", func() {
	var lrpStart LRPStartRequest
	var lrpStartPayload string

	BeforeEach(func() {
		lrpStartPayload = `{
    "desired_lrp": {
      "process_guid": "some-guid",
      "domain": "tests",
      "instances": 1,
      "stack": "some-stack",
      "start_timeout": 0,
      "root_fs": "docker:///docker.com/docker",
      "action": {"download": {
          "from": "http://example.com",
          "to": "/tmp/internet",
          "cache_key": ""
        }
      },
      "disk_mb": 512,
      "memory_mb": 1024,
      "cpu_weight": 42,
			"privileged": false,
      "ports": [
        5678
      ],
      "routes": [
        "route-1",
        "route-2"
      ],
      "log_guid": "log-guid",
      "log_source": "the cloud"
    },
    "indices": [2]
  }`

		lrpStart = LRPStartRequest{
			Indices: []uint{2},

			DesiredLRP: DesiredLRP{
				Domain:      "tests",
				ProcessGuid: "some-guid",

				RootFSPath: "docker:///docker.com/docker",
				Instances:  1,
				Stack:      "some-stack",
				MemoryMB:   1024,
				DiskMB:     512,
				CPUWeight:  42,
				Routes:     []string{"route-1", "route-2"},
				Ports: []uint32{
					5678,
				},
				LogGuid:   "log-guid",
				LogSource: "the cloud",
				Action: &DownloadAction{
					From: "http://example.com",
					To:   "/tmp/internet",
				},
			},
		}
	})

	Describe("ToJSON", func() {
		It("should JSONify", func() {
			jsonPayload, err := ToJSON(&lrpStart)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(string(jsonPayload)).Should(MatchJSON(lrpStartPayload))
		})
	})

	Describe("FromJSON", func() {
		var decodedLRPStartRequest *LRPStartRequest
		var err error

		JustBeforeEach(func() {
			decodedLRPStartRequest = &LRPStartRequest{}
			err = FromJSON([]byte(lrpStartPayload), decodedLRPStartRequest)
		})

		It("returns a LRP with correct fields", func() {
			Ω(err).ShouldNot(HaveOccurred())
			Ω(decodedLRPStartRequest).Should(Equal(&lrpStart))
		})

		Context("with an invalid payload", func() {
			BeforeEach(func() {
				lrpStartPayload = "aliens lol"
			})

			It("returns the error", func() {
				Ω(err).Should(HaveOccurred())
			})
		})

		Context("with an invalid desired lrp", func() {
			BeforeEach(func() {
				lrpStartPayload = `{
    "desired_lrp": {
      "domain": "tests",
      "instances": 1,
      "stack": "some-stack",
      "start_timeout": 0,
      "root_fs": "docker:///docker.com/docker",
      "action": {"download": {
          "from": "http://example.com",
          "to": "/tmp/internet",
          "cache_key": ""
        }
      },
      "disk_mb": 512,
      "memory_mb": 1024,
      "cpu_weight": 42,
      "ports": [
        5678
      ],
      "routes": [
        "route-1",
        "route-2"
      ],
      "log_guid": "log-guid",
      "log_source": "the cloud"
    },
    "index": 2
  }`
			})

			It("returns a validation error", func() {
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(ContainElement(ErrInvalidField{"process_guid"}))
			})
		})

		Context("with no indices", func() {
			BeforeEach(func() {
				lrpStartPayload = `{
    "desired_lrp": {
      "process_guid": "some-guid",
      "domain": "tests",
      "instances": 1,
      "stack": "some-stack",
      "start_timeout": 0,
      "root_fs": "docker:///docker.com/docker",
      "action": {"download": {
          "from": "http://example.com",
          "to": "/tmp/internet",
          "cache_key": ""
        }
      },
      "disk_mb": 512,
      "memory_mb": 1024,
      "cpu_weight": 42,
      "ports": [
        5678
      ],
      "routes": [
        "route-1",
        "route-2"
      ],
      "log_guid": "log-guid",
      "log_source": "the cloud"
    }
  }`
			})

			It("returns a validation error", func() {
				Ω(err).Should(HaveOccurred())
				Ω(err).Should(ContainElement(ErrInvalidField{"indices"}))
			})
		})

		Context("with an invalid index", func() {
			BeforeEach(func() {
				lrpStartPayload = `{
    "desired_lrp": {
      "process_guid": "some-guid",
      "domain": "tests",
      "instances": 1,
      "stack": "some-stack",
      "start_timeout": 0,
      "root_fs": "docker:///docker.com/docker",
      "action": {"download": {
          "from": "http://example.com",
          "to": "/tmp/internet",
          "cache_key": ""
        }
      },
      "disk_mb": 512,
      "memory_mb": 1024,
      "cpu_weight": 42,
      "ports": [
        5678
      ],
      "routes": [
        "route-1",
        "route-2"
      ],
      "log_guid": "log-guid",
      "log_source": "the cloud"
    },
    "indices": [-1]
  }`
			})

			It("returns a validation error", func() {
				Ω(err).Should(BeAssignableToTypeOf(&json.UnmarshalTypeError{}))
			})
		})

	})
})