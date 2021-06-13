# mysensors-simulator
![lint workflow](https://github.com/jkandasa/mysensors-simulator/actions/workflows/lint.yaml/badge.svg)
![publish container images](https://github.com/jkandasa/mysensors-simulator/actions/workflows/publish_container_images.yaml/badge.svg)

MySensors simulator. Presentation only works<br>
Used to test MyController.org home automation controller

## Download
### Container images
`master` branch images are tagged as `:master`<br>
Both released and master branch container images are published into [Docker Hub](https://hub.docker.com/r/jkandasa/mysensors-simulator)

#### Docker Run
```bash
docker run --detach --name ms-simulator \
    --env  TZ="Asia/Kolkata" \
    --restart unless-stopped \
    docker.io/jkandasa/mysensors-simulator:master
```
