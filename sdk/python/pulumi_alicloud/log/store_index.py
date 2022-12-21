# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['StoreIndexArgs', 'StoreIndex']

@pulumi.input_type
class StoreIndexArgs:
    def __init__(__self__, *,
                 logstore: pulumi.Input[str],
                 project: pulumi.Input[str],
                 field_searches: Optional[pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]]] = None,
                 full_text: Optional[pulumi.Input['StoreIndexFullTextArgs']] = None):
        """
        The set of arguments for constructing a StoreIndex resource.
        :param pulumi.Input[str] logstore: The log store name to the query index belongs.
        :param pulumi.Input[str] project: The project name to the log store belongs.
        :param pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]] field_searches: List configurations of field search index. Valid item as follows:
        :param pulumi.Input['StoreIndexFullTextArgs'] full_text: The configuration of full text index. Valid item as follows:
        """
        pulumi.set(__self__, "logstore", logstore)
        pulumi.set(__self__, "project", project)
        if field_searches is not None:
            pulumi.set(__self__, "field_searches", field_searches)
        if full_text is not None:
            pulumi.set(__self__, "full_text", full_text)

    @property
    @pulumi.getter
    def logstore(self) -> pulumi.Input[str]:
        """
        The log store name to the query index belongs.
        """
        return pulumi.get(self, "logstore")

    @logstore.setter
    def logstore(self, value: pulumi.Input[str]):
        pulumi.set(self, "logstore", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        The project name to the log store belongs.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="fieldSearches")
    def field_searches(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]]]:
        """
        List configurations of field search index. Valid item as follows:
        """
        return pulumi.get(self, "field_searches")

    @field_searches.setter
    def field_searches(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]]]):
        pulumi.set(self, "field_searches", value)

    @property
    @pulumi.getter(name="fullText")
    def full_text(self) -> Optional[pulumi.Input['StoreIndexFullTextArgs']]:
        """
        The configuration of full text index. Valid item as follows:
        """
        return pulumi.get(self, "full_text")

    @full_text.setter
    def full_text(self, value: Optional[pulumi.Input['StoreIndexFullTextArgs']]):
        pulumi.set(self, "full_text", value)


@pulumi.input_type
class _StoreIndexState:
    def __init__(__self__, *,
                 field_searches: Optional[pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]]] = None,
                 full_text: Optional[pulumi.Input['StoreIndexFullTextArgs']] = None,
                 logstore: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering StoreIndex resources.
        :param pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]] field_searches: List configurations of field search index. Valid item as follows:
        :param pulumi.Input['StoreIndexFullTextArgs'] full_text: The configuration of full text index. Valid item as follows:
        :param pulumi.Input[str] logstore: The log store name to the query index belongs.
        :param pulumi.Input[str] project: The project name to the log store belongs.
        """
        if field_searches is not None:
            pulumi.set(__self__, "field_searches", field_searches)
        if full_text is not None:
            pulumi.set(__self__, "full_text", full_text)
        if logstore is not None:
            pulumi.set(__self__, "logstore", logstore)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="fieldSearches")
    def field_searches(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]]]:
        """
        List configurations of field search index. Valid item as follows:
        """
        return pulumi.get(self, "field_searches")

    @field_searches.setter
    def field_searches(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['StoreIndexFieldSearchArgs']]]]):
        pulumi.set(self, "field_searches", value)

    @property
    @pulumi.getter(name="fullText")
    def full_text(self) -> Optional[pulumi.Input['StoreIndexFullTextArgs']]:
        """
        The configuration of full text index. Valid item as follows:
        """
        return pulumi.get(self, "full_text")

    @full_text.setter
    def full_text(self, value: Optional[pulumi.Input['StoreIndexFullTextArgs']]):
        pulumi.set(self, "full_text", value)

    @property
    @pulumi.getter
    def logstore(self) -> Optional[pulumi.Input[str]]:
        """
        The log store name to the query index belongs.
        """
        return pulumi.get(self, "logstore")

    @logstore.setter
    def logstore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logstore", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        """
        The project name to the log store belongs.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class StoreIndex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 field_searches: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StoreIndexFieldSearchArgs']]]]] = None,
                 full_text: Optional[pulumi.Input[pulumi.InputType['StoreIndexFullTextArgs']]] = None,
                 logstore: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Log store index can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:log/storeIndex:StoreIndex example tf-log:tf-log-store
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StoreIndexFieldSearchArgs']]]] field_searches: List configurations of field search index. Valid item as follows:
        :param pulumi.Input[pulumi.InputType['StoreIndexFullTextArgs']] full_text: The configuration of full text index. Valid item as follows:
        :param pulumi.Input[str] logstore: The log store name to the query index belongs.
        :param pulumi.Input[str] project: The project name to the log store belongs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: StoreIndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Log store index can be imported using the id, e.g.

        ```sh
         $ pulumi import alicloud:log/storeIndex:StoreIndex example tf-log:tf-log-store
        ```

        :param str resource_name: The name of the resource.
        :param StoreIndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(StoreIndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 field_searches: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StoreIndexFieldSearchArgs']]]]] = None,
                 full_text: Optional[pulumi.Input[pulumi.InputType['StoreIndexFullTextArgs']]] = None,
                 logstore: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = StoreIndexArgs.__new__(StoreIndexArgs)

            __props__.__dict__["field_searches"] = field_searches
            __props__.__dict__["full_text"] = full_text
            if logstore is None and not opts.urn:
                raise TypeError("Missing required property 'logstore'")
            __props__.__dict__["logstore"] = logstore
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
        super(StoreIndex, __self__).__init__(
            'alicloud:log/storeIndex:StoreIndex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            field_searches: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StoreIndexFieldSearchArgs']]]]] = None,
            full_text: Optional[pulumi.Input[pulumi.InputType['StoreIndexFullTextArgs']]] = None,
            logstore: Optional[pulumi.Input[str]] = None,
            project: Optional[pulumi.Input[str]] = None) -> 'StoreIndex':
        """
        Get an existing StoreIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['StoreIndexFieldSearchArgs']]]] field_searches: List configurations of field search index. Valid item as follows:
        :param pulumi.Input[pulumi.InputType['StoreIndexFullTextArgs']] full_text: The configuration of full text index. Valid item as follows:
        :param pulumi.Input[str] logstore: The log store name to the query index belongs.
        :param pulumi.Input[str] project: The project name to the log store belongs.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _StoreIndexState.__new__(_StoreIndexState)

        __props__.__dict__["field_searches"] = field_searches
        __props__.__dict__["full_text"] = full_text
        __props__.__dict__["logstore"] = logstore
        __props__.__dict__["project"] = project
        return StoreIndex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="fieldSearches")
    def field_searches(self) -> pulumi.Output[Optional[Sequence['outputs.StoreIndexFieldSearch']]]:
        """
        List configurations of field search index. Valid item as follows:
        """
        return pulumi.get(self, "field_searches")

    @property
    @pulumi.getter(name="fullText")
    def full_text(self) -> pulumi.Output[Optional['outputs.StoreIndexFullText']]:
        """
        The configuration of full text index. Valid item as follows:
        """
        return pulumi.get(self, "full_text")

    @property
    @pulumi.getter
    def logstore(self) -> pulumi.Output[str]:
        """
        The log store name to the query index belongs.
        """
        return pulumi.get(self, "logstore")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        The project name to the log store belongs.
        """
        return pulumi.get(self, "project")
