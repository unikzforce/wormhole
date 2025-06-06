FROM ubuntu:24.04

RUN apt-get update && \
    echo "wireshark-common wireshark-common/install-setuid boolean true" | debconf-set-selections && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y tshark

RUN apt install -y golang-go
RUN apt install -y build-essential
RUN apt install -y pkg-config
RUN apt install -y clang
RUN apt install -y llvm
RUN apt install -y m4
RUN apt install -y git
RUN apt install -y libelf-dev
RUN apt install -y libpcap-dev
RUN apt install -y iproute2
RUN apt install -y iputils-ping
RUN apt install -y linux-headers-generic
RUN apt install -y libbpf-dev
RUN apt install -y linux-libc-dev
RUN apt install -y cmake
RUN apt install -y libpcap-dev
RUN apt install -y libcap-ng-dev
RUN apt install -y libbfd-dev
RUN ln -sf /usr/include/asm-generic/ /usr/include/asm
RUN apt install -y libcap-dev
RUN apt install -y dwarves
RUN apt install -y ethtool
RUN echo 'alias ll="ls -al"' >> ~/.bashrc
RUN git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
RUN ~/.fzf/install --all

RUN mkdir /tools/

WORKDIR /tools/

RUN git clone --recurse-submodules https://github.com/libbpf/bpftool.git

RUN make -C bpftool/src/ install

RUN git clone --recurse-submodules https://github.com/xdp-project/xdp-tools.git

RUN make -C xdp-tools/ install

RUN mkdir /wormhole-source/

COPY ../../go.mod /wormhole-source
COPY ../../go.sum /wormhole-source
COPY ../../internal/ /wormhole-source/internal/
COPY ../../include/ /wormhole-source/include/
COPY ../../cmd/ /wormhole-source/cmd/

WORKDIR /wormhole-source/

#RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go mod download
RUN go generate ./internal/switch_agent/ebpf/
RUN go generate ./internal/vxlan_agent/ebpf/
RUN go generate ./internal/dummy_xdp/ebpf/

RUN go build -o /build/switch_agent ./cmd/switch_agent/main.go
RUN go build -o /build/vxlan_agent ./cmd/vxlan_agent/main.go
RUN go build -o /build/dummy_xdp ./cmd/dummy_xdp/main.go
RUN go build -o /build/test_agent ./cmd/test_agent/test_agent_grpc.server_bootstrap.go
# for debuging comment above and uncomment below 2 commands & comment above
#RUN go build -gcflags="all=-N -l" -o /build/switch_agent ./cmd/switch_agent/main.go
#EXPOSE 40000



#ENTRYPOINT ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/build/test_agent"]

#ENTRYPOINT ["sh", "-c", "trap 'exit 0' SIGTERM SIGINT; while true; do echo 'Container is running...'; sleep 10; done"]

ENTRYPOINT ["/build/test_agent"]