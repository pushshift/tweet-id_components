My first attempt at creating a Golang program to do something somewhat useful!

This will allow a bunch of tweet ids to be piped into it and it will extract the following components:

sequence_id
server_id
machine_id
datacenter_id
creation_time (milliseconds)

Example usage:

cat tweets.ndjson | jq '.id_str' | ./tweet -c all

Speed: Processes around one million ids per second (~60x faster than Python)

