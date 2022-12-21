// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CS
{
    /// <summary>
    /// This resource will help you to manage addon in Kubernetes Cluster.
    /// 
    /// &gt; **NOTE:** Available in 1.150.0+.
    /// **NOTE:** From version 1.166.0, support specifying addon customizable configuration.
    /// 
    /// ## Example Usage
    /// 
    /// **Create a managed cluster**
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tf-test";
    ///     var defaultZones = AliCloud.GetZones.Invoke(new()
    ///     {
    ///         AvailableResourceCreation = "VSwitch",
    ///     });
    /// 
    ///     var defaultInstanceTypes = AliCloud.Ecs.GetInstanceTypes.Invoke(new()
    ///     {
    ///         AvailabilityZone = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///         CpuCoreCount = 2,
    ///         MemorySize = 4,
    ///         KubernetesNodeRole = "Worker",
    ///     });
    /// 
    ///     var defaultNetwork = new AliCloud.Vpc.Network("defaultNetwork", new()
    ///     {
    ///         VpcName = name,
    ///         CidrBlock = "10.1.0.0/21",
    ///     });
    /// 
    ///     var defaultSwitch = new AliCloud.Vpc.Switch("defaultSwitch", new()
    ///     {
    ///         VswitchName = name,
    ///         VpcId = defaultNetwork.Id,
    ///         CidrBlock = "10.1.1.0/24",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Zones[0]?.Id),
    ///     });
    /// 
    ///     var defaultKeyPair = new AliCloud.Ecs.KeyPair("defaultKeyPair", new()
    ///     {
    ///         KeyPairName = name,
    ///     });
    /// 
    ///     var defaultManagedKubernetes = new List&lt;AliCloud.CS.ManagedKubernetes&gt;();
    ///     for (var rangeIndex = 0; rangeIndex &lt; (1 == true); rangeIndex++)
    ///     {
    ///         var range = new { Value = rangeIndex };
    ///         defaultManagedKubernetes.Add(new AliCloud.CS.ManagedKubernetes($"defaultManagedKubernetes-{range.Value}", new()
    ///         {
    ///             ClusterSpec = "ack.pro.small",
    ///             IsEnterpriseSecurityGroup = true,
    ///             WorkerNumber = 2,
    ///             Password = "Hello1234",
    ///             PodCidr = "172.20.0.0/16",
    ///             ServiceCidr = "172.21.0.0/20",
    ///             WorkerVswitchIds = new[]
    ///             {
    ///                 defaultSwitch.Id,
    ///             },
    ///             WorkerInstanceTypes = new[]
    ///             {
    ///                 defaultInstanceTypes.Apply(getInstanceTypesResult =&gt; getInstanceTypesResult.InstanceTypes[0]?.Id),
    ///             },
    ///         }));
    ///     }
    /// });
    /// ```
    /// **Installing of addon**
    /// When a cluster is created, some system addons and those specified at the time of cluster creation will be installed, so when an addon resource is applied:
    /// * If the addon already exists in the cluster and its version is the same as the specified version, it will be skipped and will not be reinstalled.
    /// * If the addon already exists in the cluster and its version is different from the specified version, the addon will be upgraded.
    /// * If the addon does not exist in the cluster, it will be installed.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ack_node_problem_detector = new AliCloud.CS.KubernetesAddon("ack-node-problem-detector", new()
    ///     {
    ///         ClusterId = alicloud_cs_managed_kubernetes.Default[0].Id,
    ///         Version = "1.2.7",
    ///     });
    /// 
    ///     var nginxIngressController = new AliCloud.CS.KubernetesAddon("nginxIngressController", new()
    ///     {
    ///         ClusterId = @var.Cluster_id,
    ///         Version = "v1.1.2-aliyun.2",
    ///         Config = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["CpuLimit"] = "",
    ///             ["CpuRequest"] = "100m",
    ///             ["EnableWebhook"] = true,
    ///             ["HostNetwork"] = false,
    ///             ["IngressSlbNetworkType"] = "internet",
    ///             ["IngressSlbSpec"] = "slb.s2.small",
    ///             ["MemoryLimit"] = "",
    ///             ["MemoryRequest"] = "200Mi",
    ///             ["NodeSelector"] = new[]
    ///             {
    ///             },
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// **Upgrading of addon**
    /// First, check the `next_version` field of the addon that can be upgraded to through the `.tfstate file`, then overwrite the `version` field with the value of `next_version` and apply.
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ack_node_problem_detector = new AliCloud.CS.KubernetesAddon("ack-node-problem-detector", new()
    ///     {
    ///         ClusterId = alicloud_cs_managed_kubernetes.Default[0].Id,
    ///         Version = "1.2.8",
    ///     });
    /// 
    ///     // upgrade from 1.2.7 to 1.2.8
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Cluster addon can be imported by cluster id and addon name. Then write the addon.tf file according to the result of `terraform plan`.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:cs/kubernetesAddon:KubernetesAddon my_addon &lt;cluster_id&gt;:&lt;addon_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cs/kubernetesAddon:KubernetesAddon")]
    public partial class KubernetesAddon : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Is the addon ready for upgrade.
        /// </summary>
        [Output("canUpgrade")]
        public Output<bool> CanUpgrade { get; private set; } = null!;

        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Output("clusterId")]
        public Output<string> ClusterId { get; private set; } = null!;

        /// <summary>
        /// The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
        /// </summary>
        [Output("config")]
        public Output<string?> Config { get; private set; } = null!;

        /// <summary>
        /// The name of addon.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The version which addon can be upgraded to.
        /// </summary>
        [Output("nextVersion")]
        public Output<string> NextVersion { get; private set; } = null!;

        /// <summary>
        /// Is it a mandatory addon to be installed.
        /// </summary>
        [Output("required")]
        public Output<bool> Required { get; private set; } = null!;

        /// <summary>
        /// The current version of addon.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a KubernetesAddon resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public KubernetesAddon(string name, KubernetesAddonArgs args, CustomResourceOptions? options = null)
            : base("alicloud:cs/kubernetesAddon:KubernetesAddon", name, args ?? new KubernetesAddonArgs(), MakeResourceOptions(options, ""))
        {
        }

        private KubernetesAddon(string name, Input<string> id, KubernetesAddonState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cs/kubernetesAddon:KubernetesAddon", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing KubernetesAddon resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static KubernetesAddon Get(string name, Input<string> id, KubernetesAddonState? state = null, CustomResourceOptions? options = null)
        {
            return new KubernetesAddon(name, id, state, options);
        }
    }

    public sealed class KubernetesAddonArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The name of addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The current version of addon.
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public KubernetesAddonArgs()
        {
        }
        public static new KubernetesAddonArgs Empty => new KubernetesAddonArgs();
    }

    public sealed class KubernetesAddonState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Is the addon ready for upgrade.
        /// </summary>
        [Input("canUpgrade")]
        public Input<bool>? CanUpgrade { get; set; }

        /// <summary>
        /// The id of kubernetes cluster.
        /// </summary>
        [Input("clusterId")]
        public Input<string>? ClusterId { get; set; }

        /// <summary>
        /// The custom configuration of addon. You can checkout the customizable configuration of the addon through datasource `alicloud.cs.getKubernetesAddonMetadata`, the returned format is the standard json schema. If return empty, it means that the addon does not support custom configuration yet. You can also checkout the current custom configuration through the data source `alicloud.cs.getKubernetesAddons`.
        /// </summary>
        [Input("config")]
        public Input<string>? Config { get; set; }

        /// <summary>
        /// The name of addon.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The version which addon can be upgraded to.
        /// </summary>
        [Input("nextVersion")]
        public Input<string>? NextVersion { get; set; }

        /// <summary>
        /// Is it a mandatory addon to be installed.
        /// </summary>
        [Input("required")]
        public Input<bool>? Required { get; set; }

        /// <summary>
        /// The current version of addon.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubernetesAddonState()
        {
        }
        public static new KubernetesAddonState Empty => new KubernetesAddonState();
    }
}