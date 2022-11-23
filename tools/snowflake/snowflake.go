package snowflake

import (
	"bluebell/config"
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func init() {
	var st time.Time
	cfg := config.SnowFlakeConfig
	st, err := time.Parse("2006-01-02", cfg.StartTime)
	if err != nil {
		panic(fmt.Sprintf("init snowflake failed, err:%v", err))
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(cfg.MachineID)
}

func GenIDInt64() int64 {
	return node.Generate().Int64()
}

func GenIDString() string {
	return node.Generate().String()
}
