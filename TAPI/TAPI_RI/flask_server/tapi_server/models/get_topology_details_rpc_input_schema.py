# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from tapi_server.models.base_model_ import Model
from tapi_server import util


class GetTopologyDetailsRPCInputSchema(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, topology_id_or_name: str=None):  # noqa: E501
        """GetTopologyDetailsRPCInputSchema - a model defined in Swagger

        :param topology_id_or_name: The topology_id_or_name of this GetTopologyDetailsRPCInputSchema.  # noqa: E501
        :type topology_id_or_name: str
        """
        self.swagger_types = {
            'topology_id_or_name': str
        }

        self.attribute_map = {
            'topology_id_or_name': 'topology-id-or-name'
        }

        self._topology_id_or_name = topology_id_or_name

    @classmethod
    def from_dict(cls, dikt) -> 'GetTopologyDetailsRPCInputSchema':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get-topology-detailsRPC_input_schema of this GetTopologyDetailsRPCInputSchema.  # noqa: E501
        :rtype: GetTopologyDetailsRPCInputSchema
        """
        return util.deserialize_model(dikt, cls)

    @property
    def topology_id_or_name(self) -> str:
        """Gets the topology_id_or_name of this GetTopologyDetailsRPCInputSchema.


        :return: The topology_id_or_name of this GetTopologyDetailsRPCInputSchema.
        :rtype: str
        """
        return self._topology_id_or_name

    @topology_id_or_name.setter
    def topology_id_or_name(self, topology_id_or_name: str):
        """Sets the topology_id_or_name of this GetTopologyDetailsRPCInputSchema.


        :param topology_id_or_name: The topology_id_or_name of this GetTopologyDetailsRPCInputSchema.
        :type topology_id_or_name: str
        """

        self._topology_id_or_name = topology_id_or_name
