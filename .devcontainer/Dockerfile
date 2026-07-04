# Pinned to the current LTS on purpose: a floating ubuntu:latest silently broke
# the bpftool v7.4.0 build when 26.04's glibc/gcc landed (C23 const-correct
# strchr + -Werror). Upgrade this tag deliberately, together with the tool tags
# below and cmd/test_agent/test_agent.Dockerfile.
FROM ubuntu:26.04

RUN apt update && \
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
RUN apt install -y libssl-dev
RUN ln -sf /usr/include/asm-generic/ /usr/include/asm
RUN apt install -y libcap-dev
RUN apt install -y dwarves
RUN apt install -y locales
RUN apt install -y bash-completion
RUN apt install -y cargo
RUN locale-gen en_US.UTF-8
RUN update-locale LANG=en_US.UTF-8
RUN echo 'alias ll="ls -al"' >> ~/.bashrc
RUN echo 'if [ -f /etc/bash_completion ]; then . /etc/bash_completion; fi' >> ~/.bashrc
RUN mkdir -p ~/.local/share/bash-completion/completions
RUN echo 'docker completion bash > ~/.local/share/bash-completion/completions/docker' >> ~/.bashrc
RUN git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
RUN ~/.fzf/install --all

RUN echo 'PATH=$PATH:$HOME/.cargo/bin' >> ~/.bashrc

# bpftool and xdp-tools straight from the Ubuntu archive — no source builds.
# 26.04 ships standalone bpftool (v7.7) and xdp-tools/libxdp (v1.6).
RUN apt install -y bpftool xdp-tools libxdp-dev

# zellij — not packaged by Ubuntu; built from source via cargo (latest release).
RUN cargo install --locked --root /usr/local zellij && \
    rm -rf ~/.cargo/registry ~/.cargo/git && \
    zellij --version

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin
RUN mkdir -p $GOPATH/bin


RUN apt install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# install CLAB
# containelab version will be set in devcontainer.json
# ARG _CLAB_VERSION

# Set permissions for mounts in devcontainer.json
#RUN mkdir -p /home/vscode/.vscode-server/bin
#RUN chown -R vscode:vscode /home/vscode/.vscode-server

# install some basic tools inside the container
# adjust this list based on your demands
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    sshpass \
    curl \
    iputils-ping \
    htop \
    yamllint \
    && rm -rf /var/lib/apt/lists/* \
    && rm -Rf /usr/share/doc && rm -Rf /usr/share/man \
    && apt-get clean

# install preferred version of the containerlab
RUN bash -c "$(curl -sL https://get.containerlab.dev)"
