package cli

import (
	"github.com/gruntwork-io/terragrunt/options"
	"github.com/gruntwork-io/terragrunt/util"
)

const CMD_DEPLOY = "deploy"
const CMD_AUTH = "auth"
const CMD_GENERATE = "generate"
const CMD_ADD_DEPENDENCY = "add-dependency"
const CMD_TEST = "test"
const CMD_RELEASE = "release"
const CMD_CATALOG = "catalog"
const CMD_UPDATE = "update"
const CMD_DASHBOARD = "dashboard"
const CMD_PIPELINE = "pipeline"

var allDesignSprintCommands = []string{
	CMD_DEPLOY,
	CMD_AUTH,
	CMD_GENERATE,
	CMD_ADD_DEPENDENCY,
	CMD_TEST,
	CMD_RELEASE,
	CMD_CATALOG,
	CMD_UPDATE,
	CMD_DASHBOARD,
	CMD_PIPELINE,
}

func isDesignSprintCommand(terragruntOptions *options.TerragruntOptions) bool {
	return util.ListContainsAny(terragruntOptions.TerraformCliArgs, allDesignSprintCommands)
}

func runDesignSprintCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("**** DESIGN SPRINT CODE GOES HERE ****")
	return nil
}
