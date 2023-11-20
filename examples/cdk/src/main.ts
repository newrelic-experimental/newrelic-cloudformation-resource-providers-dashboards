import { App } from "aws-cdk-lib";

import { DashboardStack } from "./stacks/stack";

const app = new App();

new DashboardStack(app, "newrelic-dashboard-cdk-example", {});

app.synth();
