// Copyright © 2018 psucoder <hungle.info@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	workDir string
	tplDir  string
	debug   bool
	port    int
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "irr",
	Short: "Online judge helper",
	Long: `irr is an online judge helper, which helps you in parsing and testing.
'irr' stands for Is It Rated? :).`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: mainServer,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	viper.BindPFlags(rootCmd.Flags())
	cobra.OnInitialize(initConfig, initLogger)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $PWD/.irr.yaml)")

	// workDir flag
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().StringVarP(&workDir, "workDir", "w", wd, "working dir")

	// tplDir flag
	rootCmd.PersistentFlags().StringVarP(&tplDir, "tplDir", "t", path.Join(wd, "templates"), "templates dir")

	// debug flag
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "debug flag (default is false)")

	// listening port
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 4243, "listener port")
}

func initLogger() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in current working dir with name ".irr" (without extension).
		viper.AddConfigPath(".")
		// Search config in home directory with name ".irr" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".irr")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		if !rootCmd.PersistentFlags().Lookup("debug").Changed {
			debug = viper.GetBool("debug")
		}
	}
}
