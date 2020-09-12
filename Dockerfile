FROM golang

ADD . /delivery-api

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel
RUN ls
RUN ls /delivery-api

CMD ["revel", "run", "-a", "/delivery-api/app"]
