package add

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"

	"github.com/weaveworks/weave-gitops/pkg/fluxops"
)

const (
	manifestsKustomize = "kustomize"
	manifestsHelm      = "helm"
)

type paramSet struct {
	name          string
	url           string
	branch        string
	path          string
	manifestsKind string
}

var params paramSet

var Cmd = &cobra.Command{
	Use:   "add [--name <name>] [--url <url>] [--branch <branch>] [--path <path>] [--manifests-kind kustomize|helm]",
	Short: "Add a workload repository to a wego cluster",
	Long: strings.TrimSpace(dedent.Dedent(`
        Associates an additional git repository with a wego cluster so that its contents may be managed via GitOps
    `)),
	Example: "wego add",
	Run:     runCmd,
}

// checkError will print a message to stderr and exit
func checkError(msg string, err interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", msg, err)
		os.Exit(1)
	}
}

func checkAddError(err interface{}) {
	checkError("Failed to add workload repository", err)
}

func init() {
	Cmd.Flags().StringVar(&params.name, "name", "", "Name of remote git repository")
	Cmd.Flags().StringVar(&params.url, "url", "", "URL of remote git repository")
	Cmd.Flags().StringVar(&params.path, "path", "", "Path to watch within git repository")
	Cmd.Flags().StringVar(&params.branch, "branch", "main", "Branch to watch within git repository")
	Cmd.Flags().StringVar(&params.manifestsKind, "manifests-kind", manifestsKustomize, "Manifests kind used to generate the templates (kustomize or helm)")

	// Required
	Cmd.MarkFlagRequired("path")
}

func updateURLIfNecessary() {
	if params.url == "" {
		url, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
		checkError("Failed to discover URL of remote repository", err)
		fmt.Printf("URL not specified; using URL of 'origin' from git config...")
		params.url = string(url)
	}
}

func generateSourceManifest() []byte {
	sourceManifest, err := fluxops.CallFlux(fmt.Sprintf(`create source git %s --url="%s" --branch="%s" --interval=30s --export`,
		params.name, params.url, params.branch))
	checkAddError(err)
	return sourceManifest
}

func generateKustomizeManifest() []byte {
	kustomizeManifest, err := fluxops.CallFlux(
		fmt.Sprintf(`create kustomization %s --path="%s" --prune=true --validation=client --interval=5m --export`, params.name, params.path))
	checkAddError(err)
	return kustomizeManifest
}

func generateHelmManifest() []byte {
	kustomizeManifest, err := fluxops.CallFlux(
		fmt.Sprintf(`create helmrelease %s --source="GitRepository/%s" --chart="%s" --interval=5m --export`, params.name, params.name, params.path))

	checkAddError(err)
	return kustomizeManifest
}

func runCmd(cmd *cobra.Command, args []string) {
	updateURLIfNecessary()

	generateTemplates := generateKustomizeManifest

	if params.manifestsKind == manifestsHelm {
		generateTemplates = generateHelmManifest
	}

	io.WriteString(os.Stdout, fmt.Sprintf("%s%s", generateSourceManifest(), generateTemplates()))
}
