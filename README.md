# Calculate Flight

This is a simple program to calculate Start and Destination of a journey given 
the set of flights.

The server is listening on port 8080 and the end point exposed in /calculate.
It is a POST API call with flight list (JSON of array of arrays) as the payload
and response is the caluated journey.

For example, of the server runs locally; below is the CURL command to validate:

curl -X POST -H "Content-Type: application/json" -d '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]' http://localhost:8080/calculate

The response is:

["SFO","EWR"]

