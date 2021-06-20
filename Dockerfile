FROM alpine
ADD scanner-job /scanner-job
ENTRYPOINT [ "/scanner-job" ]