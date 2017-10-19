FROM ubuntu:xenial

RUN apt-get update 
RUN apt-get -y install libpcre3-dev libssl-dev perl make build-essential curl
RUN apt-get -y install wget 
RUN wget -qO - https://openresty.org/package/pubkey.gpg | apt-key add -  
RUN apt-get -y install software-properties-common
RUN add-apt-repository -y "deb http://openresty.org/package/ubuntu $(lsb_release -sc) main"
RUN apt-get install build-essential
RUN apt-get -y install openresty
RUN apt-get -y install --no-install-recommends openresty

CMD ["sleep infinity"]

