package stats_test

import (
	"cpumonitor/fakes"
	"cpumonitor/stats"
	"errors"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CPUCollector", func() {
	Describe("Run", func() {
		var defaultTestTicker stats.TickerHarness

		BeforeEach(func() {
			defaultTestTicker = func(d time.Duration) <-chan time.Time {
				t := time.NewTicker(d)
				return t.C
			}
		})

		It("returns an error if cpu percent errors", func() {
			fakeCpu := new(fakes.FakeCpu)
			expectedErr := errors.New("something bad happened")
			fakeCpu.PercentReturns(nil, expectedErr)
			collector := stats.NewCPUCollector(fakeCpu, defaultTestTicker, time.Second, time.Second, false)

			Eventually(func() error { return collector.Run() }()).Should(Equal(expectedErr))
		})

		It("uses specified cpu arguments", func() {
			fakeCpu := new(fakes.FakeCpu)

			ch := make(chan time.Time)
			testTickerHarness := func(d time.Duration) <-chan time.Time {
				return ch
			}

			cpuInterval := time.Second
			perCPU := false
			collector := stats.NewCPUCollector(fakeCpu, testTickerHarness, time.Second, cpuInterval, perCPU)
			go collector.Run()

			ch <- time.Time{}

			interval, perCpu := fakeCpu.PercentArgsForCall(0)
			Expect(interval).To(Equal(cpuInterval))
			Expect(perCpu).To(Equal(perCPU))
		})
	})

	Describe("Result", func() {
		It("returns a slice of Info", func() {

			fakeCpu := new(fakes.FakeCpu)
			stat := []float64{1.1, 2.2}
			fakeCpu.PercentReturns(stat, nil)

			ch := make(chan time.Time)
			testTickerHarness := func(d time.Duration) <-chan time.Time {
				return ch
			}

			collector := stats.NewCPUCollector(fakeCpu, testTickerHarness, time.Second, time.Second, false)
			go collector.Run()

			ch <- time.Time{}

			result := collector.Result()
			Expect(len(result)).ToNot(Equal(0))
			for i := 0; i < len(result); i++ {
				Expect(result[i].Percentage).To(Equal(stat))
			}
		})
	})
})
