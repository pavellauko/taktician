package main

import (
	"flag"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/nelhage/taktician/ai"
	"github.com/nelhage/taktician/playtak"
	"github.com/nelhage/taktician/tak"
)

var (
	server   = flag.String("server", "playtak.com:10000", "playtak.com server to connect to")
	user     = flag.String("user", "", "username for login")
	pass     = flag.String("pass", "", "password for login")
	accept   = flag.String("accept", "", "accept a game from specified user")
	gameTime = flag.Duration("time", 20*time.Minute, "Length of game to offer")
	size     = flag.Int("size", 5, "size of game to offer")
	once     = flag.Bool("once", false, "play a single game and exit")
	takbot   = flag.String("takbot", "", "challenge TakBot AI")

	debug = flag.Int("debug", 1, "debug level")
	depth = flag.Int("depth", 5, "minimax depth")
	limit = flag.Duration("limit", time.Minute, "time limit per move")
	sort  = flag.Bool("sort", true, "sort moves via history heuristic")
	table = flag.Bool("table", true, "use the transposition table")

	debugClient = flag.Bool("debug-client", false, "log debug output for playtak connection")
)

const ClientName = "Taktician AI"

func main() {
	flag.Parse()
	if *accept != "" || *takbot != "" {
		*once = true
	}

	backoff := 1 * time.Second
	for {
		client := &playtak.Client{
			Debug: *debugClient,
		}
		err := client.Connect(*server)
		if err != nil {
			goto reconnect
		}
		backoff = time.Second
		client.SendClient(ClientName)
		if *user != "" {
			err = client.Login(*user, *pass)
		} else {
			err = client.LoginGuest()
		}
		if err != nil {
			log.Fatal("login: ", err)
		}
		log.Printf("login OK")
		for {
			if *accept != "" {
				for line := range client.Recv {
					if strings.HasPrefix(line, "Seek new") {
						bits := strings.Split(line, " ")
						if bits[3] == *accept {
							log.Printf("accepting game %s from %s", bits[2], bits[3])
							client.SendCommand("Accept", bits[2])
							break
						}
					}
				}
			} else {
				client.SendCommand("Seek", strconv.Itoa(*size), strconv.Itoa(int(gameTime.Seconds())))
				log.Printf("Seek OK")
				if *takbot != "" {
					client.SendCommand("Shout", "takbot: play", *takbot)
				}
			}
			for line := range client.Recv {
				if strings.HasPrefix(line, "Game Start") {
					playGame(client, &Taktician{}, line)
					break
				}
			}
			if *once {
				return
			}
			if client.Error() != nil {
				log.Printf("Disconnected: %v", client.Error())
				break
			}
		}
	reconnect:
		log.Printf("sleeping %s before reconnect...", backoff)
		time.Sleep(backoff)
		backoff = backoff * 2
		if backoff > time.Minute {
			backoff = time.Minute
		}
	}
}

func timeBound(remaining time.Duration) time.Duration {
	return *limit
}

type Taktician struct {
	ai *ai.MinimaxAI
}

func (t *Taktician) NewGame(g *Game) {
	t.ai = ai.NewMinimax(ai.MinimaxConfig{
		Size:  g.size,
		Depth: *depth,
		Debug: *debug,

		NoSort:  !*sort,
		NoTable: !*table,
	})
}

func (t *Taktician) GetMove(p *tak.Position, mine, theirs time.Duration) tak.Move {
	return t.ai.GetMove(p, timeBound(mine))
}

func (t *Taktician) HandleChat(string, string) {}
