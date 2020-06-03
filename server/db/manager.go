package db

var chCmd chan *cmd

type cmd struct {
	t int
}
