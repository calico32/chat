package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"github.com/wiisportsresort/chat/server/ent"

	_ "github.com/mattn/go-sqlite3"
)

var mode string

var clients = make(map[string]*Client)
var admins = []string{}

func main() {
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		godotenv.Load("../.env")
	}

	if os.Getenv("HOSTNAME") == "" {
		log.Fatal("$HOSTNAME not set")
	}

	if os.Getenv("API_PORT") == "" {
		log.Fatal("$API_PORT not set")
	}

	if os.Getenv("ADMINS") == "" {
		log.Println("WARNING: $ADMINS not set")
	} else {
		admins = strings.Split(os.Getenv("ADMINS"), ",")
	}

	db, err := ent.Open("sqlite3", "file:ent.sqlite?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	ctx := context.Background()
	if err := db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()

	if mode == "release" {
		r.Use(gin.Recovery())
		r.SetTrustedProxies([]string{"172.0.0.1/16"})
	}

	var originCheck func(r *http.Request) bool

	if mode == "release" {
		originCheck = func(r *http.Request) bool {
			origin := r.Header.Get("Origin")
			return origin == "https://"+os.Getenv("HOSTNAME")
		}
	} else {
		originCheck = func(r *http.Request) bool { return true }
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     originCheck,
	}

	r.GET("/c/:channel", cors, NewGetMessages(ctx, db))
	r.GET("/p/:from", cors, NewGetPrivateMessageRecipients(ctx, db))
	r.GET("/p/:from/:to", cors, NewGetPrivateMessages(ctx, db))
	r.GET("/ws", cors, NewWebsocketHandler(ctx, db, &upgrader))
	log.Println("Listening on port " + os.Getenv("API_PORT"))
	log.Fatal(r.Run(":" + os.Getenv("API_PORT")))
}

func cors(c *gin.Context) {
	allowed := "https://" + os.Getenv("HOSTNAME") + ":" + os.Getenv("PORT")
	c.Writer.Header().Set("Access-Control-Allow-Origin", allowed)
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}
