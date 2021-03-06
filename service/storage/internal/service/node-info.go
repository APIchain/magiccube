/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/

package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/bottos-project/magiccube/service/storage/proto"
)

// GetNodeInfos from db
func (c *StorageService) GetNodeInfos(ctx context.Context, request *storage.AllRequest, response *storage.NodeInfosResponse) error {

	if request == nil {
		response.Code = 0
		return errors.New("Missing storage request")
	}
	fmt.Println("GetNodeInfos")
	nodes, err := c.mgoRepo.CallGetNodeInfos()

	if nodes == nil || err != nil {
		response.Code = 0
		fmt.Println(err)
		return errors.New("Failed CallGetNodeInfos")

	}
	response.NodeList = []*storage.Node{}
	for _, node := range nodes {
		dbTag := &storage.Node{
			NodeId:   node.NodeId,
			NodeIp:   node.NodeIP,
			NodePort: node.NodePort}
		response.NodeList = append(response.NodeList, dbTag)
	}
	response.Code = 1
	return nil
}
