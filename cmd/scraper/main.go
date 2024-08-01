package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cemayan/searchengine/constants"
	"github.com/cemayan/searchengine/internal/config"
	backendpb "github.com/cemayan/searchengine/protos/backendreq"
	pb "github.com/cemayan/searchengine/protos/searchreq"
	"github.com/cemayan/searchengine/trie"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math/rand/v2"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	configPath      string // get scraper config yaml path as flag
	configPathExtra string // get write api config yaml path as flag
	page            *rod.Page
	grpcCliConn     *grpc.ClientConn
)

func init() {
	flag.StringVar(&configPath, "config", "", "Path of config yaml")
	flag.StringVar(&configPathExtra, "configExtra", "", "Path of config yaml")
	flag.Parse()
	// config initializer
	config.Init(constants.Scraper, configPath)
	// db initializer
	config.Init(constants.WriteApi, configPathExtra)
}

type server struct {
	pb.UnimplementedSearcherServer
}

type SearchResponse struct {
	Url   string `json:"url"`
	Title string `json:"title"`
}

func getMs(min, max int) int {
	return rand.IntN(max-min) + min
}

func (s server) SearchHandler(ctx context.Context, request *pb.SearchRequest) (*pb.SearchRequest, error) {

	newBackendGrpcClient()
	defer grpcCliConn.Close()

	_trie := trie.New()

	indexing := _trie.ConvertForIndexing(request.Record)

	for k, _ := range indexing {
		page.MustElement("textarea[name=q]").MustInput(k).MustType(input.Enter)
		page.MustWaitStable()

		elements := page.MustElements("#rso span[jscontroller]:not([jscontroller='']) > a")

		arr := []*backendpb.BackendRequestItem{}

		for _, element := range elements {

			item := &backendpb.BackendRequestItem{}

			href, err := element.Attribute("href")
			if err == nil {
				item.Url = *href
			}

			title, err := element.Element("h3")
			if err == nil {
				item.Title = title.MustText()
			}

			arr = append(arr, item)
		}

		client := backendpb.NewDbServiceClient(grpcCliConn)

		_, err := client.SendRequest(ctx, &backendpb.BackendRequest{Items: arr, Key: k})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		time.Sleep(time.Duration(getMs(100, 250)) * time.Millisecond)
	}

	logrus.Infoln("scraping has been completed")

	return &pb.SearchRequest{}, nil
}

func searchGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GetConfig(constants.Scraper).Scraper.Server.Port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSearcherServer(s, &server{})
	logrus.Printf("grpc server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}

func newBackendGrpcClient() {
	writeApiConf := config.GetConfig(constants.WriteApi)

	addr := fmt.Sprintf("%s:%d", writeApiConf.Scraper.Server.Host, writeApiConf.Scraper.Server.Port)

	// Set up a connection to the server.
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	grpcCliConn = conn
}

func main() {

	go func() {
		searchGrpcServer()
	}()

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	path, _ := launcher.LookPath()
	u := launcher.New().Bin(path).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()

	defer browser.MustClose()

	logrus.Infoln("browser has been launched")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Create a new page
	page = browser.MustPage("https://google.com").MustWaitStable()

	logrus.Infoln("page has been created")

	<-done

	logrus.Infoln("browser has been closed")
	defer cancel()
}
