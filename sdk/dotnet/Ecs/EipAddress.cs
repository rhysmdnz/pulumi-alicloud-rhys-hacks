// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    /// <summary>
    /// Provides a EIP Address resource.
    /// 
    /// For information about EIP Address and how to use it, see [What is EIP Address](https://www.alibabacloud.com/help/en/doc-detail/36016.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.126.0+.
    /// 
    /// &gt; **NOTE:** BGP (Multi-ISP) lines are supported in all regions. BGP (Multi-ISP) Pro lines are supported only in the China (Hong Kong) region.
    /// 
    /// &gt; **NOTE:** The resource only supports to create `PayAsYouGo PayByTraffic`  or `Subscription PayByBandwidth` elastic IP for international account. Otherwise, you will happened error `COMMODITY.INVALID_COMPONENT`.
    /// Your account is international if you can use it to login in [International Web Console](https://account.alibabacloud.com/login/login.htm).
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var example = new AliCloud.Ecs.EipAddress("example", new()
    ///     {
    ///         AddressName = "tf-testAcc1234",
    ///         InternetChargeType = "PayByBandwidth",
    ///         Isp = "BGP",
    ///         PaymentType = "PayAsYouGo",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// EIP Address can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:ecs/eipAddress:EipAddress example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ecs/eipAddress:EipAddress")]
    public partial class EipAddress : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The activity id.
        /// </summary>
        [Output("activityId")]
        public Output<string?> ActivityId { get; private set; } = null!;

        /// <summary>
        /// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        /// </summary>
        [Output("addressName")]
        public Output<string> AddressName { get; private set; } = null!;

        /// <summary>
        /// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `auto_pay` is `true`, The order will be automatically paid. When `auto_pay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `payment_type` is `Subscription`, this parameter is valid.
        /// </summary>
        [Output("autoPay")]
        public Output<bool?> AutoPay { get; private set; } = null!;

        /// <summary>
        /// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
        /// </summary>
        [Output("bandwidth")]
        public Output<string> Bandwidth { get; private set; } = null!;

        /// <summary>
        /// Whether enable the deletion protection or not. Default value: `false`.
        /// </summary>
        [Output("deletionProtection")]
        public Output<bool> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The description of the EIP.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Field `instance_charge_type` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `payment_type` instead.
        /// </summary>
        [Output("instanceChargeType")]
        public Output<string> InstanceChargeType { get; private set; } = null!;

        /// <summary>
        /// The metering method of the EIP. 
        /// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `payment_type` is "Subscription".
        /// </summary>
        [Output("internetChargeType")]
        public Output<string> InternetChargeType { get; private set; } = null!;

        /// <summary>
        /// The address of the EIP.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
        /// </summary>
        [Output("isp")]
        public Output<string> Isp { get; private set; } = null!;

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `address_name` instead.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of the network. Valid value is `public` (Internet).
        /// </summary>
        [Output("netmode")]
        public Output<string?> Netmode { get; private set; } = null!;

        /// <summary>
        /// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
        /// </summary>
        [Output("paymentType")]
        public Output<string> PaymentType { get; private set; } = null!;

        /// <summary>
        /// The duration that you will buy the resource, in month. It is valid when `payment_type` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The ID of the IP address pool. The EIP is allocated from the IP address pool. **NOTE:** The feature is available only to users whose accounts are included in the whitelist. If you want to use the feature,[submit a ticket](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
        /// </summary>
        [Output("publicIpAddressPoolId")]
        public Output<string?> PublicIpAddressPoolId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Output("resourceGroupId")]
        public Output<string> ResourceGroupId { get; private set; } = null!;

        /// <summary>
        /// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to `AntiDDoS_Enhanced`, Anti-DDoS Pro(Premium) is used.
        /// </summary>
        [Output("securityProtectionTypes")]
        public Output<ImmutableArray<string>> SecurityProtectionTypes { get; private set; } = null!;

        /// <summary>
        /// The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, object>?> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a EipAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EipAddress(string name, EipAddressArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/eipAddress:EipAddress", name, args ?? new EipAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EipAddress(string name, Input<string> id, EipAddressState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ecs/eipAddress:EipAddress", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/rhysmdnz/pulumi-alicloud",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EipAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EipAddress Get(string name, Input<string> id, EipAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new EipAddress(name, id, state, options);
        }
    }

    public sealed class EipAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activity id.
        /// </summary>
        [Input("activityId")]
        public Input<string>? ActivityId { get; set; }

        /// <summary>
        /// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        /// </summary>
        [Input("addressName")]
        public Input<string>? AddressName { get; set; }

        /// <summary>
        /// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `auto_pay` is `true`, The order will be automatically paid. When `auto_pay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `payment_type` is `Subscription`, this parameter is valid.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// Whether enable the deletion protection or not. Default value: `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The description of the EIP.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field `instance_charge_type` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `payment_type` instead.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The metering method of the EIP. 
        /// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `payment_type` is "Subscription".
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `address_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the network. Valid value is `public` (Internet).
        /// </summary>
        [Input("netmode")]
        public Input<string>? Netmode { get; set; }

        /// <summary>
        /// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource, in month. It is valid when `payment_type` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ID of the IP address pool. The EIP is allocated from the IP address pool. **NOTE:** The feature is available only to users whose accounts are included in the whitelist. If you want to use the feature,[submit a ticket](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
        /// </summary>
        [Input("publicIpAddressPoolId")]
        public Input<string>? PublicIpAddressPoolId { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityProtectionTypes")]
        private InputList<string>? _securityProtectionTypes;

        /// <summary>
        /// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to `AntiDDoS_Enhanced`, Anti-DDoS Pro(Premium) is used.
        /// </summary>
        public InputList<string> SecurityProtectionTypes
        {
            get => _securityProtectionTypes ?? (_securityProtectionTypes = new InputList<string>());
            set => _securityProtectionTypes = value;
        }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EipAddressArgs()
        {
        }
        public static new EipAddressArgs Empty => new EipAddressArgs();
    }

    public sealed class EipAddressState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The activity id.
        /// </summary>
        [Input("activityId")]
        public Input<string>? ActivityId { get; set; }

        /// <summary>
        /// The name of the EIP instance. This name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as "-",".","_", and must not begin or end with a hyphen, and must not begin with http:// or https://.
        /// </summary>
        [Input("addressName")]
        public Input<string>? AddressName { get; set; }

        /// <summary>
        /// Whether to pay automatically. Valid values: `true` and `false`. Default value: `true`. When `auto_pay` is `true`, The order will be automatically paid. When `auto_pay` is `false`, The order needs to go to the order center to complete the payment. **NOTE:** When `payment_type` is `Subscription`, this parameter is valid.
        /// </summary>
        [Input("autoPay")]
        public Input<bool>? AutoPay { get; set; }

        /// <summary>
        /// The maximum bandwidth of the EIP. Valid values: `1` to `200`. Unit: Mbit/s. Default value: `5`.
        /// </summary>
        [Input("bandwidth")]
        public Input<string>? Bandwidth { get; set; }

        /// <summary>
        /// Whether enable the deletion protection or not. Default value: `false`.
        /// </summary>
        [Input("deletionProtection")]
        public Input<bool>? DeletionProtection { get; set; }

        /// <summary>
        /// The description of the EIP.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Field `instance_charge_type` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `payment_type` instead.
        /// </summary>
        [Input("instanceChargeType")]
        public Input<string>? InstanceChargeType { get; set; }

        /// <summary>
        /// The metering method of the EIP. 
        /// Valid values: `PayByDominantTraffic`, `PayByBandwidth` and `PayByTraffic`. Default to `PayByBandwidth`. **NOTE:** It must be set to "PayByBandwidth" when `payment_type` is "Subscription".
        /// </summary>
        [Input("internetChargeType")]
        public Input<string>? InternetChargeType { get; set; }

        /// <summary>
        /// The address of the EIP.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The line type. You can set this parameter only when you create a `PayAsYouGo` EIP. Valid values: `BGP`: BGP (Multi-ISP) lines.Up to 89 high-quality BGP lines are available worldwide. Direct connections with multiple Internet Service Providers (ISPs), including Telecom, Unicom, Mobile, Railcom, Netcom, CERNET, China Broadcast Network, Dr. Peng, and Founder, can be established in all regions in mainland China. `BGP_PRO`:  BGP (Multi-ISP) Pro lines optimize data transmission to mainland China and improve connection quality for international services. Compared with BGP (Multi-ISP), when BGP (Multi-ISP) Pro provides services to clients in mainland China (excluding data centers), cross-border connections are established without using international ISP services. This reduces network latency.
        /// </summary>
        [Input("isp")]
        public Input<string>? Isp { get; set; }

        /// <summary>
        /// Field `name` has been deprecated from provider version 1.126.0, and it will be removed in the future version. Please use the new attribute `address_name` instead.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the network. Valid value is `public` (Internet).
        /// </summary>
        [Input("netmode")]
        public Input<string>? Netmode { get; set; }

        /// <summary>
        /// The billing method of the EIP. Valid values: `Subscription` and `PayAsYouGo`. Default value is `PayAsYouGo`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// The duration that you will buy the resource, in month. It is valid when `payment_type` is `Subscription`. Valid values: [1-9, 12, 24, 36]. At present, the provider does not support modify "period" and you can do that via web console.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The ID of the IP address pool. The EIP is allocated from the IP address pool. **NOTE:** The feature is available only to users whose accounts are included in the whitelist. If you want to use the feature,[submit a ticket](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/429100).
        /// </summary>
        [Input("publicIpAddressPoolId")]
        public Input<string>? PublicIpAddressPoolId { get; set; }

        /// <summary>
        /// The ID of the resource group.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        [Input("securityProtectionTypes")]
        private InputList<string>? _securityProtectionTypes;

        /// <summary>
        /// The edition of Anti-DDoS. If you do not set this parameter, Anti-DDoS Origin Basic is used. If you set the value to `AntiDDoS_Enhanced`, Anti-DDoS Pro(Premium) is used.
        /// </summary>
        public InputList<string> SecurityProtectionTypes
        {
            get => _securityProtectionTypes ?? (_securityProtectionTypes = new InputList<string>());
            set => _securityProtectionTypes = value;
        }

        /// <summary>
        /// The status of the EIP. Valid values:  `Associating`: The EIP is being associated. `Unassociating`: The EIP is being disassociated. `InUse`: The EIP is allocated. `Available`:The EIP is available.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// A mapping of tags to assign to the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public EipAddressState()
        {
        }
        public static new EipAddressState Empty => new EipAddressState();
    }
}