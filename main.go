package main

// Paket vreme
import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Prozvan/vreme/vreme"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Name: "vreme",

		// Privzeto
		Flags: []cli.Flag{
			&cli.Float64Flag{
				Name:  "minTemp",
				Value: -50,
			},
			&cli.Float64Flag{
				Name:  "maxTemp",
				Value: 50,
			},
			&cli.IntFlag{
				Name:  "stMeritev",
				Value: 3,
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			// Stikala
			// go run .\cmd\ --stMeritev=2
			minTemp := cmd.Float("minTemp")
			maxTemp := cmd.Float("maxTemp")
			stMeritev := cmd.Int("stMeritev")

			fmt.Println("Nastavljena najnižja temperatura:", minTemp)
			fmt.Println("Nastavljena najvišja temperatura:", maxTemp)
			fmt.Println("Nastavljeno število meritev:", stMeritev)
			fmt.Println()

			postaje := make(map[string]vreme.Postaja)
			vreme.DodajPostajo(postaje, "LJ01", "Kovinarska", "Ljubljana")

			vreme.DodajMeritev(postaje, "LJ01", 25.3, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 32.1, minTemp, maxTemp)
			vreme.DodajMeritev(postaje, "LJ01", 21.5, minTemp, maxTemp)

			// Izpis meritev in poročila
			vreme.IzpisMeritev(postaje)
			fmt.Println()
			vreme.IzpisPorocila(postaje, stMeritev)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
