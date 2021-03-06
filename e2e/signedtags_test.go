package e2e

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/ossf/scorecard/checker"
	"github.com/ossf/scorecard/checks"
)

var _ = Describe("E2E TEST:Signedtags", func() {
	Context("E2E TEST:Validating signed tags", func() {
		It("Should return valid signed tags", func() {
			l := log{}
			checker := checker.Checker{
				Ctx:         context.Background(),
				Client:      ghClient,
				HttpClient:  client,
				Owner:       "bitcoin",
				Repo:        "bitcoin",
				GraphClient: graphClient,
				Logf:        l.Logf,
			}
			result := checks.SignedTags(checker)
			Expect(result.Error).Should(BeNil())
			Expect(result.Pass).Should(BeTrue())
		})
	})
})
