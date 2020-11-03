FROM ubuntu:20.10
RUN apt-get update && apt-get install gnupg2 -y
ADD image-signature-webhook /usr/local/bin/image-signature-webhook
ENTRYPOINT ["/usr/local/bin/image-signature-webhook"]
