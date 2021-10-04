# coding=utf-8
# *** WARNING: this file was generated by test. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._enums import *

__all__ = [
    'example_func',
]

def example_func(enums: Optional[Sequence[Union[str, 'MyEnum']]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['enums'] = enums
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('my8110::exampleFunc', __args__, opts=opts).value

