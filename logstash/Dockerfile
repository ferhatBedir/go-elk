FROM docker.elastic.co/logstash/logstash-oss:8.7.1
#docker.elastic.co/logstash/logstash:8.7.1
COPY ./logstash/logstash.conf /etc/logstash/conf.d/
CMD logstash -f /etc/logstash/conf.d/logstash.conf