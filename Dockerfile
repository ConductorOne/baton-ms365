FROM gcr.io/distroless/static-debian11:nonroot
ENTRYPOINT ["/baton-ms365"]
COPY baton-ms365 /