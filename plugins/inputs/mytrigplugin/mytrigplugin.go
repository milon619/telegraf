package mytrigplugin

// simple.go

import (
    "github.com/influxdata/telegraf"
    "github.com/influxdata/telegraf/plugins/inputs"
)

type MyTrigPlugin struct {
}

var mytrigconfig = `
  ## Set the amplitude
  amplitude = 10.0
`
func (s *MyTrigPlugin) Description() string {
    return "My first custom telegraf plugin. Inserts sine and cosine wave for demonstration purposes."
}

func (s *MyTrigPlugin) SampleConfig() string {
    return ""
}

func (s *MyTrigPlugin) Gather(acc telegraf.Accumulator) error {
      return nil
}

func init() {
    inputs.Add("mytrigplugin", func() telegraf.Input { return &MyTrigPlugin{} })
}
