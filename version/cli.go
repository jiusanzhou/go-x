/*
 * Copyright (c) 2020 wellwell.work, LLC by Zoe
 *
 * Licensed under the Apache License 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package version

import "go.zoe.im/x/cli"

var (
	helpTempl = `Version: %s
BuildDate: %s
GitCommit: %s
GitTreeState: %s
GoVersion: %s
Compiler: %s
Platform: %s
`
)

// NewCommand return a new command for version
func NewCommand(opts ...cli.Option) *cli.Command {
	var c = &struct{ Short bool }{}
	return cli.New(
		cli.Name("version"),
		cli.Short("Print the application version information"),
		cli.Config(c),
		cli.Run(func(cmd *cli.Command, args ...string) {
			i := Get()
			if c.Short {
				fmt.Println(i.GitVersion)
				return
			}

			// print detail.
			fmt.Printf(
				helpTempl,
				i.GitVersion,
				i.BuildDate, i.GitCommit, i.GitTreeState,
				i.GoVersion, i.Compiler, i.Platform,
			)
		})
		...opts,
	)
}

// NewOption return a option to set version
func NewOption(needCmd bool, opts ...Option) Option {
	return func(c *Command) {
		c.Command.Version = gitVersion
		// if needCmd is true, we need to install a version command
		if isCmd {
			c.Register(NewCommand(opts...))
		}
	}
}