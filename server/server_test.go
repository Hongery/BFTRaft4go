package server

import (
	"testing"
	"time"
	"os"
)

func initDB(dbPath string, t *testing.T) {
	if err := os.MkdirAll(dbPath, os.ModePerm); err != nil {
		t.Fatal(err)
	}
	InitDatabase(dbPath)
}

func getServer(dbPath string, addr string, bootstraps []string, t *testing.T) *BFTRaftServer {
	initDB(dbPath, t)
	s, err := GetServer(Options{
		DBPath: dbPath,
		Address: addr,
		Bootstrap: bootstraps,
		ConsensusTimeout: 1 * time.Minute,
	})

	if err != nil {
		t.Fatal(err)
	}
	return s
}

func TestServerStartup(t *testing.T) {
	dbPath := "test_data/ServerStartup"
	s := getServer(dbPath, "localhost:4560", []string{}, t)
	defer os.RemoveAll(dbPath)
	go func() {
		if err := s.StartServer(); err != nil {
			t.Fatal(err)
		}
	}()
}

func TestColdStart(t *testing.T) {
	// test for creating a cold started node and add a member to join it
	dbPath1 := "test_data/TestColdStart1"
	dbPath2 := "test_data/TestColdStart2"
	addr1 := "localhost:4561"
	addr2 := "localhost:4562"
	os.RemoveAll(dbPath1)
	os.RemoveAll(dbPath2)
	defer os.RemoveAll(dbPath1)
	defer os.RemoveAll(dbPath2)
	println("start server 1")
	s1 := getServer(dbPath1, addr1, []string{}, t)
	s1.StartServer()
	println("start server 2")
	time.Sleep(1 * time.Second)
	s2 := getServer(dbPath2, addr2, []string{addr1}, t)
	s2.StartServer()
	time.Sleep(1 * time.Second)
	s2.NodeJoin(1)
}