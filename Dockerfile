FROM golang:1.21

WORKDIR /golang/app

COPY . /golang/app

RUN go install

EXPOSE 1200

CMD [ "go", "run", "main.go" ]