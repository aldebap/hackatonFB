FROM java:8
MAINTAINER EBCDIC <ebcdic@cielo.com.br>
VOLUME /tmp
ADD springBootDocker-ebcedic.jar Myapp.jar
RUN sh -c 'touch /MyApp-ebcedic.jar'
ENV JAVA_OPTS=""
ENTRYPOINT [ "sh", "-c", "java $JAVA_OPTS -Djava.security.egd=file:/dev/./urandom -jar /MyApp-ebcedic.jar" ]