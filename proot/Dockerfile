FROM debian:9
RUN apt-get update && apt-get install -q -y build-essential git libseccomp-dev libtalloc-dev
ADD PRoot /PRoot
RUN cd PRoot/src && make 

FROM scratch
COPY --from=0 /PRoot/src/proot /runrootless-proot
