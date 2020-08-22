// Example of using the graphiteNop feature in action:
//
//     package mylib
//
//     import (
//         "github.com/msaf1980/graphite-golang"
//         "log"
//     )
//
//     func init() {
//
//         // load your configuration file / mechanism
//         config := newConfig()
//
//         // try to connect a graphite server
//         if config.GraphiteEnabled {
//             Graphite, err = graphite.NewGraphite(config.Graphite.Address)
//         } else {
//             Graphite = graphite.NewGraphiteNop(config.Graphite.Address)
//         }
//         // if you couldn't connect to graphite, use a nop
//         if err != nil {
//             Graphite = graphite.NewGraphiteNop(config.Graphite.Adress)
//         }
//
//         log.Printf("Loaded Graphite connection: %#v", Graphite)
//         Graphite.SimpleSend("stats.graphite_loaded", 1)
//     }
//
//     func doWork() {
//         // this will work just fine, regardless of if you're working with a graphite
//         // nop or not
//         Graphite.SimpleSend("stats.doing_work", 1)
//     }
package graphite
