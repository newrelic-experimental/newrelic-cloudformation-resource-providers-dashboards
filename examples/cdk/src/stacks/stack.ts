import { Stack, StackProps, CfnOutput } from "aws-cdk-lib";
import { Construct } from "constructs";
import { CfnDashboards } from "@cdk-cloudformation/newrelic-observability-dashboards";

export class DashboardStack extends Stack {
  constructor(scope: Construct, id: string, props: StackProps = {}) {
    super(scope, id, props);

    const buildADashboard = {
      dashboard: {
        description: "CloudFormation test dashboard",
        name: "CloudFormation Test Dashboard",
        pages: {
          description: "TD PAGE",
          name: "TD Page",
          widgets: {
            title: "Widget Title",
            configuration: { markdown: { text: "Some Markdown" } },
          },
        },
        permissions: "PUBLIC_READ_ONLY",
      },
    };

    const dash = new CfnDashboards(this, "NewRelicDashboard", {
      dashboard: this.toGqlString(buildADashboard),
    });
    new CfnOutput(this, "NewRelicDashboardGuid", {
      exportName: "NewRelicDashCDKExampleGuid",
      value: dash.attrGuid,
    });
  }

  // This is a workaround to get a json-ified object into the correct GQL format
  // There are GQL libraries that will do this for you, and you should probably use those.
  // However, for a minimal example, I felt it best to go with fewer libraries.
  private toGqlString(obj: any): string {
    const firstPass = JSON.stringify(obj).replace(/\"([\w_-]+?)\"\:/g, "$1:");
    const permissionsUpdated = firstPass.replace(
      /permissions:\"([A-Z_]+)\"/g,
      "permissions:$1"
    );
    const removeOuterBrace = permissionsUpdated
      .substring(0, permissionsUpdated.lastIndexOf("}"))
      .replace("{", "");
    return removeOuterBrace;
  }
}
