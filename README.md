# docker-utils
docker utilities: toolbox to help work efficiently with docker.
It provides three functionality.
- Remove untagged images locally from docker, to free up spaces.
- Deletes containers with "exited" status.
- Flatten the docker Image i.e. merges the layers of docker image, to save space.
    - NOTE: This utility, do deletes original image.

    Below is help from docker-utils
    ```
    $ ./docker-utils --help
    NAME:
       docker-utils - Toolchain for docker

    USAGE:
       docker-utils [global options] command [command options] [arguments...]

    VERSION:
        0.1.0

    COMMANDS:
        rmi          deletes the docker images
        rm           deletes docker containers
        flatten      Compacts the images by flattening
        help, h      Shows a list of commands or help for one command

        GLOBAL OPTIONS:
          --help, -h           show help
          --version, -v        print the version
    ```

# remove untagged images

To remove untagged images, "rmi" subcommand is provided.
- ``--untagged`` flag is mandatory flag to be passed with this subcommand.
- ``--dry`` flag, shows, the images that will be deleted, but no deletion happens.

```
$ ./docker-utils rmi --help
NAME:
   rmi - deletes the docker images

   USAGE:
      command rmi [command options] [arguments...]

   OPTIONS:
      --dry        [Optional] dry_run the command
      --untagged   [Required] deletes untagged images
```

example
  ``$ docker-utils rmi --untagged``

# remove exited containers

To remove exited or killed images, "rmi" subcommand is provided.
- ``--exited`` flag is mandatory flag to be passed with this subcommand.
- ``--dry`` flag, shows, the conatiners that will be deleted, but no deletion happens.

```
$ ./docker-utils rm --help
NAME:
   rm - deletes docker containers

   USAGE:
      command rm [command options] [arguments...]

   OPTIONS:
      --dry        [Optional] dry_run the command
      --exited     [Required] deletes exited containers
   ```
example
   ``$ docker-utils rm  --exited``


# Flatten the docker images
To flatten any image, ``flatten`` subcommand is used.
- Here ``--image`` is actual name of image in format <repo-name>:<tag> need to be provided.
- ``--name`` is new name which will be used while importing the image.
- ``--tag`` is new tag which will be used while importing the image.

```
$ ./docker-utils flatten --help
NAME:
   flatten - Compacts the images by flattening

   USAGE:
      command flatten [command options] [arguments...]

   OPTIONS:
      --image      [Required] Image file to flatten
      --name       [Required] New name of Image file
      --tag        [Required] tag for new image
```
example
   ``$docker-utils flatten --image centos:latest --name kk/centos --tag:7``

