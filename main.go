package goinfras_tool

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"time"
)

func main()  {
	app := cli.NewApp()
	app.Name = "gt"
	app.Version = "1.0.0"
	app.Compiled = time.Now()
	app.Usage = "A generation tool of go app scaffold which base on bb-orz/goinfras."
	app.UsageText = "gt [option] [command] [args]"
	app.ArgsUsage = "[args and such]"
	app.UseShortOptionHandling = true

	app.Action = func(c *cli.Context) error {
		fmt.Println("gt (goinfras tool) is a generation tool of go app scaffold which base on bb-orz/goinfras.")
		return nil
	}
}