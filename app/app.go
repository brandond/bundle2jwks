package app

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/brandond/bundle2jwks/pkg/util"
	"github.com/brandond/bundle2jwks/pkg/version"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func init() {
	// Add a trailing newline after the help to separate it from error output
	cli.AppHelpTemplate = cli.AppHelpTemplate + "\n"
}

func New() *cli.App {
	return &cli.App{
		Name:            "bundle2jwks",
		Usage:           "Convert an x509 CA bundle to go-jose JSONWebKeySet",
		Version:         fmt.Sprintf("%s (%.8s)", version.GitVersion, version.GitCommit),
		HideHelpCommand: true,
		Args:            true,
		ArgsUsage:       "[CA-BUNDLE-FILE]",
		Writer:          os.Stderr,
		ErrWriter:       os.Stderr,
		Before: func(clx *cli.Context) error {
			if !clx.Args().Present() {
				cli.ShowAppHelp(clx)
				return fmt.Errorf("missing required arg: CA-BUNDLE-FILE")
			}
			return nil
		},
		Action: func(clx *cli.Context) error {
			ks, err := util.GetKeySet(clx.Args().First())
			if err != nil {
				return errors.Wrap(err, "failed to get JSONWebKeySet")
			}

			b, err := json.MarshalIndent(ks, "", "  ")
			if err != nil {
				return errors.Wrap(err, "failed to marshal JSONWebKeySet")
			}

			fmt.Fprintf(os.Stdout, "%s\n", string(b))
			return nil
		},
	}
}
