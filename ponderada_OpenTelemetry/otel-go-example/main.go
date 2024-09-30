package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "time"

    "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
    "go.opentelemetry.io/otel/sdk/resource"
    "go.opentelemetry.io/otel/sdk/trace"
    semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
    tp, err := initTracer()
    if err != nil {
        log.Fatalf("Erro ao inicializar o Tracer Provider: %v", err)
    }
    defer func() {
        if err := tp.Shutdown(context.Background()); err != nil {
            log.Fatalf("Erro ao desligar o Tracer Provider: %v", err)
        }
    }()

    // Definir o Tracer Provider global
    otel.SetTracerProvider(tp)

    // Instrumentar um cliente HTTP
    client := http.Client{
        Transport: otelhttp.NewTransport(http.DefaultTransport),
    }

    // Fazer uma requisição HTTP simulada
    req, _ := http.NewRequest("GET", "https://example.com", nil)
    ctx := context.Background()
    req = req.WithContext(ctx)

	// Obter o Tracer
	tracer := otel.Tracer("custom-tracer")

	// Iniciar um Span personalizado
	ctx, span := tracer.Start(ctx, "custom-operation")
	defer span.End()

	// Simular uma operação
	fmt.Println("Executando operação personalizada...")
	time.Sleep(2 * time.Second)
	fmt.Println("Operação personalizada concluída.")

    fmt.Println("Fazendo requisição HTTP...")
    res, err := client.Do(req)
    if err != nil {
        log.Fatalf("Erro na requisição HTTP: %v", err)
    }
    res.Body.Close()
    fmt.Println("Requisição concluída.")

    // Esperar um pouco para garantir que os spans sejam exportados
    time.Sleep(5 * time.Second)
}

func initTracer() (*trace.TracerProvider, error) {
    ctx := context.Background()

    // Configurar o exportador OTLP para enviar dados ao SigNoz
    exporter, err := otlptracehttp.New(ctx, otlptracehttp.WithEndpoint("localhost:4318"), otlptracehttp.WithInsecure())
    if err != nil {
        return nil, err
    }

    // Criar o Tracer Provider com o exportador
    tracerProvider := trace.NewTracerProvider(
        trace.WithBatcher(exporter),
        trace.WithResource(resource.NewWithAttributes(
            semconv.SchemaURL,
            semconv.ServiceNameKey.String("otel-go-example"),
        )),
    )

    return tracerProvider, nil
}

func helloWorld() string {
    return "Hello, OpenTelemetry!"
}