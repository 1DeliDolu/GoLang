# Go Testing Examples

Dieses Verzeichnis demonstriert die grundlegenden Konzepte und Praktiken für das Testen in Go.

## Einführung

Go bietet ein eingebautes Testing-Framework über das `testing`-Paket. Tests in Go sind:
- Einfach zu schreiben und auszuführen
- Integriert in das Go-Toolset
- Unterstützen verschiedene Testarten wie Unit-Tests, Benchmark-Tests und Beispiel-Tests

## Testdateien in Go

Testdateien in Go:
- Enden mit `_test.go` (z.B. `add_test.go`)
- Befinden sich im selben Paket wie der zu testende Code (in diesem Fall `package main`)
- Enthalten Funktionen, die mit `Test`, `Benchmark` oder `Example` beginnen

## Einen Test ausführen

Um Tests auszuführen:

```bash
# Im aktuellen Verzeichnis
go test

# Mit ausführlicher Ausgabe
go test -v

# Einen bestimmten Test ausführen
go test -run TestAdd

# Benchmarks ausführen
go test -bench=.
```

## Beispiel

Dieses Verzeichnis enthält einen einfachen Test für eine Addition-Funktion:

1. `add.go` - enthält die Funktion `add(a, b int) int`, die wir testen möchten
2. `add_test.go` - enthält die Testfunktion `TestAdd`, die verschiedene Testfälle überprüft

### Die Table-Driven Tests Methode

In `add_test.go` verwenden wir die "Table-Driven Tests"-Methode, ein gängiges Muster in Go:

```go
func TestAdd(t *testing.T) {
    // test cases
    cases := []struct {
        a, b, want int
    }{
        {1, 2, 3},
        {2, 3, 5},
        // ...weitere Testfälle
    }
    
    // Über alle Testfälle iterieren
    for _, c := range cases {
        got := add(c.a, c.b)
        if got != c.want {
            t.Errorf("add(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
        }
    }
}
```

Vorteile dieser Methode:
- Einfaches Hinzufügen weiterer Testfälle
- Übersichtlicher Code
- Wartungsfreundlich
- Konsistente Fehlermeldungen

## Weitere Testarten

### Benchmark-Tests

Werden verwendet, um die Performance zu messen:

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        add(1, 2)
    }
}
```

### Beispiel-Tests

Dienen als Dokumentation und testbare Beispiele:

```go
func ExampleAdd() {
    sum := add(1, 2)
    fmt.Println(sum)
    // Output: 3
}
```

## Best Practices für Go Tests

1. **Isolierte Tests**: Tests sollten unabhängig voneinander ausführbar sein
2. **Aussagekräftige Fehlermeldungen**: Verwenden Sie `t.Errorf()` oder `t.Fatalf()` mit klaren Nachrichten
3. **Table-Driven Tests**: Für wiederholbare Testfälle
4. **Test-Helper**: Extrahieren Sie wiederverwendbare Testlogik in separate Hilfsfunktionen
5. **Testabdeckung**: Überprüfen Sie die Testabdeckung mit `go test -cover`

## Testabdeckung messen

Führen Sie folgende Befehle aus, um die Testabdeckung zu messen:

```bash
# Testabdeckung als Prozentsatz anzeigen
go test -cover

# HTML-Bericht generieren
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## Tipps für effektives Testen

- Schreiben Sie Tests während der Entwicklung, nicht danach
- Testen Sie Edge Cases (Randfälle) und Fehlerszenarien
- Halten Sie Tests einfach und fokussiert
- Verwenden Sie das Muster "Arrange-Act-Assert" (Vorbereiten-Ausführen-Überprüfen)

## Weitere Informationen

- [Go Testing-Paket Dokumentation](https://pkg.go.dev/testing)
- [Go Blog: Using Subtests and Sub-benchmarks](https://blog.golang.org/subtests)
- [Go Proverbs](https://go-proverbs.github.io/)
