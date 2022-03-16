package cli

import (
	"fmt"
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
	switch terragruntOptions.TerraformCommand {
	case CMD_DEPLOY:
		return runDeployCommand(terragruntOptions)
	case CMD_AUTH:
		return runAuthCommand(terragruntOptions)
	case CMD_GENERATE:
		return runGenerateCommand(terragruntOptions)
	case CMD_ADD_DEPENDENCY:
		return runAddDependencyCommand(terragruntOptions)
	case CMD_TEST:
		return runTestCommand(terragruntOptions)
	case CMD_RELEASE:
		return runReleaseCommand(terragruntOptions)
	case CMD_CATALOG:
		return runCatalogCommand(terragruntOptions)
	case CMD_UPDATE:
		return runUpdateCommand(terragruntOptions)
	case CMD_DASHBOARD:
		return runDashboardCommand(terragruntOptions)
	case CMD_PIPELINE:
		return runPipelineCommand(terragruntOptions)
	default:
		return fmt.Errorf("unrecognized command: %s", terragruntOptions.TerraformCommand)
	}
}

func runDeployCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runAuthCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runGenerateCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runAddDependencyCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runTestCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runReleaseCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runCatalogCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runUpdateCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runDashboardCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

func runPipelineCommand(terragruntOptions *options.TerragruntOptions) error {
	terragruntOptions.Logger.Infof("Command '%s' is not yet implemented.", terragruntOptions.TerraformCommand)
	return nil
}

