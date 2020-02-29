<!-- shell file and commands -->
* there are 2 streams/channels in shell: stderr and stdout
* 
* bash shell file: <name>.sh and the first line is always: `#!/usr/bin/env bash`

* a shell file runs commands in sequence.
  
* `set -x` : it is a debug option, it not only executes all the commands but also prints them before it executes them (to stderr).
* `set -xe` : it is a debug option, it not only executes all the commands but also prints them before it executes them (to stderr). the flag -e terminates and exits when facing an error. if a command fails it does not go to the next command.
* add `time` at the beginning on every command to show the total time for that command.
* to redirect an stdout to a log file: <command which outputs stdout> > build.log
* to redirect an stdout to a black hole, hide it from screen : <command which outputs stdout> > /dev/null
* to direct an stdout to sterr: <command which ouputs stout> >$2   (2 is stderr and 1 is stdout)


<!-- Docker -->
* In the first iteration the workflow was to:
    1. run a docker (golang)image (and not building a docker) and mount a volume: `docker run -it -v "$PWD":/usr/src/app -w /usr/src/app golang`
    2. inside the container, install the go dependencies. for our case:`go get -u cgt.name/pkg/go-mwclient`
    3. build and run our source go file inside the container: `go run connectTest.go`

We wanted to automate the process; when the docker is running, during the runtime stage and not build stage, we want to be able to change the source code and see the changes without having to rebuild the docker everytime. We want to:
1. Create images:
   1. specify the specifications -> Dockerfile
   2. build the image -> `docker image build -t <our image name>`
2. (optional: distribute the image)
3. (Repeat) Run containers based on images `docker run ...`


If we use Dockerfile <a file for specifications for our custom image>, using `COPY` commands will detect the changes in the source code and therefore kicks in rebuilding the container. We want to run the code changs during the runtime `docker run...` and not during the build `Dockerfile`. So at the end, in the Dockerfile, we defined out custom image by adding `FROM` and `RUN` and in the `RUN` (when the docker is built), we added a command to install the go dependency packages. This way the packages are installed once when the docker is built and not everytime that the docker is running.

To automate the workflow, a bash shell script was created to automatically run some of the followig steps. The script run the docker run command which runs the container from golang image and then while inside the container, it takes another command as argument which is another bash file `runInDocker.sh`:
``` bash
#!/usr/bin/env bash
set -xe
docker image build -t snpedia:1.0 . > build.log # send the stdout to stderr
docker run --rm -v "$PWD":/usr/src/app -w /usr/src/app snpedia:1.0 ./runInDocker.sh

<inside runInDocker.sh>:
#!/usr/bin/env bash
time go run connectTest.go 

```
* to add execution privileges to the runInDocker.sh we need to run the following in the container only once:
* `chmod +x runInDocker.sh`
  
From the above, when the container is running, the second script builds and runs our source code.