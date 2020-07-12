package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// reference: https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go

// birçok cli uygulaması, bir takım araçları tek komut altında toplamak için
// genellikle alt komutlardan (sub-commands) faydalanır
// her bir alt komutun kendi flag koleksiyonuna sahip olabildiğini bilmekte fayda var!

type GreetCommand struct {
	fs *flag.FlagSet

	name string
}

func NewGreetCommand() *GreetCommand {
	gc := &GreetCommand{
		fs: flag.NewFlagSet("greet", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "name", "World", "name of the person to be greeted")

	return gc
}

func (g *GreetCommand) Name() string {
	return g.fs.Name()
}

func (g *GreetCommand) Init(args []string) error {
	return g.fs.Parse(args)
}

func (g *GreetCommand) Run() error {
	fmt.Println("Hello", g.name, "!")
	return nil
}

type Runner interface {
	Init([]string) error
	Run() error
	Name() string
}

// root fonksiyonu, tüm subcommandlerin tanımlanacağı Runner interfaceini tanımlıyor
func root(args []string) error {
	if len(args) < 1 {
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		NewGreetCommand(),
	}
	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {
	// herhangi bir fonksiyon hata dönerse, if koşulumuz onu yakalıyor
	// programa verilen tüm argümanlar root'a gönderiliyor
	// os.Args[1:] sebebi programın kendi isminin de komut olarak geçmesi
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
