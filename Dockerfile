FROM alpine

RUN adduser -h /home/cnitch -D cnitch cnitch
#RUN usermod -aG docker cnitch

COPY ./cmd/cnitch /home/cnitch/
RUN chmod +x /home/cnitch/cnitch

USER cnitch

ENTRYPOINT ["/home/cnitch/cnitch"]
