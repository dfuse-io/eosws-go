`eosws` Go bindings (from the [dfuse API](https://dfuse.io/))
--------------------------------------

WebSocket consumer for the https://dfuse.io API on EOS networks.

## Examples

You will find a bunch of example scripts in the [examples](./examples) folder. Here the list
of provided examples:

 * [Get all actions for a specific account](./examples/get-actions/main.go)

### Instructions

All examples expect to have an API key being set in the `EOSWS_API_KEY` environment variable.
To ease fiddling with the examples, copy the file `.envrc.dist` (at root of project) to a
new file named `.envrc`. Edit the file to put your API key in the file.

Then install [direnv](https://direnv.net/) on your machine to auto-load the variable
from your terminal when you `cd` in the directory.

To run the examples, simply do the following:

```
go run ./examples/get-actions/main.go
```

**Note** Of course, change `get-actions` to the actual example you want to test out.
