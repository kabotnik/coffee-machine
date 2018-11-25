# coffee-machine
Lightweight app simulating a coffee maker. Meant to test infrastructure. Has ports in Go, NodeJS, and soon in .NET Core.

[![Codefresh build status]( https://g.codefresh.io/api/badges/pipeline/kabotnik/kabotnik%2Fcoffee-machine%2Fcoffee-machine?branch=master&key=eyJhbGciOiJIUzI1NiJ9.NWJmNzNiYTE2MzEzZjBhYTRjMGM1ZGM5.fkm41RqK2FVEetIqJ2vts5tbOiHuuwYtm8BIpze6ZdA&type=cf-1)]( https://g.codefresh.io/pipelines/coffee-machine/builds?repoOwner=kabotnik&repoName=coffee-machine&serviceName=kabotnik%2Fcoffee-machine&filter=trigger:build~Build;branch:master;pipeline:5bf8c2906da223d75f6c35c3~coffee-machine)

## Building
Each port contains a `Dockerfile`. `cd` to the appropriate directory and run
```
docker build . -t coffee-machine
```

The resulting image can then be run with
```
docker run -it --rm -p 9090:9090 coffee-machine
```

This API can be reached at http://localhost:9090/coffee
