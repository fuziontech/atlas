// Copyright 2021-present The Atlas Authors. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package cmdapi

import (
	"fmt"
	"strings"

	"ariga.io/atlas/sql/schema"
	"ariga.io/atlas/sql/sqlclient"

	"github.com/spf13/cobra"
)

type diffCmdOpts struct {
	fromURL string
	toURL   string
}

// newDiffCmd returns a new *cobra.Command that runs cmdDiffRun with the given flags and mux.
func newDiffCmd() *cobra.Command {
	var opts diffCmdOpts
	cmd := &cobra.Command{
		Use:   "diff",
		Short: "Calculate and print the diff between two schemas.",
		Long: `'atlas schema diff' connects to two given databases, inspects
them, calculates the difference in their schemas, and prints a plan of
SQL statements to migrate the "from" database to the schema of the "to" database.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmdDiffRun(cmd, &opts)
		},
	}
	cmd.Flags().StringVarP(&opts.fromURL, "from", "", "", "[driver://username:password@protocol(address)/dbname?param=value] select a database using the URL format")
	cmd.Flags().StringVarP(&opts.toURL, "to", "", "", "[driver://username:password@protocol(address)/dbname?param=value] select a database using the URL format")
	cobra.CheckErr(cmd.MarkFlagRequired("from"))
	cobra.CheckErr(cmd.MarkFlagRequired("to"))
	return cmd
}

func init() {
	diffCmd := newDiffCmd()
	schemaCmd.AddCommand(diffCmd)
}

// cmdDiffRun connects to the given databases, and prints an SQL plan to get from
// the "from" schema to the "to" schema.
func cmdDiffRun(cmd *cobra.Command, flags *diffCmdOpts) {
	ctx := cmd.Context()
	fromC, err := sqlclient.Open(cmd.Context(), flags.fromURL)
	cobra.CheckErr(err)
	defer fromC.Close()
	toC, err := sqlclient.Open(cmd.Context(), flags.toURL)
	cobra.CheckErr(err)
	defer toC.Close()
	fromS := fromC.URL.Schema
	toS := toC.URL.Schema
	var diff []schema.Change
	switch {
	case fromS == "" && toS == "":
		// compare realm.
		fromRealm, err := fromC.InspectRealm(ctx, nil)
		cobra.CheckErr(err)
		toRealm, err := toC.InspectRealm(ctx, nil)
		cobra.CheckErr(err)
		diff, err = toC.RealmDiff(fromRealm, toRealm)
		cobra.CheckErr(err)
	case fromS == "":
		cobra.CheckErr(fmt.Errorf("cannot diff schema %q with a database connection", fromS))
	case toS == "":
		cobra.CheckErr(fmt.Errorf("cannot diff database connection with a schema %q", toS))
	default:
		// compare schemas.
		fromSchema, err := fromC.InspectSchema(ctx, fromS, nil)
		cobra.CheckErr(err)
		toSchema, err := toC.InspectSchema(ctx, toS, nil)
		cobra.CheckErr(err)
		// SchemaDiff checks for name equality which is irrelevant in the case
		// the user wants to compare their contents, if the names are different
		// we reset them to allow the comparison.
		if fromS != toS {
			toSchema.Name = ""
			fromSchema.Name = ""
		}
		diff, err = toC.SchemaDiff(fromSchema, toSchema)
		cobra.CheckErr(err)
	}
	p, err := toC.PlanChanges(ctx, "plan", diff)
	cobra.CheckErr(err)
	if len(p.Changes) == 0 {
		cmd.Println("Schemas are synced, no changes to be made.")
		return
	}
	for _, c := range p.Changes {
		if c.Comment != "" {
			cmd.Println("--", strings.ToUpper(c.Comment[:1])+c.Comment[1:])
		}
		cmd.Println(c.Cmd)
	}
}
