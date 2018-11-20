# coding: utf-8

from __future__ import absolute_import
from datetime import date, datetime  # noqa: F401

from typing import List, Dict  # noqa: F401

from tapi_server.models.base_model_ import Model
from tapi_server.models.node_edge_point import NodeEdgePoint  # noqa: F401,E501
from tapi_server import util


class GetNodeEdgePointDetailsRPCOutputSchema(Model):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    def __init__(self, node_edge_point: NodeEdgePoint=None):  # noqa: E501
        """GetNodeEdgePointDetailsRPCOutputSchema - a model defined in Swagger

        :param node_edge_point: The node_edge_point of this GetNodeEdgePointDetailsRPCOutputSchema.  # noqa: E501
        :type node_edge_point: NodeEdgePoint
        """
        self.swagger_types = {
            'node_edge_point': NodeEdgePoint
        }

        self.attribute_map = {
            'node_edge_point': 'node-edge-point'
        }

        self._node_edge_point = node_edge_point

    @classmethod
    def from_dict(cls, dikt) -> 'GetNodeEdgePointDetailsRPCOutputSchema':
        """Returns the dict as a model

        :param dikt: A dict.
        :type: dict
        :return: The get-node-edge-point-detailsRPC_output_schema of this GetNodeEdgePointDetailsRPCOutputSchema.  # noqa: E501
        :rtype: GetNodeEdgePointDetailsRPCOutputSchema
        """
        return util.deserialize_model(dikt, cls)

    @property
    def node_edge_point(self) -> NodeEdgePoint:
        """Gets the node_edge_point of this GetNodeEdgePointDetailsRPCOutputSchema.


        :return: The node_edge_point of this GetNodeEdgePointDetailsRPCOutputSchema.
        :rtype: NodeEdgePoint
        """
        return self._node_edge_point

    @node_edge_point.setter
    def node_edge_point(self, node_edge_point: NodeEdgePoint):
        """Sets the node_edge_point of this GetNodeEdgePointDetailsRPCOutputSchema.


        :param node_edge_point: The node_edge_point of this GetNodeEdgePointDetailsRPCOutputSchema.
        :type node_edge_point: NodeEdgePoint
        """

        self._node_edge_point = node_edge_point
