FROM golang:1.6

ENV builddir=/go/src/github.com/tebriel/follow_stats

RUN mkdir -p ${builddir}
WORKDIR ${builddir}

CMD ["go-wrapper", "run"]

COPY . ${builddir}
RUN go-wrapper download
RUN go-wrapper install

EXPOSE 8080
