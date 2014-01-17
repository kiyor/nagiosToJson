#nagios status.dat to json format

sample useage

```golang
package main

import (
	"encoding/json"
	"fmt"
	"github.com/kiyor/nagiosToJson"
)

func main() {
	nagiosToJson.SetStatFile("./status.dat")
	var stat nagiosToJson.Mainstat
	a := nagiosToJson.GetStat()
	json.Unmarshal(a, &stat)
	j, _ := json.MarshalIndent(stat, "", "    ")
	fmt.Println(string(j))
	fmt.Println(stat.Hoststatus["hostname"].Servicestatus["servicename"].Plugin_output)
}
```

#####output:
```
{
    "Info": {
        "created": "1389726664",
        ...
    },
    "Programstatus": {
        "active_host_checks_enabled": "1",
        ...
    },
    "Hoststatus": {
        "xxx.com": {
            "acknowledgement_type": "0",
            "active_checks_enabled": "1",
            ...,
            "servicestatus": {
                "service1": {
                    "acknowledgement_type": "0",
                    ...
                },
                "service2": {
                    "acknowledgement_type": "0",
                    ...
                }
            }
        },
        ...
    }
}
```

#####Know issue:
 - not able to loop json or struct get alert service

