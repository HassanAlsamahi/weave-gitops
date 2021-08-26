package acceptance

import (
	"fmt"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sclevine/agouti"
	// . "github.com/sclevine/agouti/matchers"
)

var webDriver *agouti.WebDriver
var page *agouti.Page

func initializeUISettings() {

	By("And I install wego to my active cluster", func() {
		installAndVerifyWego(WEGO_DEFAULT_NAMESPACE)
	})

	By("And I have my default ssh key on path "+DEFAULT_SSH_KEY_PATH, func() {
		setupSSHKey(DEFAULT_SSH_KEY_PATH)
	})

	By("And run wego ui", func() {
		_ = runCommandAndReturnSessionOutput(fmt.Sprintf("%s ui run &", WEGO_BIN_PATH))
	})
}

var _ = Describe("Weave GitOps UI Test", func() {

	deleteWegoRuntime := false
	if os.Getenv("DELETE_WEGO_RUNTIME_ON_EACH_TEST") == "true" {
		deleteWegoRuntime = true
	}

	BeforeEach(func() {
		By("Given I have all the setup ready", func() {
			var err error
			_, err = ResetOrCreateCluster(WEGO_DEFAULT_NAMESPACE, deleteWegoRuntime)
			Expect(err).ShouldNot(HaveOccurred())

			Expect(FileExists(WEGO_BIN_PATH)).To(BeTrue())
			initializeUISettings()

			webDriver = initializeWebDriver("remote")

			By("When I navigate to WeGO dashboard", func() {
				page, err = webDriver.NewPage(agouti.Desired(agouti.Capabilities{}.Browser("chrome")))
				Expect(err).NotTo(HaveOccurred())
			})

		})
	})

	AfterEach(func() {
		Expect(page.Destroy()).To(Succeed())
	})

	It("SmokeTest - Verify wego can run UI without apps installed", func() {
		Expect(page.Navigate(WEGO_UI_URL)).To(Succeed())
		Expect(page.Title()).To(ContainSubstring("Weave GitOps"))
		time.Sleep(5 * time.Second)
		fmt.Println(page.Title())
		installAndVerifyWego(WEGO_DEFAULT_NAMESPACE)
	})
})
