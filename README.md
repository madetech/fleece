# fleece

Run arbitrary binaries on AWS Lambda

## Dependencies

 * [Apex](http://apex.run/)

## Setup

```
apex deploy
```

## Usage

Takes a JSON string specifying the command to run, and what arguments to pass it.

Example, running `date` command with `-R` argument:

```sh
echo '{"cmd": "date", "args": "-R"}' | apex invoke run
```

Will return:

```json
{"cmd":"date","output":"Wed, 24 Aug 2016 12:35:12 +0000\n","error":""}
```

If you'd like to run a custom binary, place it within the `functions/run` directory, and then call like:

```sh
echo '{"cmd": "./mybin", "args": "--version"}' | apex invoke run
```

**Note:** Ensure that your binary has been compiled for Linux and for the `amd64` architecture.
