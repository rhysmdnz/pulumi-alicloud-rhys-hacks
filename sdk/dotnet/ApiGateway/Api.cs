// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ApiGateway
{
    /// <summary>
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
    ///     var apiGroup = new AliCloud.ApiGateway.Group("apiGroup", new()
    ///     {
    ///         Description = "description of the api group",
    ///     });
    /// 
    ///     var apiGatewayApi = new AliCloud.ApiGateway.Api("apiGatewayApi", new()
    ///     {
    ///         GroupId = apiGroup.Id,
    ///         Description = "your description",
    ///         AuthType = "APP",
    ///         ForceNonceCheck = false,
    ///         RequestConfig = new AliCloud.ApiGateway.Inputs.ApiRequestConfigArgs
    ///         {
    ///             Protocol = "HTTP",
    ///             Method = "GET",
    ///             Path = "/test/path1",
    ///             Mode = "MAPPING",
    ///         },
    ///         ServiceType = "HTTP",
    ///         HttpServiceConfig = new AliCloud.ApiGateway.Inputs.ApiHttpServiceConfigArgs
    ///         {
    ///             Address = "http://apigateway-backend.alicloudapi.com:8080",
    ///             Method = "GET",
    ///             Path = "/web/cloudapi",
    ///             Timeout = 12,
    ///             AoneName = "cloudapi-openapi",
    ///         },
    ///         RequestParameters = new[]
    ///         {
    ///             new AliCloud.ApiGateway.Inputs.ApiRequestParameterArgs
    ///             {
    ///                 Name = "aaa",
    ///                 Type = "STRING",
    ///                 Required = "OPTIONAL",
    ///                 In = "QUERY",
    ///                 InService = "QUERY",
    ///                 NameService = "testparams",
    ///             },
    ///         },
    ///         StageNames = new[]
    ///         {
    ///             "RELEASE",
    ///             "TEST",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Api gateway api can be imported using the id.Format to `&lt;API Group Id&gt;:&lt;API Id&gt;` e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:apigateway/api:Api example "ab2351f2ce904edaa8d92a0510832b91:e4f728fca5a94148b023b99a3e5d0b62"
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:apigateway/api:Api")]
    public partial class Api : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the api of api gateway.
        /// </summary>
        [Output("apiId")]
        public Output<string> ApiId { get; private set; } = null!;

        /// <summary>
        /// The authorization Type including APP and ANONYMOUS. Defaults to null.
        /// </summary>
        [Output("authType")]
        public Output<string> AuthType { get; private set; } = null!;

        /// <summary>
        /// constant_parameters defines the constant parameters of the api.
        /// </summary>
        [Output("constantParameters")]
        public Output<ImmutableArray<Outputs.ApiConstantParameter>> ConstantParameters { get; private set; } = null!;

        /// <summary>
        /// The description of Constant parameter.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// fc_service_config defines the config when service_type selected 'FunctionCompute'.
        /// </summary>
        [Output("fcServiceConfig")]
        public Output<Outputs.ApiFcServiceConfig?> FcServiceConfig { get; private set; } = null!;

        /// <summary>
        /// Whether to prevent API replay attack. Default value: `false`.
        /// </summary>
        [Output("forceNonceCheck")]
        public Output<bool> ForceNonceCheck { get; private set; } = null!;

        /// <summary>
        /// The api gateway that the api belongs to. Defaults to null.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// http_service_config defines the config when service_type selected 'HTTP'.
        /// </summary>
        [Output("httpServiceConfig")]
        public Output<Outputs.ApiHttpServiceConfig?> HttpServiceConfig { get; private set; } = null!;

        /// <summary>
        /// http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.
        /// </summary>
        [Output("httpVpcServiceConfig")]
        public Output<Outputs.ApiHttpVpcServiceConfig?> HttpVpcServiceConfig { get; private set; } = null!;

        /// <summary>
        /// http_service_config defines the config when service_type selected 'MOCK'.
        /// </summary>
        [Output("mockServiceConfig")]
        public Output<Outputs.ApiMockServiceConfig?> MockServiceConfig { get; private set; } = null!;

        /// <summary>
        /// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Request_config defines how users can send requests to your API.
        /// </summary>
        [Output("requestConfig")]
        public Output<Outputs.ApiRequestConfig> RequestConfig { get; private set; } = null!;

        /// <summary>
        /// request_parameters defines the request parameters of the api.
        /// </summary>
        [Output("requestParameters")]
        public Output<ImmutableArray<Outputs.ApiRequestParameter>> RequestParameters { get; private set; } = null!;

        /// <summary>
        /// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
        /// </summary>
        [Output("serviceType")]
        public Output<string> ServiceType { get; private set; } = null!;

        /// <summary>
        /// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
        /// </summary>
        [Output("stageNames")]
        public Output<ImmutableArray<string>> StageNames { get; private set; } = null!;

        /// <summary>
        /// system_parameters defines the system parameters of the api.
        /// </summary>
        [Output("systemParameters")]
        public Output<ImmutableArray<Outputs.ApiSystemParameter>> SystemParameters { get; private set; } = null!;


        /// <summary>
        /// Create a Api resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Api(string name, ApiArgs args, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/api:Api", name, args ?? new ApiArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Api(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:apigateway/api:Api", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Api resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Api Get(string name, Input<string> id, ApiState? state = null, CustomResourceOptions? options = null)
        {
            return new Api(name, id, state, options);
        }
    }

    public sealed class ApiArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authorization Type including APP and ANONYMOUS. Defaults to null.
        /// </summary>
        [Input("authType", required: true)]
        public Input<string> AuthType { get; set; } = null!;

        [Input("constantParameters")]
        private InputList<Inputs.ApiConstantParameterArgs>? _constantParameters;

        /// <summary>
        /// constant_parameters defines the constant parameters of the api.
        /// </summary>
        public InputList<Inputs.ApiConstantParameterArgs> ConstantParameters
        {
            get => _constantParameters ?? (_constantParameters = new InputList<Inputs.ApiConstantParameterArgs>());
            set => _constantParameters = value;
        }

        /// <summary>
        /// The description of Constant parameter.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// fc_service_config defines the config when service_type selected 'FunctionCompute'.
        /// </summary>
        [Input("fcServiceConfig")]
        public Input<Inputs.ApiFcServiceConfigArgs>? FcServiceConfig { get; set; }

        /// <summary>
        /// Whether to prevent API replay attack. Default value: `false`.
        /// </summary>
        [Input("forceNonceCheck")]
        public Input<bool>? ForceNonceCheck { get; set; }

        /// <summary>
        /// The api gateway that the api belongs to. Defaults to null.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        /// <summary>
        /// http_service_config defines the config when service_type selected 'HTTP'.
        /// </summary>
        [Input("httpServiceConfig")]
        public Input<Inputs.ApiHttpServiceConfigArgs>? HttpServiceConfig { get; set; }

        /// <summary>
        /// http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.
        /// </summary>
        [Input("httpVpcServiceConfig")]
        public Input<Inputs.ApiHttpVpcServiceConfigArgs>? HttpVpcServiceConfig { get; set; }

        /// <summary>
        /// http_service_config defines the config when service_type selected 'MOCK'.
        /// </summary>
        [Input("mockServiceConfig")]
        public Input<Inputs.ApiMockServiceConfigArgs>? MockServiceConfig { get; set; }

        /// <summary>
        /// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Request_config defines how users can send requests to your API.
        /// </summary>
        [Input("requestConfig", required: true)]
        public Input<Inputs.ApiRequestConfigArgs> RequestConfig { get; set; } = null!;

        [Input("requestParameters")]
        private InputList<Inputs.ApiRequestParameterArgs>? _requestParameters;

        /// <summary>
        /// request_parameters defines the request parameters of the api.
        /// </summary>
        public InputList<Inputs.ApiRequestParameterArgs> RequestParameters
        {
            get => _requestParameters ?? (_requestParameters = new InputList<Inputs.ApiRequestParameterArgs>());
            set => _requestParameters = value;
        }

        /// <summary>
        /// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
        /// </summary>
        [Input("serviceType", required: true)]
        public Input<string> ServiceType { get; set; } = null!;

        [Input("stageNames")]
        private InputList<string>? _stageNames;

        /// <summary>
        /// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
        /// </summary>
        public InputList<string> StageNames
        {
            get => _stageNames ?? (_stageNames = new InputList<string>());
            set => _stageNames = value;
        }

        [Input("systemParameters")]
        private InputList<Inputs.ApiSystemParameterArgs>? _systemParameters;

        /// <summary>
        /// system_parameters defines the system parameters of the api.
        /// </summary>
        public InputList<Inputs.ApiSystemParameterArgs> SystemParameters
        {
            get => _systemParameters ?? (_systemParameters = new InputList<Inputs.ApiSystemParameterArgs>());
            set => _systemParameters = value;
        }

        public ApiArgs()
        {
        }
        public static new ApiArgs Empty => new ApiArgs();
    }

    public sealed class ApiState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the api of api gateway.
        /// </summary>
        [Input("apiId")]
        public Input<string>? ApiId { get; set; }

        /// <summary>
        /// The authorization Type including APP and ANONYMOUS. Defaults to null.
        /// </summary>
        [Input("authType")]
        public Input<string>? AuthType { get; set; }

        [Input("constantParameters")]
        private InputList<Inputs.ApiConstantParameterGetArgs>? _constantParameters;

        /// <summary>
        /// constant_parameters defines the constant parameters of the api.
        /// </summary>
        public InputList<Inputs.ApiConstantParameterGetArgs> ConstantParameters
        {
            get => _constantParameters ?? (_constantParameters = new InputList<Inputs.ApiConstantParameterGetArgs>());
            set => _constantParameters = value;
        }

        /// <summary>
        /// The description of Constant parameter.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// fc_service_config defines the config when service_type selected 'FunctionCompute'.
        /// </summary>
        [Input("fcServiceConfig")]
        public Input<Inputs.ApiFcServiceConfigGetArgs>? FcServiceConfig { get; set; }

        /// <summary>
        /// Whether to prevent API replay attack. Default value: `false`.
        /// </summary>
        [Input("forceNonceCheck")]
        public Input<bool>? ForceNonceCheck { get; set; }

        /// <summary>
        /// The api gateway that the api belongs to. Defaults to null.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// http_service_config defines the config when service_type selected 'HTTP'.
        /// </summary>
        [Input("httpServiceConfig")]
        public Input<Inputs.ApiHttpServiceConfigGetArgs>? HttpServiceConfig { get; set; }

        /// <summary>
        /// http_vpc_service_config defines the config when service_type selected 'HTTP-VPC'.
        /// </summary>
        [Input("httpVpcServiceConfig")]
        public Input<Inputs.ApiHttpVpcServiceConfigGetArgs>? HttpVpcServiceConfig { get; set; }

        /// <summary>
        /// http_service_config defines the config when service_type selected 'MOCK'.
        /// </summary>
        [Input("mockServiceConfig")]
        public Input<Inputs.ApiMockServiceConfigGetArgs>? MockServiceConfig { get; set; }

        /// <summary>
        /// System parameter name which supports values including in [system parameter list](https://www.alibabacloud.com/help/doc-detail/43677.html).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Request_config defines how users can send requests to your API.
        /// </summary>
        [Input("requestConfig")]
        public Input<Inputs.ApiRequestConfigGetArgs>? RequestConfig { get; set; }

        [Input("requestParameters")]
        private InputList<Inputs.ApiRequestParameterGetArgs>? _requestParameters;

        /// <summary>
        /// request_parameters defines the request parameters of the api.
        /// </summary>
        public InputList<Inputs.ApiRequestParameterGetArgs> RequestParameters
        {
            get => _requestParameters ?? (_requestParameters = new InputList<Inputs.ApiRequestParameterGetArgs>());
            set => _requestParameters = value;
        }

        /// <summary>
        /// The type of backend service. Type including HTTP,VPC and MOCK. Defaults to null.
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        [Input("stageNames")]
        private InputList<string>? _stageNames;

        /// <summary>
        /// Stages that the api need to be deployed. Valid value: `RELEASE`,`PRE`,`TEST`.
        /// </summary>
        public InputList<string> StageNames
        {
            get => _stageNames ?? (_stageNames = new InputList<string>());
            set => _stageNames = value;
        }

        [Input("systemParameters")]
        private InputList<Inputs.ApiSystemParameterGetArgs>? _systemParameters;

        /// <summary>
        /// system_parameters defines the system parameters of the api.
        /// </summary>
        public InputList<Inputs.ApiSystemParameterGetArgs> SystemParameters
        {
            get => _systemParameters ?? (_systemParameters = new InputList<Inputs.ApiSystemParameterGetArgs>());
            set => _systemParameters = value;
        }

        public ApiState()
        {
        }
        public static new ApiState Empty => new ApiState();
    }
}