package main

import (
	"bufio"
	"fmt"
	"github.com/pborman/getopt"
	"os"
	"strconv"
	"strings"
)

type Tweet struct {
	id uint64
}

func (i *Tweet) sequence_id() string {
	return strconv.FormatUint(i.id & (1<<12 - 1),10)
}

func (i *Tweet) machine_id() string {
	return strconv.FormatUint((i.id >> 12) & (1<<5 - 1),10)
}

func (i *Tweet) server_id() string {
	return strconv.FormatUint((i.id >> 12) & (1<<10 - 1),10)
}

func (i *Tweet) datacenter_id() string {
	return strconv.FormatUint((i.id >> 17) & (1<<5 - 1),10)
}

func (i *Tweet) creation_time() string {
    return strconv.FormatUint((i.id >> 22) + 1288834974657,10)
}

func (i *Tweet) all() string {
    return strconv.FormatUint(i.id,10) + "," + i.sequence_id() + "," + i.machine_id() + "," + i.server_id() + "," + i.datacenter_id() + "," + i.creation_time()
}

func main() {

    // This program allows tweet ids to be piped into it and tweet id components to be extracted from those ids

	// Usage Example
	//cat twitter_ids | ./tweet -c server_id

	// Get Component Flag
	component := getopt.StringLong("component", 'c', "", `Tweet component. Can be one of: sequence_id, machine_id, server_id, datacenter_id, creation_time or all. All will
    print every component in the following order (id,sequence_id, machine_id, server_id, datacenter_id, creation_time)`)
	getopt.Parse()
	*component = strings.ToLower(*component)

	// Create alias to correct method based on flag
	var y func(*Tweet) string

    switch *component {
	case "creation_time":
		y = (*Tweet).creation_time
	case "server_id":
		y = (*Tweet).server_id
	case "machine_id":
		y = (*Tweet).machine_id
	case "datacenter_id":
		y = (*Tweet).datacenter_id
	case "sequence_id":
		y = (*Tweet).sequence_id
    case "all":
        y = (*Tweet).all
	default:
		fmt.Println("Unknown type: " + *component + "\n")
		getopt.Usage()
		return
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		id, err := strconv.ParseUint(input.Text(), 10, 64)
		if err == nil {
			tweet := Tweet{id}
			fmt.Println(y(&tweet))
		}
	}
}
