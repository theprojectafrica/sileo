FROM gcr.io/distroless/static
COPY sileo /usr/local/bin/sileo
ENTRYPOINT [ "/usr/local/bin/sileo" ]
