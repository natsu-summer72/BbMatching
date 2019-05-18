package main

import (
	"context"
	"database/sql"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"

	bbm "github.com/natsu-summer72/BbMatching/controller"
	"github.com/natsu-summer72/BbMatching/gen/user"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "localhost", "Server host (valid values: localhost)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD") + "@(" + os.Getenv("MYSQL_HOST") + ":"+ os.Getenv("MYSQL_PORT") +")" + "/" + os.Getenv("MYSQL_DATABASE")
	db, err := sql.Open("mysql", host)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS"))
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}


	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
		authClient *auth.Client
		Database *sql.DB
	)
	{
		logger = log.New(os.Stderr, "[bbm] ", log.Ltime)
		// Get an auth client from the firebase.App
		client, err := app.Auth(ctx)
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
		}
		authClient = client
		Database = db
	}



	// Databaseのテスト(MatchAPIを実装したら消す)
	type Recruit struct {
		id int64
		user_id string
		location string
		date string
		comment string
		disabled bool
		created_at string
	}


	rows, err := Database.Query("SELECT * from match_recruit")
	defer rows.Close()
	if err!=nil{
		panic(err.Error())
	}
	for rows.Next() {
		recruit:=Recruit{}
		err = rows.Scan(&recruit.id, &recruit.user_id, &recruit.location, &recruit.date, &recruit.comment, &recruit.disabled, &recruit.created_at)
		if err != nil {
			panic(err.Error())
		}
		println(recruit.location)
	}


	// Initialize the services.
	var (
		userSvc user.Service
	)
	{
		userSvc = bbm.NewUser(logger, authClient)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		userEndpoints *user.Endpoints
	)
	{
		userEndpoints = user.NewEndpoints(userSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		errc <- fmt.Errorf("%s", <-c)
	}()

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "localhost":
		{
			addr := "http://localhost:8080"
			u, err := url.Parse(addr)
			if err != nil {
				fmt.Fprintf(os.Stderr, "invalid URL %#v: %s", addr, err)
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, userEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: localhost)", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
