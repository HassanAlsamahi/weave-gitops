package remove

// Provides support for removing an application from wego management.

import (
	"context"
	"fmt"
	"strings"

	"github.com/lithammer/dedent"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/weaveworks/weave-gitops/cmd/wego/version"
	"github.com/weaveworks/weave-gitops/pkg/apputils"
	"github.com/weaveworks/weave-gitops/pkg/services/app"
	"github.com/weaveworks/weave-gitops/pkg/utils"
)

var params app.RemoveParams

var Cmd = &cobra.Command{
	Use:   "remove [--private-key <keyfile>] <app name>",
	Short: "Remove an app from a wego cluster",
	Long: strings.TrimSpace(dedent.Dedent(`
        Removes an application from a wego cluster so it will no longer be managed via GitOps
    `)),
	Example: `
  # Remove application from wego control via immediate commit
  wego app remove podinfo
`,
	Args:          cobra.MinimumNArgs(1),
	RunE:          runCmd,
	SilenceUsage:  true,
	SilenceErrors: true,
	PostRun: func(cmd *cobra.Command, args []string) {
		version.CheckVersion(version.CheckpointParamsWithFlags(version.CheckpointParams(), cmd))
	},
}

func init() {
	Cmd.Flags().StringVar(&params.Name, "name", "", "Name of application")
	Cmd.Flags().StringVar(&params.PrivateKey, "private-key", "", "Private key to access git repository over ssh")
	Cmd.Flags().BoolVar(&params.DryRun, "dry-run", false, "If set, 'wego remove' will not make any changes to the system; it will just display the actions that would have been taken")
}

func runCmd(cmd *cobra.Command, args []string) error {
	ctx := context.Background()

	params.Name = args[0]
	params.Namespace, _ = cmd.Parent().Flags().GetString("namespace")

	appService, appError := apputils.GetAppService(ctx, params.Name, params.Namespace)
	if appError != nil {
		return fmt.Errorf("failed to create app service: %w", appError)
	}

	utils.SetCommmitMessage(fmt.Sprintf("wego app remove %s", params.Name))

	if err := appService.Remove(params); err != nil {
		return errors.Wrapf(err, "failed to remove the app %s", params.Name)
	}

	return nil
}
