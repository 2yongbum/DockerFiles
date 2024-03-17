# DockerFiles
Dcoker files for https://hub.docker.com/repositories/2yongbum

## cling jupyterlab docker

[https://gist.github.com/dsuess/059b86ea55d639bb99175c9a8cd2ca3e](https://gist.github.com/dsuess/059b86ea55d639bb99175c9a8cd2ca3e)

### Dockerfile
```
FROM frolvlad/alpine-miniconda3

RUN conda install -y -c conda-forge bash jupyterlab jupyter_contrib_nbextensions
RUN conda install -y -c conda-forge xeus-cling xtensor
RUN mkdir /notebooks
```

### build
docker build -t jupyter-pp .

### run
docker run -d --rm -p 8888:8888 --workdir /notebooks -it jupyter-pp jupyter-lab  --allow-root --ip=0.0.0.0

## Wine
* https://github.com/DIV2014A/wine/blob/master/README.md
* https://hub.docker.com/r/suchja/wine
* https://www.winehq.org
