package h2c

func ExampleServeHTTP() {
        log.Fatal(http.ListenAndServe(":8080", &Server{}))
}

