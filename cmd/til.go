package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	name string
	tags []string
)

func configureTil() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "til",
		Short: "#til generator for huyng.dev",
		Run:   tilcmd,
	}
	f := cmd.Flags()
	f.StringVarP(&name, "name", "n", "", "#til that")
	f.StringArrayVarP(&tags, "tags", "t", nil, "tags")
	return cmd
}

func tilcmd(cmd *cobra.Command, args []string) {
	// first create a file in <my_src>/<name>.md
	path := path.Join(ws, blogDir, "content/til", name+".md")

	// bootstrap with some meta info
	f, _ := os.Create(path)
	f.WriteString(tilTpl())

	// open vim to start edit

	// after save, go push to git, with after hook
}

func tilTpl() string {
	return fmt.Sprintf(`---
title: %s
kind: "til"
date: %v
tags: [%s]
---
	`, strings.Title(name), time.Now(), strings.Join(tags, ","))
}
