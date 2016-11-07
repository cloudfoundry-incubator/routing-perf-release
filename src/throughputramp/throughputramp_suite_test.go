package main_test

import (
	"os/exec"
	"strconv"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
	"github.com/tedsuo/ifrit/ginkgomon"

	"testing"
)

var (
	binPath string
)

func NewThroughputRamp(throughputRampPath string, args Args) *ginkgomon.Runner {
	return ginkgomon.New(ginkgomon.Config{
		Name:    "throughputramp",
		Command: exec.Command(throughputRampPath, args.ArgSlice()...),
	})
}

type Args struct {
	NumRequests        int
	ConcurrentRequests int
	StartRateLimit     int
	EndRateLimit       int
	RateLimitStep      int
	URL                string
}

func (args Args) ArgSlice() []string {
	return []string{
		"-n", strconv.Itoa(args.NumRequests),
		"-c", strconv.Itoa(args.ConcurrentRequests),
		"--lower-throughput", strconv.Itoa(args.StartRateLimit),
		"--upper-throughput", strconv.Itoa(args.EndRateLimit),
		"--throughput-step", strconv.Itoa(args.RateLimitStep),
		args.URL,
	}
}

func TestThroughputramp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Throughputramp Suite")
}

var _ = BeforeSuite(func() {
	var err error
	binPath, err = gexec.Build("throughputramp", "-race")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})
