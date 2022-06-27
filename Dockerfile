## NOTE: This image uses goreleaser to build image
# if building manually please run go build ./cmd/traverse first and then build

# Choose alpine as a base image to make this useful for CI, as many
# CI tools expect an interactive shell inside the container
FROM alpine:latest as production

#COPY --from=builder /build/traverse /usr/bin/traverse
COPY traverse /usr/bin/traverse
RUN chmod +x /usr/bin/traverse

WORKDIR /workdir

ENTRYPOINT ["/usr/bin/traverse"]
CMD ["--help"]
