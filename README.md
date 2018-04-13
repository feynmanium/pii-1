# pii
Identifying Personally Identifiable Information

Honestly, this doesn't even really work. I thought it was kind of interesting, but what I've realised is I've essentially
reimplemented the Lintol server so.....

What this does is provides an endpoint, where a list of data can be passed, as well as the list of processors that you
wish to run against the data. Being able to pass in options with the processor gives additional opportunities for users
to fine tune the processing

```
# Sample Request
{
    "data": [
           "Field data",
           "Some names like Phil",
           "My IP address is 192.168.0.1",
           "Phils National insurance number is QQ123456C"
    ],
    "processors": [
        {
            "name": "ip-address",
        },
        {
            "name": "name",
            "options": {
                "additional-names": [
                    "Phil",
                    "Andrew"
                ]
            }
        },
        {
            "name": "national-insurance"
        }
    ]
}
```

This data would be passed via a GET request to the server, which would run each of the 3 selected processors against
each field in the data. The report would then come back with the prescribed JSON response, with each field being flagged
where the processor has identified a problem. This endpoint accepts a single dimension array rather than a two
dimensional array, as the processing of the original data file is best left to the content owner who will be more
familiar with the data. If they are not it is straightforward to write processors for such files, knowing that they can
all call the same endpoint

The server provides a single binary which can provide both a HTTP server which could be used on a server or as the
backend for an electron application. Using the Go plugin functionality means that anyone can write a Go (or C) shared
library as a processor, using the example in the documentations (see [GoDoc](https://godoc.org/github.com/tyndyll/pii/domain)),
place it in a directory and have it automatically picked up and made available. This could effectively create a
marketplace for shared libraries where processors could be made available by anyone to users. It also means that the
default processors that I have created (ipaddress, name and national-insurance) could be replaced by better, versions
