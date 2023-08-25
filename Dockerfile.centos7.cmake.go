FROM centos:7

RUN yum install -y gcc gcc-c++ make wget perl libgcc

RUN wget https://www.openssl.org/source/openssl-1.1.1i.tar.gz --no-check-certificate
RUN tar xvf openssl-1.1.1i.tar.gz
RUN cd openssl-1.1.1i && ./config && make -j8 && make install
RUN rm -rf openssl-1.1.1i*

RUN wget https://github.com/Kitware/CMake/releases/download/v3.18.2/cmake-3.18.2.tar.gz
RUN tar xvf cmake-3.18.2.tar.gz
RUN cd cmake-3.18.2 && ./bootstrap && gmake -j8 && make install
RUN rm -rf cmake-3.18.2*

RUN yum install -y libstdc++-devel.i686 libgcc.i686 glibc-devel.i686

RUN wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz && rm -f go1.21.0.linux-amd64.tar.gz
RUN echo "export PATH=\$PATH:/usr/local/go/bin" >> /root/.bashrc
