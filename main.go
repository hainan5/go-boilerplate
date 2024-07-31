package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-boilerplate/app/cmd"
	"go-boilerplate/bootstrap"
	btsConig "go-boilerplate/config"
	"go-boilerplate/pkg/config"
	"go-boilerplate/pkg/console"
	"os"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConig.Initialize()
}

func main() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   "go-boilerplate",
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)
			fmt.Println("初始化命令成功")

			// 初始化 Logger
			bootstrap.SetupLogger()
			fmt.Println("初始化日志成功")

			// 初始化数据库
			bootstrap.SetupDB()
			fmt.Println("初始化数据库成功")

			// 初始化 Redis
			bootstrap.SetupRedis()
			fmt.Println("初始化Redis成功")

			// 初始化缓存
			bootstrap.SetupCache()
			fmt.Println("初始化缓存成功")
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdCache,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
