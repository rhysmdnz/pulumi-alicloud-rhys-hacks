// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./accountPrivilege";
export * from "./backupPolicy";
export * from "./cluster";
export * from "./database";
export * from "./endpoint";
export * from "./endpointAddress";
export * from "./getAccounts";
export * from "./getClusters";
export * from "./getDatabases";
export * from "./getEndpoints";
export * from "./getGlobalDatabaseNetworks";
export * from "./getNodeClasses";
export * from "./getZones";
export * from "./globalDatabaseNetwork";

// Import resources to register:
import { Account } from "./account";
import { AccountPrivilege } from "./accountPrivilege";
import { BackupPolicy } from "./backupPolicy";
import { Cluster } from "./cluster";
import { Database } from "./database";
import { Endpoint } from "./endpoint";
import { EndpointAddress } from "./endpointAddress";
import { GlobalDatabaseNetwork } from "./globalDatabaseNetwork";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "alicloud:polardb/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "alicloud:polardb/accountPrivilege:AccountPrivilege":
                return new AccountPrivilege(name, <any>undefined, { urn })
            case "alicloud:polardb/backupPolicy:BackupPolicy":
                return new BackupPolicy(name, <any>undefined, { urn })
            case "alicloud:polardb/cluster:Cluster":
                return new Cluster(name, <any>undefined, { urn })
            case "alicloud:polardb/database:Database":
                return new Database(name, <any>undefined, { urn })
            case "alicloud:polardb/endpoint:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "alicloud:polardb/endpointAddress:EndpointAddress":
                return new EndpointAddress(name, <any>undefined, { urn })
            case "alicloud:polardb/globalDatabaseNetwork:GlobalDatabaseNetwork":
                return new GlobalDatabaseNetwork(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("alicloud", "polardb/account", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/accountPrivilege", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/backupPolicy", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/cluster", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/database", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/endpoint", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/endpointAddress", _module)
pulumi.runtime.registerResourceModule("alicloud", "polardb/globalDatabaseNetwork", _module)