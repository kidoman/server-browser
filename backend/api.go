package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-martini/martini"
	"github.com/golang/groupcache/lru"
)

var locationCache *lru.Cache

type Server struct {
	IP       string `json:"ip"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Ping     int    `json:"ping,omitempty"`
	Location string `json:"location,omitempty"`
}

var types = []string{"PUG", "Scrim", "DM"}

func randomName() string {
	return fmt.Sprintf("SoStronk %v Server %v", types[rand.Intn(3)], rand.Intn(20)+2)
}

var statuses = []string{"Available", "Waiting", "Busy", "Live"}

func randomStatus() string {
	return statuses[rand.Intn(4)]
}

func GetServerDetails(enc Encoder, params martini.Params) string {
	ip := params["ip"]

	l, ok := locationCache.Get(ip)
	if !ok {
		var err error
		l, err = location(ip)
		if err != nil {
			l = ""
		} else {
			locationCache.Add(ip, l)
		}
	}

	ping, err := ping(ip)
	if err != nil {
		ping = 999
	}
	s := Server{
		IP:       ip,
		Name:     randomName(),
		Status:   randomStatus(),
		Ping:     ping,
		Location: l.(string),
	}
	return Must(enc.Encode(s))
}

func GetServers(enc Encoder) string {
	r := rand.Intn(len(servers))
	return Must(enc.Encode(servers[:r]))
}

func init() {
	rand.Seed(time.Now().UnixNano())

	locationCache = lru.New(1000)
}
