/*
   Copyright 2014 Franc[e]sco (lolisamurai@tfwno.gf)
   This file is part of kagami.
   kagami is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.
   kagami is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
   GNU General Public License for more details.
   You should have received a copy of the GNU General Public License
   along with kagami. If not, see <http://www.gnu.org/licenses/>.
*/

package worlds

import (
	"github.com/Francesco149/kagami/common"
	"github.com/Francesco149/kagami/common/config"
)

// A World holds information about a single world, such as the player load, listening port and so on
type World struct {
	connected  bool
	id         byte
	port       uint32
	playerLoad int32
	worldcon   *common.EncryptedConnection
	config     *config.WorldConf
	channels   map[byte]*Channel
}

func (w *World) Connected() bool                       { return w.connected }
func (w *World) Id() byte                              { return w.id }
func (w *World) Port() uint32                          { return w.port }
func (w *World) PlayerLoad() int32                     { return w.playerLoad }
func (w *World) WorldCon() *common.EncryptedConnection { return w.worldcon }
func (w *World) Conf() *config.WorldConf               { return w.config }
func (w *World) Channel(id byte) *Channel              { return w.channels[id] }

func (w *World) SetConnected(connected bool)                      { w.connected = connected }
func (w *World) SetId(id byte)                                    { w.id = id }
func (w *World) SetPort(port uint32)                              { w.port = port }
func (w *World) SetPlayerLoad(playerLoad int32)                   { w.playerLoad = playerLoad }
func (w *World) SetWorldCon(worldcon *common.EncryptedConnection) { w.worldcon = worldcon }
func (w *World) SetConf(config *config.WorldConf)                 { w.config = config }

// ClearChannels deletes all of the channels in this world
func (w *World) ClearChannels() { w.channels = make(map[byte]*Channel) }

// RemoveChannel removes a channel by id
func (w *World) RemoveChannel(id byte) { delete(w.channels, id) }

// AddChannel adds a channel with the given id (overwrites if the given id already exists)
func (w *World) AddChannel(id byte, ch *Channel) { w.channels[id] = ch }