package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/pdslly/go-directory-tree/tree"
)

var (
	a bool
	v bool
)

var rootCmd = &cobra.Command{
	Use: "tree <path>",
	Short: "图形显示指定文件夹的树形结构",
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if v {
			fmt.Println(tree.VERSION)
		} else {
			args = append(args, "./")
			tree.Parse(args[0], a)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&v, "version", "v", false, "显示当前版本")
	rootCmd.PersistentFlags().BoolVarP(&a, "all", "a", false, "显示每个文件夹中文件的名称，默认只显示文件夹")
}

func Execute() error {
	return rootCmd.Execute()
}
