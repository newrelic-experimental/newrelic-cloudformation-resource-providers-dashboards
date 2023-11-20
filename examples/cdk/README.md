# CDK Example

This is an example of setting up a basic dashboard using CDK.
It leverages the NewRelic::Observability::Dashboards plugin in CloudFormation for its deploy.

# Usage

- Begin with `npm i` in this directory
- Log into your AWS on the command line. For simplicity, this example assumes you have set up a `~/.aws/config` file and have exported the proper profile as an `AWS_PROFILE` environment variable. If that is not right for you, you can alter the `deploy`, `destroy`, and `diff` commands in `package.json` to have the correct `cdk` flags for your setup.
- To view just the CloudFormation json created, run `npm run synth`. Json will be output in the `cdk.out` directory.
- To deploy the stack in the environment you're logged in to, run `npm run deploy`
- To completely remove the stack, run `npm run destroy`.
