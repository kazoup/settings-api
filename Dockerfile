FROM scratch
ADD settings-api /
ENTRYPOINT [ "/settings-api" ]
