package main

import (
  

)
func main () {
    http. Handle ("/metrics", promhttp. Handler ())
    log. Fatal (http. Listenandserve (": 8080", nil))
}