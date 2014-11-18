package lrp_bbs_test

import (
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/shared"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/cloudfoundry/dropsonde/metric_sender/fake"
	"github.com/cloudfoundry/dropsonde/metrics"
	"github.com/cloudfoundry/storeadapter"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LrpConvergence", func() {
	var (
		sender *fake.FakeMetricSender

		cellID string
	)

	processGuid := "process-guid"

	BeforeEach(func() {
		sender = fake.NewFakeMetricSender()
		metrics.Initialize(sender)

		cellID = "the-cell-id"
		etcdClient.Create(storeadapter.StoreNode{
			Key:   shared.CellSchemaPath(cellID),
			Value: []byte{},
		})

		_, err := bbs.ReportActualLRPAsStarting(processGuid, "instance-guid-1", cellID, "domain", 0)
		Ω(err).ShouldNot(HaveOccurred())
		_, err = bbs.ReportActualLRPAsStarting(processGuid, "instance-guid-2", cellID, "domain", 1)
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("bumps the convergence counter", func() {
		Ω(sender.GetCounter("ConvergenceLRPRuns")).Should(Equal(uint64(0)))
		bbs.ConvergeLRPs()
		Ω(sender.GetCounter("ConvergenceLRPRuns")).Should(Equal(uint64(1)))
		bbs.ConvergeLRPs()
		Ω(sender.GetCounter("ConvergenceLRPRuns")).Should(Equal(uint64(2)))
	})

	It("reports the duration that it took to converge", func() {
		bbs.ConvergeLRPs()

		reportedDuration := sender.GetValue("ConvergenceLRPDuration")
		Ω(reportedDuration.Unit).Should(Equal("nanos"))
		Ω(reportedDuration.Value).ShouldNot(BeZero())
	})

	Describe("pruning LRPs by cell", func() {
		JustBeforeEach(func() {
			bbs.ConvergeLRPs()
		})

		Context("when no cell is missing", func() {
			It("should not prune any LRPs", func() {
				Ω(bbs.ActualLRPs()).Should(HaveLen(2))
			})
		})

		Context("when an cell is missing", func() {
			BeforeEach(func() {
				etcdClient.Delete(shared.CellSchemaPath(cellID))
			})

			It("should delete LRPs associated with said cell", func() {
				Ω(bbs.ActualLRPs()).Should(BeEmpty())
			})

			It("should prune LRP directories for apps that are no longer running", func() {
				actual, err := etcdClient.ListRecursively(shared.ActualLRPSchemaRoot)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(actual.ChildNodes).Should(BeEmpty())
			})
		})
	})

	Describe("when there is a desired LRP", func() {
		var desiredEvents <-chan models.DesiredLRPChange
		var desiredLRP models.DesiredLRP

		commenceWatching := func() {
			desiredEvents, _, _ = bbs.WatchForDesiredLRPChanges()
		}

		BeforeEach(func() {
			desiredLRP = models.DesiredLRP{
				Domain:      "tests",
				ProcessGuid: processGuid,
				Instances:   2,
				Stack:       "pancake",
				Action: models.ExecutorAction{
					Action: models.DownloadAction{
						From: "http://example.com",
						To:   "/tmp/internet",
					},
				},
			}
		})

		Context("when the desired LRP has malformed JSON", func() {
			BeforeEach(func() {
				err := etcdClient.SetMulti([]storeadapter.StoreNode{
					{
						Key:   shared.DesiredLRPSchemaPathByProcessGuid("bogus-desired"),
						Value: []byte("ß"),
					},
				})

				Ω(err).ShouldNot(HaveOccurred())
			})

			It("should delete the bogus entry", func() {
				bbs.ConvergeLRPs()
				_, err := etcdClient.Get(shared.DesiredLRPSchemaPathByProcessGuid("bogus-desired"))
				Ω(err).Should(MatchError(storeadapter.ErrorKeyNotFound))
			})

			It("bumps the deleted LRPs convergence counter", func() {
				Ω(sender.GetCounter("ConvergenceLRPsDeleted")).Should(Equal(uint64(0)))
				bbs.ConvergeLRPs()
				Ω(sender.GetCounter("ConvergenceLRPsDeleted")).Should(Equal(uint64(1)))
			})
		})

		Context("when the desired LRP has all its actual LRPs, and there are no extras", func() {
			BeforeEach(func() {
				bbs.DesireLRP(desiredLRP)
			})

			It("should not kick the desired LRP", func() {
				commenceWatching()
				bbs.ConvergeLRPs()

				Consistently(desiredEvents).ShouldNot(Receive())
			})
		})

		Context("when the desired LRP is missing actuals", func() {
			BeforeEach(func() {
				desiredLRP.Instances = 3
				bbs.DesireLRP(desiredLRP)
			})

			It("should kick the desired LRP", func() {
				commenceWatching()
				bbs.ConvergeLRPs()

				var noticedOnce models.DesiredLRPChange
				Eventually(desiredEvents).Should(Receive(&noticedOnce))
				Ω(*noticedOnce.After).Should(Equal(desiredLRP))
			})

			It("bumps the compare-and-swapped LRPs convergence counter", func() {
				Ω(sender.GetCounter("ConvergenceLRPsKicked")).Should(Equal(uint64(0)))
				bbs.ConvergeLRPs()
				Ω(sender.GetCounter("ConvergenceLRPsKicked")).Should(Equal(uint64(1)))
			})
		})

		Context("when there are extra actual LRPs", func() {
			BeforeEach(func() {
				desiredLRP.Instances = 1
				bbs.DesireLRP(desiredLRP)
			})

			It("should kick the desired LRP", func() {
				commenceWatching()
				bbs.ConvergeLRPs()

				var noticedOnce models.DesiredLRPChange
				Eventually(desiredEvents).Should(Receive(&noticedOnce))
				Ω(*noticedOnce.After).Should(Equal(desiredLRP))
			})

			It("bumps the compare-and-swapped LRPs convergence counter", func() {
				Ω(sender.GetCounter("ConvergenceLRPsKicked")).Should(Equal(uint64(0)))
				bbs.ConvergeLRPs()
				Ω(sender.GetCounter("ConvergenceLRPsKicked")).Should(Equal(uint64(1)))
			})
		})

		Context("when there are duplicate actual LRPs", func() {
			BeforeEach(func() {
				bbs.ReportActualLRPAsStarting(processGuid, "instance-guid-duplicate", cellID, "domain", 0)
				bbs.DesireLRP(desiredLRP)
			})

			It("should kick the desired LRP", func() {
				commenceWatching()
				bbs.ConvergeLRPs()

				var noticedOnce models.DesiredLRPChange
				Eventually(desiredEvents).Should(Receive(&noticedOnce))
				Ω(*noticedOnce.After).Should(Equal(desiredLRP))
			})

			It("bumps the compare-and-swapped LRPs convergence counter", func() {
				Ω(sender.GetCounter("ConvergenceLRPsKicked")).Should(Equal(uint64(0)))
				bbs.ConvergeLRPs()
				Ω(sender.GetCounter("ConvergenceLRPsKicked")).Should(Equal(uint64(1)))
			})
		})
	})

	Context("when there is an actual LRP with no matching desired LRP", func() {
		It("should emit a stop for the actual LRP", func() {
			bbs.ConvergeLRPs()
			stops, err := bbs.StopLRPInstances()
			Ω(err).ShouldNot(HaveOccurred())
			Ω(stops).Should(HaveLen(2))

			Ω(stops).Should(ContainElement(models.StopLRPInstance{
				ProcessGuid:  processGuid,
				InstanceGuid: "instance-guid-1",
				Index:        0,
			}))

			Ω(stops).Should(ContainElement(models.StopLRPInstance{
				ProcessGuid:  processGuid,
				InstanceGuid: "instance-guid-2",
				Index:        1,
			}))
		})

		It("bumps the stopped LRPs convergence counter", func() {
			Ω(sender.GetCounter("ConvergenceLRPsStopped")).Should(Equal(uint64(0)))
			bbs.ConvergeLRPs()
			Ω(sender.GetCounter("ConvergenceLRPsStopped")).Should(Equal(uint64(2)))
		})
	})
})
