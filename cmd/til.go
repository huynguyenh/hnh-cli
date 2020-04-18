package cmd

import (
	"fmt"
	"os"
	"os/exec"
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
	p := path.Join(ws, blogDir, "content/til", name+".md")

	// bootstrap with some meta info
	f, _ := os.Create(p)
	f.WriteString(tilTpl())

	// open vim to start edit
	vim := exec.Command("vim", p)
	vim.Stdin = os.Stdin
	vim.Stdout = os.Stdout
	vim.Run()

	// run git push script
	exe := exec.Command("sh", path.Dir(p)+"/til.sh", "add "+name+".md")
	exe.Stdout = os.Stdout
	exe.Run()
}

func tilTpl() string {
	return fmt.Sprintf(`---
title: %s
kind: "til"
date: %v
tags: [%s]
---`, strings.Title(name), time.Now(), strings.Join(tags, ","))
}
