// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Rds
{
    public static class GetRdsBackups
    {
        /// <summary>
        /// This data source provides the Rds Backups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.149.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = AliCloud.Rds.GetRdsBackups.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstRdsBackupId"] = example.Apply(getRdsBackupsResult =&gt; getRdsBackupsResult.Backups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetRdsBackupsResult> InvokeAsync(GetRdsBackupsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRdsBackupsResult>("alicloud:rds/getRdsBackups:getRdsBackups", args ?? new GetRdsBackupsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Rds Backups of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.149.0+.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var example = AliCloud.Rds.GetRdsBackups.Invoke(new()
        ///     {
        ///         DbInstanceId = "example_value",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["firstRdsBackupId"] = example.Apply(getRdsBackupsResult =&gt; getRdsBackupsResult.Backups[0]?.Id),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetRdsBackupsResult> Invoke(GetRdsBackupsInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetRdsBackupsResult>("alicloud:rds/getRdsBackups:getRdsBackups", args ?? new GetRdsBackupsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdsBackupsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// BackupMode.
        /// </summary>
        [Input("backupMode")]
        public string? BackupMode { get; set; }

        /// <summary>
        /// Backup task status. **NOTE:** This parameter will only be returned when a task is executed. Value:
        /// * **NoStart**: Not started
        /// * **Checking**: check the backup
        /// * **Preparing**: Prepare a backup
        /// * **Waiting**: Waiting for backup
        /// * **Uploading**: Upload backup
        /// * **Finished**: Complete backup
        /// * **Failed**: backup Failed
        /// </summary>
        [Input("backupStatus")]
        public string? BackupStatus { get; set; }

        /// <summary>
        /// The db instance id.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public string DbInstanceId { get; set; } = null!;

        /// <summary>
        /// The end time.
        /// </summary>
        [Input("endTime")]
        public string? EndTime { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Backup IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The start time.
        /// </summary>
        [Input("startTime")]
        public string? StartTime { get; set; }

        public GetRdsBackupsArgs()
        {
        }
        public static new GetRdsBackupsArgs Empty => new GetRdsBackupsArgs();
    }

    public sealed class GetRdsBackupsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// BackupMode.
        /// </summary>
        [Input("backupMode")]
        public Input<string>? BackupMode { get; set; }

        /// <summary>
        /// Backup task status. **NOTE:** This parameter will only be returned when a task is executed. Value:
        /// * **NoStart**: Not started
        /// * **Checking**: check the backup
        /// * **Preparing**: Prepare a backup
        /// * **Waiting**: Waiting for backup
        /// * **Uploading**: Upload backup
        /// * **Finished**: Complete backup
        /// * **Failed**: backup Failed
        /// </summary>
        [Input("backupStatus")]
        public Input<string>? BackupStatus { get; set; }

        /// <summary>
        /// The db instance id.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        /// <summary>
        /// The end time.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Backup IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The start time.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        public GetRdsBackupsInvokeArgs()
        {
        }
        public static new GetRdsBackupsInvokeArgs Empty => new GetRdsBackupsInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdsBackupsResult
    {
        public readonly string? BackupMode;
        public readonly string? BackupStatus;
        public readonly ImmutableArray<Outputs.GetRdsBackupsBackupResult> Backups;
        public readonly string DbInstanceId;
        public readonly string? EndTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? OutputFile;
        public readonly string? StartTime;

        [OutputConstructor]
        private GetRdsBackupsResult(
            string? backupMode,

            string? backupStatus,

            ImmutableArray<Outputs.GetRdsBackupsBackupResult> backups,

            string dbInstanceId,

            string? endTime,

            string id,

            ImmutableArray<string> ids,

            string? outputFile,

            string? startTime)
        {
            BackupMode = backupMode;
            BackupStatus = backupStatus;
            Backups = backups;
            DbInstanceId = dbInstanceId;
            EndTime = endTime;
            Id = id;
            Ids = ids;
            OutputFile = outputFile;
            StartTime = startTime;
        }
    }
}