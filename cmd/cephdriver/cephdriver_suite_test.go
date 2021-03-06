package main_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"testing"
)

func TestNfsV3Driver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cephdriver Main Suite")
}

var driverPath string

var _ = BeforeSuite(func() {
	// this test suite shares an os environment and therefore cannot run in parallel
	Expect(config.GinkgoConfig.ParallelTotal).To(Equal(1))

	var err error
	driverPath, err = Build("code.cloudfoundry.org/cephdriver/cmd/cephdriver")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	CleanupBuildArtifacts()
})
