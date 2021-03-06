package main

import (
	"fmt"
	"strings"
)

const MAX_PLAYERS = 10

type PUG struct {
	mapName string
	validMaps []string
	players[] string
	pugAdmin string

	pugStarted, pugActive bool
}

func (p *PUG) PugStarted() bool {
	return p.pugStarted
}

func (p *PUG) PugActive() bool {
	return p.pugActive
}

func (p *PUG) JoinPug(player string) bool {
	if (!p.pugStarted || len(p.players) >= MAX_PLAYERS) {
		return false
	}

	if (len(p.players) == 0) {
		p.pugAdmin = player;
	}

	players := strings.Join(p.players, " ")
	if (strings.Contains(players, player)) {
		return false
	} else {
		p.players = append(p.players, player)
		return true
	}
}

func (p *PUG) GetPlayerCount() int {
	return len(p.players)
}

func (p *PUG) LeavePug(player string) bool {
	if (!p.pugStarted) {
		return false
	}
	
	for i := range p.players {
		if (player == p.players[i]) {
			p.players = append(p.players[:i], p.players[i+1:]...)
			return true
		}
	}
	return false
}

func (p *PUG) SetMap(mapName string) {
	p.mapName = mapName
}

func (p *PUG) GetMap() string {
	return p.mapName
}

func (p *PUG) GetAdmin() string {
	return p.pugAdmin
}

func (p *PUG) GetPlayers() []string {
	return p.players
}

func (p *PUG) StartPug() {
	if (p.pugStarted) {
		return;
	}
	
	if (len(p.mapName) > 0) && p.IsValidMap(p.mapName) {
		fmt.Printf("Pug map is %s", p.mapName)
	} else {
		p.mapName = "de_dust2"
	}

	p.pugStarted = true
}

func (p *PUG) EndPug() {
	if (!p.pugStarted) {
		return
	}

	p.pugStarted = false
	p.pugActive = false
	p.mapName = ""
	p.players = nil
}

func (p *PUG) SetAllowedMaps(maps []string) {
	p.validMaps = maps;
}

func (p *PUG) IsValidMap(mapName string) bool {
	validMaps := strings.Join(p.validMaps, " ")
	if (strings.Contains(validMaps, mapName)) {
		return true
	} else {
		return false
	}
}
