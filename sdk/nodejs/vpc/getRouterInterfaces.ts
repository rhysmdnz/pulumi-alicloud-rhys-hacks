// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as alicloud from "@pulumi/alicloud";
 *
 * const routerInterfacesDs = pulumi.output(alicloud.vpc.getRouterInterfaces({
 *     nameRegex: "^testenv",
 *     status: "Active",
 * }));
 *
 * export const firstRouterInterfaceId = routerInterfacesDs.interfaces[0].id;
 * ```
 */
export function getRouterInterfaces(args?: GetRouterInterfacesArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterInterfacesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("alicloud:vpc/getRouterInterfaces:getRouterInterfaces", {
        "ids": args.ids,
        "nameRegex": args.nameRegex,
        "oppositeInterfaceId": args.oppositeInterfaceId,
        "oppositeInterfaceOwnerId": args.oppositeInterfaceOwnerId,
        "outputFile": args.outputFile,
        "role": args.role,
        "routerId": args.routerId,
        "routerType": args.routerType,
        "specification": args.specification,
        "status": args.status,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouterInterfaces.
 */
export interface GetRouterInterfacesArgs {
    /**
     * A list of router interface IDs.
     */
    ids?: string[];
    /**
     * A regex string used to filter by router interface name.
     */
    nameRegex?: string;
    /**
     * ID of the peer router interface.
     */
    oppositeInterfaceId?: string;
    /**
     * Account ID of the owner of the peer router interface.
     */
    oppositeInterfaceOwnerId?: string;
    outputFile?: string;
    /**
     * Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and 
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `routerType` is set to `VBR`.
     */
    role?: string;
    /**
     * ID of the VRouter located in the local region.
     */
    routerId?: string;
    /**
     * Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     */
    routerType?: string;
    /**
     * Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     */
    specification?: string;
    /**
     * Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     */
    status?: string;
}

/**
 * A collection of values returned by getRouterInterfaces.
 */
export interface GetRouterInterfacesResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of router interface IDs.
     */
    readonly ids: string[];
    /**
     * A list of router interfaces. Each element contains the following attributes:
     */
    readonly interfaces: outputs.vpc.GetRouterInterfacesInterface[];
    readonly nameRegex?: string;
    /**
     * A list of router interface names.
     */
    readonly names: string[];
    /**
     * Peer router interface ID.
     */
    readonly oppositeInterfaceId?: string;
    /**
     * Account ID of the owner of the peer router interface.
     */
    readonly oppositeInterfaceOwnerId?: string;
    readonly outputFile?: string;
    /**
     * Router interface role. Possible values: `InitiatingSide` and `AcceptingSide`.
     */
    readonly role?: string;
    /**
     * ID of the VRouter located in the local region.
     */
    readonly routerId?: string;
    /**
     * Router type in the local region. Possible values: `VRouter` and `VBR`.
     */
    readonly routerType?: string;
    /**
     * Router interface specification. Possible values: `Small.1`, `Middle.1`, `Large.2`, ...etc.
     */
    readonly specification?: string;
    /**
     * Router interface status. Possible values: `Active`, `Inactive` and `Idle`.
     */
    readonly status?: string;
}

export function getRouterInterfacesOutput(args?: GetRouterInterfacesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterInterfacesResult> {
    return pulumi.output(args).apply(a => getRouterInterfaces(a, opts))
}

/**
 * A collection of arguments for invoking getRouterInterfaces.
 */
export interface GetRouterInterfacesOutputArgs {
    /**
     * A list of router interface IDs.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A regex string used to filter by router interface name.
     */
    nameRegex?: pulumi.Input<string>;
    /**
     * ID of the peer router interface.
     */
    oppositeInterfaceId?: pulumi.Input<string>;
    /**
     * Account ID of the owner of the peer router interface.
     */
    oppositeInterfaceOwnerId?: pulumi.Input<string>;
    outputFile?: pulumi.Input<string>;
    /**
     * Role of the router interface. Valid values are `InitiatingSide` (connection initiator) and 
     * `AcceptingSide` (connection receiver). The value of this parameter must be `InitiatingSide` if the `routerType` is set to `VBR`.
     */
    role?: pulumi.Input<string>;
    /**
     * ID of the VRouter located in the local region.
     */
    routerId?: pulumi.Input<string>;
    /**
     * Router type in the local region. Valid values are `VRouter` and `VBR` (physical connection).
     */
    routerType?: pulumi.Input<string>;
    /**
     * Specification of the link, such as `Small.1` (10Mb), `Middle.1` (100Mb), `Large.2` (2Gb), ...etc.
     */
    specification?: pulumi.Input<string>;
    /**
     * Expected status. Valid values are `Active`, `Inactive` and `Idle`.
     */
    status?: pulumi.Input<string>;
}
