# Drone Enterprise

[![Build Status](https://github.drone.webzyno.com/api/badges/dockcenter/drone/status.svg)](https://github.drone.webzyno.com/dockcenter/drone)
![Docker Image Version (latest semver)](https://img.shields.io/docker/v/dockcenter/drone)
![Docker Pulls](https://img.shields.io/docker/pulls/dockcenter/drone?color=important)
![GitHub](https://img.shields.io/github/license/dockcenter/drone)

This project builds and pushes [dockcenter/drone](hub.docker.com/r/dockcenter/drone), a containerized drone enterprise version without build limitations.

## Why?

The [drone/drone]( https://hub.docker.com/r/drone/drone) container image you pull from Docker Hub is the Enterprise Edition and without a license is free for a trial period of 5000 builds after which you will need to obtain a license via their [website](https://drone.io/enterprise). 

As per <https://docs.drone.io/enterprise/#what-is-the-difference-between-open-source-and-enterprise>, the Drone Enterprise Edition is free for organizations with under $1 million US dollars in annual gross revenue.

You can build the Enterprise Edition as well as the severely limited Open Source Edition -- The two can be compared [here](https://docs.drone.io/enterprise/#what-is-the-difference-between-open-source-and-enterprise). -- from source using the build tags described [here](https://docs.drone.io/enterprise/#how-do-i-use-the-enterprise-edition-for-free). 

This project builds and containerized the Enterprise Edition with nolimits so you can use it for free, if you or your organization falls within the requirements of the [license](https://github.com/drone/drone/blob/master/LICENSE).

## License

Be careful using this container image as you must meet the obligations and conditions of the [license](https://github.com/drone/drone/blob/master/LICENSE) as not doing so will be subject you or your organization to penalty under US Federal and International copyright law.

A copy of the Drone Enterprise Edition license can be found [here](https://github.com/drone/drone/blob/master/LICENSE).

The code for [the project](https://github.com/nemonik/drone) that builds the [dockcenter/drone](https://hub.docker.com/r/dockcenter/drone) image and pushes it to Docker Hub is distributed under the [MIT License](https://github.com/dockcenter/drone/blob/main/LICENSE).

Please, don't confuse the two licenses.

## How to use this image

The server container can be started with the below command. The container is configured through environment variables. For a full list of configuration parameters, please see the [configuration reference](https://docs.drone.io/server/reference/).

```bash
docker run \
  --volume=/var/lib/drone:/data \
  --env=DRONE_GITHUB_CLIENT_ID=your-id \
  --env=DRONE_GITHUB_CLIENT_SECRET=super-duper-secret \
  --env=DRONE_RPC_SECRET=super-duper-secret \
  --env=DRONE_SERVER_HOST=drone.company.com \
  --env=DRONE_SERVER_PROTO=https \
  --publish=80:80 \
  --publish=443:443 \
  --restart=always \
  --detach=true \
  --name=drone \
  dockcenter/drone:2
```

For more installation information, please refer to Drone official [documentation](https://docs.drone.io/server/overview/).