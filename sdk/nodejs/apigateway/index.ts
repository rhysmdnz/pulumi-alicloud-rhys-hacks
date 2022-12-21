// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./api";
export * from "./app";
export * from "./appAttachment";
export * from "./backend";
export * from "./getApis";
export * from "./getApps";
export * from "./getBackends";
export * from "./getGroups";
export * from "./getService";
export * from "./group";
export * from "./vpcAccess";

// Import resources to register:
import { Api } from "./api";
import { App } from "./app";
import { AppAttachment } from "./appAttachment";
import { Backend } from "./backend";
import { Group } from "./group";
import { VpcAccess } from "./vpcAccess";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:apigateway/api:Api":
                return new Api(name, <any>undefined, { urn })
            case "alicloud:apigateway/app:App":
                return new App(name, <any>undefined, { urn })
            case "alicloud:apigateway/appAttachment:AppAttachment":
                return new AppAttachment(name, <any>undefined, { urn })
            case "alicloud:apigateway/backend:Backend":
                return new Backend(name, <any>undefined, { urn })
            case "alicloud:apigateway/group:Group":
                return new Group(name, <any>undefined, { urn })
            case "alicloud:apigateway/vpcAccess:VpcAccess":
                return new VpcAccess(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "apigateway/api", _module)
pulumi.runtime.registerResourceModule("alicloud", "apigateway/app", _module)
pulumi.runtime.registerResourceModule("alicloud", "apigateway/appAttachment", _module)
pulumi.runtime.registerResourceModule("alicloud", "apigateway/backend", _module)
pulumi.runtime.registerResourceModule("alicloud", "apigateway/group", _module)
pulumi.runtime.registerResourceModule("alicloud", "apigateway/vpcAccess", _module)