// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Sae.Inputs
{

    public sealed class LoadBalancerInternetInternetGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The SSL certificate. `https_cert_id` is required when HTTPS is selected
        /// </summary>
        [Input("httpsCertId")]
        public Input<string>? HttpsCertId { get; set; }

        /// <summary>
        /// The SLB Port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The Network protocol. Valid values: `TCP` ,`HTTP`,`HTTPS`.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The Container port.
        /// </summary>
        [Input("targetPort")]
        public Input<int>? TargetPort { get; set; }

        public LoadBalancerInternetInternetGetArgs()
        {
        }
        public static new LoadBalancerInternetInternetGetArgs Empty => new LoadBalancerInternetInternetGetArgs();
    }
}