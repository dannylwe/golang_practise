package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

func main() {
	const PORT = ":8181"
	fmt.Println("starting application on port " + PORT)

	cfg := jaegercfg.Configuration{
		ServiceName: "sleeper",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans: true,
		},
	}

	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	tracer, closer, err := cfg.NewTracer(
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)

	opentracing.SetGlobalTracer(tracer)

	if err != nil {
		log.Fatalf("could not initialize jaeger tracer: %s", err.Error())
	}
	defer closer.Close()

	f1 := traceFunction1(tracer)
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "")
		return
	})
	http.HandleFunc("/", f1)
	log.Fatalln(http.ListenAndServe(":8181", nil))
}

func f2(ctx context.Context) string {
	span, ctx := opentracing.StartSpanFromContext(ctx, "f2")
	defer span.Finish()

	slp := time.Duration(rand.Intn(900)) * time.Nanosecond
	span.LogKV("sleep", slp)
	time.Sleep(slp)

	return "f2 done"
}

func f1(ctx context.Context) (string, string) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "f1")
	defer span.Finish()

	slp := time.Duration(rand.Intn(1000)) * time.Millisecond
	span.LogKV("sleep", slp)
	time.Sleep(slp)

	f2Val := f2(ctx)
	f1Val := "f1 done: "

	return f1Val, f2Val
}

func traceFunction1(tracer opentracing.Tracer) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		spanCtx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil {
			log.Println(err)
		}
		span := tracer.StartSpan("callF1", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		ctx = opentracing.ContextWithSpan(ctx, span)
		val1, val2 := f1(ctx)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, val1+val2)
	}
}
