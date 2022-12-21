// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Alb.Outputs
{

    [OutputType]
    public sealed class GetListenersListenerQuicConfigResult
    {
        /// <summary>
        /// The ID of the QUIC listener to be associated. If QuicUpgradeEnabled is set to true, this parameter is required. Only HTTPS listeners support this parameter.
        /// </summary>
        public readonly string QuicListenerId;
        /// <summary>
        /// Indicates whether quic upgrade is enabled. Valid values: true and false. Default value: false.
        /// </summary>
        public readonly bool QuicUpgradeEnabled;

        [OutputConstructor]
        private GetListenersListenerQuicConfigResult(
            string quicListenerId,

            bool quicUpgradeEnabled)
        {
            QuicListenerId = quicListenerId;
            QuicUpgradeEnabled = quicUpgradeEnabled;
        }
    }
}