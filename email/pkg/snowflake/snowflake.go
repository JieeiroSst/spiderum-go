package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"os"
)

type Snowflake struct {}

func NewSnowflake()*Snowflake {
	return &Snowflake{}
}

func (s *Snowflake) GearedID() int {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	return int(n.Generate().Int64())
}
