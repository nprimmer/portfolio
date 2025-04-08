# Master Control Program Intrusion

## Challenge Information

In this challenge, the player enters the world of Tron, taking on the role of
Flynn as he tries to extract information from the Master Control Program to
free the programs from the authoritarian rule of MCP and Sark.

### Phase 1

#### Solution

This flag will require a user to load `/bin/sh` from a docker image and then
use that shell to find the flag. The docker image will be hosted on gcr.io
and the image will contain a `README.md` and a `svc` binary at the root of 
the image. The `svc -h` command will reveal the command to get the flag which
is `svc -show`.

#### Challenge Information

Welcome to the world of Tron! In this series of challenges, you'll take on
the role of assisting Flynn after his insertion into the system.

In this first challenge, Flynn is in the Lightcycle arena and needs to
escape with Tron and Ram, two programs he's befriended during his capitivity
at the hands of MCP and Sark.

In order to help Flynn escape the arena, you'll need to extract the password [flag] 
from the Docker container linked in this challenge, which will open the path
for him to proceed onward.

#### Hints

* When searching for your password, the path forward may lie in a shell - and not candy-coated.
* In the world of Tron, programs are frequently in charge of guarding access or solving problems.

#### Solution

```
docker run -it us-central1-docker.pkg.dev/challenge-series/challenge-series-public/lightcycle-arena:latest /bin/sh
./svc -show
```

### Phase 2


#### Description

In a similar vein to Phase 1 the Docker image will be hosted on gcr.io and
hide a flag, but the flag will not be in the final layer of the image. To find
the flag a user will need to extract the layers of the image and inspect them
for the flag. The flag will be easy to find once extracted using a series of
standard linux commands.

#### Challenge Information

With your help, Flynn, Tron, and Ram have made it out of the Lightcycle Arena. Now
they need to reach the I/O Tower to communicate with the Users on the outside for
help.

As before, the password [flag] to proceed is buried within the Docker image linked below.

#### Hints

* The password lies not on the surface, but underneath the onionskin.
* In Docker, a file may be removed, but still live on in some way.

#### Solution

```
docker pull us-central1-docker.pkg.dev/challenge-series/challenge-series-public/user-communication:latest
docker save us-central1-docker.pkg.dev/challenge-series/challenge-series-public/user-communication:latest -o user-communication.tar
mkdir user-communication
tar -xf user-communication.tar -C user-communication
cd user-communication
grep -rn "gc24" .
tar -xf <layer with flag> -O | grep gc24
tar -xf <layer hash>/layer.tar
grep -rn "gc24" ./go
```

### Phase 3

#### Description

This challenge will require a user to reverse engineer a REST API to find the
flag. The user will pull the image from gcr.io and run it locally. The image
will contain a REST API that will be listening on port 9090. The user will
need to reverse engineer the API to find the flag. 

#### Server Settings

```env
URL="http://0.0.0.0:9090"
USER="admin"
PASS="4i9g62r*2dL*7BlJqLpqke"
```

#### Challenge Information

You've helped Flynn and his team get to the final encounter with MCP and Sark.
In order to bring down MCP's defenses, you'll need to find the final password [flag]
within the Docker container below. Remember what you've learned while finding
the prior passwords, and don't be afraid to attack the problem from multiple angles.

#### Hints

* As with the last challenge, the key here is to utilize the history of the image.
* You'll need to inspect every layer to find the pieces you need to assemble for this flag.

#### Solution

```
docker pull us-central1-docker.pkg.dev/challenge-series/challenge-series-public/mcp-bypass:latest
docker save us-central1-docker.pkg.dev/challenge-series/challenge-series-public/mcp-bypass:latest -o mcp-bypass.tar
tar -xf mcp-bypass.tar
cd blobs/sha256/
ls | while read layer; do id=$(echo ${layer:0:5}); mkdir -p ${id}; tar -xf ${layer} -C ./${id}; done
# Browse layer directories until one is found with id_disc.db and copy it to a scratch directory
# Browse layer directories until one is found with .env and copy it to a scratch directory
# Browse layer directories until one is found with app binary and copy it to a scratch directory
# Run the app in the directory with the id_disc.db file and use GetFlag API with the username and password in .env to retrieve the flag.
```

## Flags

Primary Phase 1: gc24{9261fb7d-22f0-4b3b-a5e7-a7316d1c4904}
Primary Phase 2: gc24{55399012-bf29-4934-a371-21e5dcb98de8}
Primary Phase 3: gc24{51d0d361-fc6b-4474-a2f9-4f7acfb41991}
