To run jaeger  

docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one  

  References:  
  -https://medium.com/@carlosedp/instrumenting-go-for-tracing-c5bdabe1fc81  
  -https://medium.com/opentracing/take-opentracing-for-a-hotrod-ride-f6e3141f7941  
  -https://github.com/carlosedp/go-playground