# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from tapi_server.models.base_model_ import Model
from tapi_server import util


class GetConnectivityServiceDetailsRPCInputSchema(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, service_id_or_name: str=None):  # noqa: E501
        """GetConnectivityServiceDetailsRPCInputSchema - a model defined in Swagger

        :param service_id_or_name: The service_id_or_name of this GetConnectivityServiceDetailsRPCInputSchema.  # noqa: E501
        :type service_id_or_name: str
        """
        self.swagger_types = {
            'service_id_or_name': str
        }

        self.attribute_map = {
            'service_id_or_name': 'service-id-or-name'
        }

        self._service_id_or_name = service_id_or_name

    @classmethod
    def from_dict(cls, dikt) -> 'GetConnectivityServiceDetailsRPCInputSchema':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get-connectivity-service-detailsRPC_input_schema of this GetConnectivityServiceDetailsRPCInputSchema.  # noqa: E501
        :rtype: GetConnectivityServiceDetailsRPCInputSchema
        """
        return util.deserialize_model(dikt, cls)

    @property
    def service_id_or_name(self) -> str:
        """Gets the service_id_or_name of this GetConnectivityServiceDetailsRPCInputSchema.


        :return: The service_id_or_name of this GetConnectivityServiceDetailsRPCInputSchema.
        :rtype: str
        """
        return self._service_id_or_name

    @service_id_or_name.setter
    def service_id_or_name(self, service_id_or_name: str):
        """Sets the service_id_or_name of this GetConnectivityServiceDetailsRPCInputSchema.


        :param service_id_or_name: The service_id_or_name of this GetConnectivityServiceDetailsRPCInputSchema.
        :type service_id_or_name: str
        """

        self._service_id_or_name = service_id_or_name
