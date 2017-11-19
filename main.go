/////////////////////////////////////////////////
// Author: Steve Horan <shoran@theatsgroup.com //
// This REST application will stop and start   //
// Zookeeper and Kafka Docker instances.       //
//											   //
// Visit README.md for endpoints			   //
/////////////////////////////////////////////////

package main

////////////
// Import //
////////////

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

///////////
// Types //
///////////

// I created to seperate structs to keep this backward compatible with the
// Flask REST API

type KafkaResp struct {
	Kafka string `json:"Kafka"`
}

type ZkResp struct {
	Zookeeper string `json:"Zookeeper"`
}

///////////////
// Functions //
///////////////

func DockerClient() (*client.Client, []types.Container) {
	// This function will attempt to connect to the docker socket and return
	// A byte slice of containers along with the client.
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		panic(err)
	}

	return cli, containers
}

///////////////////
// mux Functions //
///////////////////

func brokerStatus(w http.ResponseWriter, r *http.Request) {
	_, containers := DockerClient()

	for _, container := range containers {
		if container.Names[0] == "/kafka" && container.State == "running" {
			r := KafkaResp{Kafka: "running"}
			json.NewEncoder(w).Encode(r)
		} else {
			r := KafkaResp{Kafka: "Not Running"}
			json.NewEncoder(w).Encode(r)
		}
	}
}

func brokerKill(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, containers := DockerClient()

	for _, container := range containers {
		if container.Names[0] == "/kafka" {
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				r := KafkaResp{Kafka: "Invalid Container"}
				json.NewEncoder(w).Encode(r)
			}
			r := KafkaResp{Kafka: "Killed"}
			json.NewEncoder(w).Encode(r)
		} else {
			r := KafkaResp{Kafka: "Not Running"}
			json.NewEncoder(w).Encode(r)
		}
	}
}

func brokerStart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, containers := DockerClient()

	for _, container := range containers {
		if container.Names[0] == "/kafka" && container.State == "running" {
			r := KafkaResp{Kafka: "Already Running"}
			json.NewEncoder(w).Encode(r)
		} else {
			if err := cli.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
				r := KafkaResp{Kafka: "Cannot Start"}
				json.NewEncoder(w).Encode(r)
			}
			r := KafkaResp{Kafka: "Started"}
			json.NewEncoder(w).Encode(r)
		}
	}
}

func zookeeperStatus(w http.ResponseWriter, r *http.Request) {
	_, containers := DockerClient()

	for _, container := range containers {
		if container.Names[0] == "/zookeeper" && container.State == "running" {
			r := ZkResp{Zookeeper: "running"}
			json.NewEncoder(w).Encode(r)
		} else {
			r := ZkResp{Zookeeper: "Not Running"}
			json.NewEncoder(w).Encode(r)
		}
	}
}

func zookeeperKill(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, containers := DockerClient()

	for _, container := range containers {
		if container.Names[0] == "/zookeeper" {
			if err := cli.ContainerStop(ctx, container.ID, nil); err != nil {
				r := ZkResp{Zookeeper: "Invalid Container"}
				json.NewEncoder(w).Encode(r)
			}
			r := ZkResp{Zookeeper: "Killed"}
			json.NewEncoder(w).Encode(r)
		} else {
			r := ZkResp{Zookeeper: "Not Running"}
			json.NewEncoder(w).Encode(r)
		}
	}
}

func zookeeperStart(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	cli, containers := DockerClient()

	for _, container := range containers {
		if container.Names[0] == "/zookeeper" && container.State == "running" {
			r := ZkResp{Zookeeper: "Already Running"}
			json.NewEncoder(w).Encode(r)
		} else {
			if err := cli.ContainerStart(ctx, container.ID, types.ContainerStartOptions{}); err != nil {
				r := ZkResp{Zookeeper: "Cannot Start"}
				json.NewEncoder(w).Encode(r)
			}
			r := ZkResp{Zookeeper: "Started"}
			json.NewEncoder(w).Encode(r)
		}
	}
}

//////////
// Main //
//////////

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/broker/status", brokerStatus).Methods("GET")
	router.HandleFunc("/broker/kill", brokerKill).Methods("GET")
	router.HandleFunc("/broker/start", brokerStart).Methods("GET")
	router.HandleFunc("/zookeeper/status", zookeeperStatus).Methods("GET")
	router.HandleFunc("/zookeeper/kill", zookeeperKill).Methods("GET")
	router.HandleFunc("/zookeeper/start", zookeeperStart).Methods("GET")
	log.Fatal(http.ListenAndServe(":4444", router))
}
